package logic

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/database"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/mq/producer"
	"github.com/nguyenhoang711/downloader/internal/generated/grpc/go_load"
	"go.uber.org/zap"
)

type CreateDownloadTaskParams struct {
	Token        string
	DownloadType go_load.DownloadType
	URL          string
}

type CreateDownloadTaskOutput struct {
	DownloadTask go_load.DownloadTask
}

type GetDownloadTaskListParams struct {
	Token  string
	Limit  int8
	Offset int8
}

type GetDownloadTaskListOutput struct {
	ListDownloadTasks      []go_load.DownloadTask
	TotalDownloadTaskCount uint64
}

type UpdateDownloadTaskParams struct {
	ID    uint64
	Token string
	URL   string
}

type UpdateDownloadTaskOutput struct {
	UpdatedDownloadTask go_load.DownloadTask
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
	goquDatabase                *goqu.Database
	logger                      *zap.Logger
	downloadTaskCreatedProducer producer.DownloadTaskCreatedProducer
}

func NewDownloadTask(
	tokenLogic Token,
	downloadTaskDataAccessor database.DownloadTaskDataAccessor,
	goquDatabase *goqu.Database,
	logger *zap.Logger,
	downloadTaskCreatedProducer producer.DownloadTaskCreatedProducer,
) DownloadTaskLogic {
	return &downloadTask{
		tokenLogic:                  tokenLogic,
		downloadTaskDataAccessor:    downloadTaskDataAccessor,
		goquDatabase:                goquDatabase,
		logger:                      logger,
		downloadTaskCreatedProducer: downloadTaskCreatedProducer,
	}
}

// CreateDownloadTask implements DownloadTaskLogic.
func (d downloadTask) CreateDownloadTask(
	ctx context.Context,
	params CreateDownloadTaskParams,
) (CreateDownloadTaskOutput, error) {
	accId, _, err := d.tokenLogic.GetAccountIDAndExpireTime(ctx, params.Token)
	if err != nil {
		return CreateDownloadTaskOutput{}, err
	}
	databaseDownloadTask := database.DownloadTask{
		OfAccountID:    accId,
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
			CreateDownloadTask(ctx, databaseDownloadTask)
		if createDownloadTaskErr != nil {
			return createDownloadTaskErr
		}

		databaseDownloadTask.ID = downloadTaskID
		// produce event send to consumers
		produceErr := d.downloadTaskCreatedProducer.Produce(ctx, producer.DownloadTaskCreated{
			DownloadTask: databaseDownloadTask,
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
		DownloadTask: go_load.DownloadTask{
			Id:             databaseDownloadTask.ID,
			OfAccount:      nil,
			DownloadType:   databaseDownloadTask.DownloadType,
			Url:            databaseDownloadTask.URL,
			DownloadStatus: databaseDownloadTask.DownloadStatus,
		},
	}, nil
}

// DeleteDownloadTask implements DownloadTaskLogic.
func (d downloadTask) DeleteDownloadTask(ctx context.Context, params DeleteDownloadTaskParams) error {
	panic("unimplemented")
}

// GetDownloadTaskList implements DownloadTaskLogic.
func (d downloadTask) GetDownloadTaskList(ctx context.Context, params GetDownloadTaskListParams) (GetDownloadTaskListOutput, error) {
	panic("unimplemented")
}

// UpdateDownloadTask implements DownloadTaskLogic.
func (d downloadTask) UpdateDownloadTask(ctx context.Context, params UpdateDownloadTaskParams) (UpdateDownloadTaskOutput, error) {
	panic("unimplemented")
}
