package logic

import (
	"context"
	"errors"
	"io"
	"net/url"
	"path"

	"github.com/doug-martin/goqu/v9"
	"github.com/gammazero/workerpool"
	"github.com/nguyenhoang711/downloader/internal/configs"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/database"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/file"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/mq/producer"
	go_load "github.com/nguyenhoang711/downloader/internal/generated/go_load/v1"
	"github.com/nguyenhoang711/downloader/internal/utils"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	downloadTaskMetadataFieldNameFileName = "file-name"
)

type CreateDownloadTaskParams struct {
	Token        string
	DownloadType go_load.DownloadType
	URL          string
}

type CreateDownloadTaskOutput struct {
	DownloadTask *go_load.DownloadTask
}

type GetDownloadTaskListParams struct {
	Token  string
	Limit  uint64
	Offset uint64
}

type GetDownloadTaskFileParams struct {
	Token          string
	DownloadTaskID uint64
}

type GetDownloadTaskListOutput struct {
	DownloadTaskList       []*go_load.DownloadTask
	TotalDownloadTaskCount uint64
}

type UpdateDownloadTaskParams struct {
	DownloadTaskID uint64
	Token          string
	URL            string
}

type UpdateDownloadTaskOutput struct {
	UpdatedDownloadTask *go_load.DownloadTask
}

type DeleteDownloadTaskParams struct {
	Token          string
	DownloadTaskID uint64
}

type DownloadTaskLogic interface {
	CreateDownloadTask(ctx context.Context, params CreateDownloadTaskParams) (CreateDownloadTaskOutput, error)
	GetDownloadTaskList(ctx context.Context, params GetDownloadTaskListParams) (GetDownloadTaskListOutput, error)
	UpdateDownloadTask(ctx context.Context, params UpdateDownloadTaskParams) (UpdateDownloadTaskOutput, error)
	DeleteDownloadTask(ctx context.Context, params DeleteDownloadTaskParams) error
	ExecuteDownloadTask(ctx context.Context, id uint64) error
	GetDownloadTaskFile(context.Context, GetDownloadTaskFileParams) (io.ReadCloser, error)
	ExecuteAllPendingDownloadTask(context.Context) error
	UpdateDownloadingAndFailedDownloadTaskStatusToPending(context.Context) error
}

type downloadTask struct {
	tokenLogic                  Token
	downloadTaskDataAccessor    database.DownloadTaskDataAccessor
	accountDataAccessor         database.AccountDataAccessor
	goquDatabase                *goqu.Database
	fileClient                  file.Client
	logger                      *zap.Logger
	cronConfig                  configs.Cron
	downloadTaskCreatedProducer producer.DownloadTaskCreatedProducer
}

func NewDownloadTask(
	tokenLogic Token,
	downloadTaskDataAccessor database.DownloadTaskDataAccessor,
	accountDataAccessor database.AccountDataAccessor,
	fileClient file.Client,
	goquDatabase *goqu.Database,
	logger *zap.Logger,
	cronConfig configs.Cron,
	downloadTaskCreatedProducer producer.DownloadTaskCreatedProducer,
) DownloadTaskLogic {
	return &downloadTask{
		tokenLogic:                  tokenLogic,
		downloadTaskDataAccessor:    downloadTaskDataAccessor,
		accountDataAccessor:         accountDataAccessor,
		goquDatabase:                goquDatabase,
		fileClient:                  fileClient,
		logger:                      logger,
		cronConfig:                  cronConfig,
		downloadTaskCreatedProducer: downloadTaskCreatedProducer,
	}
}

func extractFileName(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return path.Base(parsedURL.Path), nil
}

func (d downloadTask) databaseDownloadTaskToProtoDownloadTask(
	downloadTask database.DownloadTask,
	account database.Account,
) *go_load.DownloadTask {
	return &go_load.DownloadTask{
		Id: downloadTask.ID,
		OfAccount: &go_load.Account{
			Id:          account.ID,
			AccountName: account.AccountName,
		},
		DownloadType:   downloadTask.DownloadType,
		Url:            downloadTask.URL,
		DownloadStatus: go_load.DownloadStatus_DOWNLOAD_STATUS_PENDING,
	}
}

