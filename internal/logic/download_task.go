package logic

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/database"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/mq/producer"
	"github.com/nguyenhoang711/downloader/internal/generated/grpc/go_load"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
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

type GetDownloadTaskListOutput struct {
	DownloadTaskList       []*go_load.DownloadTask
	TotalDownloadTaskCount uint64
}

type UpdateDownloadTaskParams struct {
	ID    uint64
	Token string
	URL   string
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
}

type downloadTask struct {
	tokenLogic                  Token
	downloadTaskDataAccessor    database.DownloadTaskDataAccessor
	accountDataAccessor         database.AccountDataAccessor
	goquDatabase                *goqu.Database
	logger                      *zap.Logger
	downloadTaskCreatedProducer producer.DownloadTaskCreatedProducer
}

func NewDownloadTask(
	tokenLogic Token,
	downloadTaskDataAccessor database.DownloadTaskDataAccessor,
	accountDataAccessor database.AccountDataAccessor,
	goquDatabase *goqu.Database,
	logger *zap.Logger,
	downloadTaskCreatedProducer producer.DownloadTaskCreatedProducer,
) DownloadTaskLogic {
	return &downloadTask{
		tokenLogic:                  tokenLogic,
		downloadTaskDataAccessor:    downloadTaskDataAccessor,
		accountDataAccessor:         accountDataAccessor,
		goquDatabase:                goquDatabase,
		logger:                      logger,
		downloadTaskCreatedProducer: downloadTaskCreatedProducer,
	}
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
		DownloadStatus: go_load.DownloadStatus_Pending,
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
		DownloadStatus: go_load.DownloadStatus_Pending,
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
	panic("unimplemented")
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
		DownloadTaskList: lo.Map(downloadTaskList, func(item database.DownloadTask, index int) *go_load.DownloadTask {
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
		output.DownloadTask = d.databaseDownloadTaskToProtoDownloadTask(downloadTask, account)
		return d.downloadTaskDataAccessor.WithDatabase(td).UpdateDownloadTask(ctx, downloadTask)
	})
	if txErr != nil {
		return UpdateDownloadTaskOutput{}, txErr
	}

	return output, nil
}
