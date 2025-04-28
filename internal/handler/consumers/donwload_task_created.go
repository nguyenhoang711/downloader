package consumers

import (
	"context"

	"github.com/nguyenhoang711/downloader/internal/dataaccess/mq/producer"
	"github.com/nguyenhoang711/downloader/internal/logic"
	"github.com/nguyenhoang711/downloader/internal/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DownloadTaskCreated interface {
	Handle(ctx context.Context, event producer.DownloadTaskCreated) error
}

type downloadTaskCreated struct {
	downloadTaskLogic logic.DownloadTaskLogic
	logger            *zap.Logger
}

func NewDownloadTaskCreated(
	downloadTaskLogic logic.DownloadTaskLogic,
	logger *zap.Logger,
) DownloadTaskCreated {
	return &downloadTaskCreated{
		downloadTaskLogic: downloadTaskLogic,
		logger:            logger,
	}
}

func (d downloadTaskCreated) Handle(ctx context.Context, event producer.DownloadTaskCreated) error {
	logger := utils.LoggerWithContext(ctx, d.logger).With(zap.Any("event", event))
	logger.Info("download task created event received")

	if err := d.downloadTaskLogic.ExecuteDownloadTask(ctx, event.ID); err != nil {
		logger.With(zap.Error(err)).Error("failed to handle download task")
		return status.Error(codes.Internal, "failed to handle download task")
	}

	return nil
}