// CreateDownloadTask implements DownloadTaskLogic.
func (d downloadTask) CreateDownloadTask(
	ctx context.Context,
	params CreateDownloadTaskParams,
) (CreateDownloadTaskOutput, error) {
	accID, _, err := d.tokenLogic.GetAccountIDAndExpireTime(ctx, params.Token)
	if err != nil {
		return CreateDownloadTaskOutput{}, err
	}

	account, err := d.accountDataAccessor.GetAccountByID(ctx, accID)
	if err != nil {
		return CreateDownloadTaskOutput{}, err
	}
	downloadTask := database.DownloadTask{
		OfAccountID:    accID,
		DownloadType:   params.DownloadType,
		URL:            params.URL,
		DownloadStatus: go_load.DownloadStatus_DOWNLOAD_STATUS_PENDING,
		Metadata: database.JSON{
			Data: make(map[string]any),
		},
	}

	txErr := d.goquDatabase.WithTx(func(td *goqu.TxDatabase) error {
		downloadTaskID, createDownloadTaskErr := d.downloadTaskDataAccessor.
			WithDatabase(td).
			CreateDownloadTask(ctx, downloadTask)
		if createDownloadTaskErr != nil {
			return createDownloadTaskErr
		}

		downloadTask.ID = downloadTaskID
		// produce event send to consumers
		produceErr := d.downloadTaskCreatedProducer.Produce(ctx, producer.DownloadTaskCreated{
			ID: downloadTaskID,
		})
		if produceErr != nil {
			return produceErr
		}
		return nil
	})
	if txErr != nil {
		return CreateDownloadTaskOutput{}, txErr
	}
	return CreateDownloadTaskOutput{
		DownloadTask: d.databaseDownloadTaskToProtoDownloadTask(downloadTask, account),
	}, nil
}

// DeleteDownloadTask implements DownloadTaskLogic.
func (d downloadTask) DeleteDownloadTask(ctx context.Context, params DeleteDownloadTaskParams) error {
	accountID, _, err := d.tokenLogic.GetAccountIDAndExpireTime(ctx, params.Token)
	if err != nil {
		return err
	}

	return d.goquDatabase.WithTx(func(td *goqu.TxDatabase) error {
		downloadTask, getDownloadTaskWithXLockErr := d.downloadTaskDataAccessor.WithDatabase(td).
			GetDownloadTaskWithXLock(ctx, params.DownloadTaskID)
		if getDownloadTaskWithXLockErr != nil {
			return getDownloadTaskWithXLockErr
		}

		if downloadTask.OfAccountID != accountID {
			return status.Error(codes.PermissionDenied, "trying to delete a download task the account does not own")
		}

		return d.downloadTaskDataAccessor.WithDatabase(td).DeleteDownloadTask(ctx, params.DownloadTaskID)
	})
}

// GetDownloadTaskList implements DownloadTaskLogic.
func (d downloadTask) GetDownloadTaskList(
	ctx context.Context,
	params GetDownloadTaskListParams,
) (GetDownloadTaskListOutput, error) {
	accountID, _, err := d.tokenLogic.GetAccountIDAndExpireTime(ctx, params.Token)
	if err != nil {
		return GetDownloadTaskListOutput{}, err
	}

	account, err := d.accountDataAccessor.GetAccountByID(ctx, accountID)
	if err != nil {
		return GetDownloadTaskListOutput{}, err
	}

	totalDownloadTaskCount, err := d.downloadTaskDataAccessor.GetDownloadTaskCountOfAccount(ctx, accountID)
	if err != nil {
		return GetDownloadTaskListOutput{}, err
	}

	downloadTaskList, err := d.downloadTaskDataAccessor.
		GetDownloadTaskListOfAccount(ctx, accountID, params.Offset, params.Limit)
	if err != nil {
		return GetDownloadTaskListOutput{}, err
	}

	return GetDownloadTaskListOutput{
		TotalDownloadTaskCount: totalDownloadTaskCount,
		DownloadTaskList: lo.Map(downloadTaskList, func(item database.DownloadTask, _ int) *go_load.DownloadTask {
			return d.databaseDownloadTaskToProtoDownloadTask(item, account)
		}),
	}, nil
}

// UpdateDownloadTask implements DownloadTaskLogic.
func (d downloadTask) UpdateDownloadTask(
	ctx context.Context,
	params UpdateDownloadTaskParams,
) (UpdateDownloadTaskOutput, error) {
	accountID, _, err := d.tokenLogic.GetAccountIDAndExpireTime(ctx, params.Token)
	if err != nil {
		return UpdateDownloadTaskOutput{}, err
	}

	account, err := d.accountDataAccessor.GetAccountByID(ctx, accountID)
	if err != nil {
		return UpdateDownloadTaskOutput{}, err
	}

	output := UpdateDownloadTaskOutput{}
	txErr := d.goquDatabase.WithTx(func(td *goqu.TxDatabase) error {
		downloadTask, getDownloadTaskWithXLockErr := d.downloadTaskDataAccessor.WithDatabase(td).
			GetDownloadTaskWithXLock(ctx, params.DownloadTaskID)
		if getDownloadTaskWithXLockErr != nil {
			return getDownloadTaskWithXLockErr
		}

		if downloadTask.OfAccountID != accountID {
			return status.Error(codes.PermissionDenied, "trying to update a download task the account does not own")
		}

		downloadTask.URL = params.URL
		output.UpdatedDownloadTask = d.databaseDownloadTaskToProtoDownloadTask(downloadTask, account)
		return d.downloadTaskDataAccessor.WithDatabase(td).UpdateDownloadTask(ctx, downloadTask)
	})
	if txErr != nil {
		return UpdateDownloadTaskOutput{}, txErr
	}

	return output, nil
}

func (d downloadTask) updateDownloadTaskStatusFromPendingToDownloading(
	ctx context.Context,
	id uint64,
) (bool, database.DownloadTask, error) {
	var (
		logger       = utils.LoggerWithContext(ctx, d.logger).With(zap.Uint64("id", id))
		updated      = false
		downloadTask database.DownloadTask
		err          error
	)

	txErr := d.goquDatabase.WithTx(func(td *goqu.TxDatabase) error {
		// get download task by id
		// WithDatabase: su dung database trong thuc hien transaction
		downloadTask, err = d.downloadTaskDataAccessor.WithDatabase(td).GetDownloadTaskWithXLock(ctx, id)
		if err != nil {
			if errors.Is(err, database.ErrDownloadTaskNotFound) {
				logger.Warn("download task not found, will skip")
				return nil
			}

			logger.With(zap.Error(err)).Error("failed to get download task")
			return err
		}

		if downloadTask.DownloadStatus != go_load.DownloadStatus_DOWNLOAD_STATUS_PENDING {
			logger.Warn("download task is not in pending status, will not execute")
			updated = false
			return nil
		}

		downloadTask.DownloadStatus = go_load.DownloadStatus_DOWNLOAD_STATUS_DOWNLOADING
		// update download_task status (only one request update)
		err = d.downloadTaskDataAccessor.WithDatabase(td).UpdateDownloadTask(ctx, downloadTask)
		if err != nil {
			logger.With(zap.Error(err)).Error("failed to update download task")
			return err
		}

		updated = true
		return nil
	})
	if txErr != nil {
		return false, database.DownloadTask{}, err
	}

	return updated, downloadTask, nil
}

func (d downloadTask) ExecuteDownloadTask(ctx context.Context, id uint64) error {
	logger := utils.LoggerWithContext(ctx, d.logger)

	updated, downloadTask, err := d.updateDownloadTaskStatusFromPendingToDownloading(ctx, id)
	if err != nil {
		return err
	}
	if !updated {
		return nil
	}
	// viet vao 1 writer (noi dung download)
	var downloader Downloader
	//nolint:exhaustive // No need to check unsupported download type
	switch downloadTask.DownloadType {
	case go_load.DownloadType_DOWNLOAD_TYPE_HTTP:
		downloader = NewHTTPDownloader(downloadTask.URL, d.logger)
	default:
		logger.With(zap.Any("download_type", downloadTask.DownloadType)).Error("unsupported download type")
		d.updateDownloadTaskStatusToFailed(ctx, downloadTask)
		return nil
	}

	fileName, err := extractFileName(downloadTask.URL)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get download file name")
		return err
	}
	// cung cap viet vao dau
	fileWriteCloser, err := d.fileClient.Write(ctx, fileName)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get download file writer")
		d.updateDownloadTaskStatusToFailed(ctx, downloadTask)
		return err
	}

	// download ve fileWriterCloser
	defer fileWriteCloser.Close()

	metadata, err := downloader.Download(ctx, fileWriteCloser)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to download")
		d.updateDownloadTaskStatusToFailed(ctx, downloadTask)
		return err
	}

	metadata[downloadTaskMetadataFieldNameFileName] = fileName
	downloadTask.DownloadStatus = go_load.DownloadStatus_DOWNLOAD_STATUS_SUCCESS
	downloadTask.Metadata = database.JSON{
		Data: metadata,
	}
	err = d.downloadTaskDataAccessor.UpdateDownloadTask(ctx, downloadTask)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to update download task status to success")
		return err
	}

	logger.Info("download task executed successfully")

	return nil
}

func (d downloadTask) GetDownloadTaskFile(
	ctx context.Context,
	params GetDownloadTaskFileParams,
) (io.ReadCloser, error) {
	accountID, _, err := d.tokenLogic.GetAccountIDAndExpireTime(ctx, params.Token)
	if err != nil {
		return nil, err
	}

	downloadTask, err := d.downloadTaskDataAccessor.GetDownloadTask(ctx, params.DownloadTaskID)
	if err != nil {
		return nil, err
	}

	if downloadTask.OfAccountID != accountID {
		return nil, status.Error(codes.PermissionDenied, "trying to get file of a download task the account does not own")
	}

	if downloadTask.DownloadStatus != go_load.DownloadStatus_DOWNLOAD_STATUS_SUCCESS {
		return nil, status.Error(codes.InvalidArgument, "download task does not have status of success")
	}

	downloadTaskMetadata, ok := downloadTask.Metadata.Data.(map[string]any)
	if !ok {
		return nil, status.Error(codes.Internal, "download task metadata is not a map[string]any")
	}

	fileName, ok := downloadTaskMetadata[downloadTaskMetadataFieldNameFileName]
	if !ok {
		return nil, status.Error(codes.Internal, "download task metadata does not contain file name")
	}

	return d.fileClient.Read(ctx, fileName.(string))
}

// cronjob chủ động query db để thực hiện
func (d downloadTask) ExecuteAllPendingDownloadTask(ctx context.Context) error {
	logger := utils.LoggerWithContext(ctx, d.logger)

	//collect all id of pending tasks
	pendingDownloadTaskIDList, err := d.downloadTaskDataAccessor.GetPendingDownloadTaskIDList(ctx)
	if err != nil {
		return err
	}
	if len(pendingDownloadTaskIDList) == 0 {
		logger.Info("no pending download task found")
		return nil
	}

	logger.
		With(zap.Int("len(pending_download_task_id_list)", len(pendingDownloadTaskIDList))).
		Info("pending download task found")

	// các tác vụ chủ động đấy đặt ở trong worker pool
	// concurrency limit: số lượng tác vụ chạy đồng thời (CPU chờ sẵn)
	workerPool := workerpool.New(d.cronConfig.ExecuteAllPendingDownloadTask.ConcurrencyLimit)
	for _, id := range pendingDownloadTaskIDList {
		workerPool.Submit(func() {
			// run execute for each task
			// hàm này được cắm thụ động vào kafka để nhận các sự kiện execute download task  mới đc tạo ra
			if executeDownloadTaskErr := d.ExecuteDownloadTask(ctx, id); executeDownloadTaskErr != nil {
				logger.
					With(zap.Uint64("download_task_id", id)).
					With(zap.Error(executeDownloadTaskErr)).
					Error("failed to execute download_task")
			}
		})
	}

	workerPool.StopWait()
	return nil
}

func (d downloadTask) updateDownloadTaskStatusToFailed(ctx context.Context, downloadTask database.DownloadTask) {
	logger := utils.LoggerWithContext(ctx, d.logger).With(zap.Uint64("id", downloadTask.ID))

	downloadTask.DownloadStatus = go_load.DownloadStatus_DOWNLOAD_STATUS_FAILED
	updateDownloadTaskErr := d.downloadTaskDataAccessor.UpdateDownloadTask(ctx, downloadTask)
	if updateDownloadTaskErr != nil {
		logger.With(zap.Error(updateDownloadTaskErr)).Warn("failed to update download task status to failed")
	}
}

// download task trạng thái downloading quá lâu hoặc failed về pending
func (d downloadTask) UpdateDownloadingAndFailedDownloadTaskStatusToPending(ctx context.Context) error {
	return d.downloadTaskDataAccessor.UpdateDownloadingAndFailedDownloadTaskStatusToPending(ctx)
}
