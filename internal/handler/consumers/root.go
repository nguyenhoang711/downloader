package consumers

import (
	"context"
	"encoding/json"

	"github.com/nguyenhoang711/downloader/internal/dataaccess/mq/consumer"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/mq/producer"
	"go.uber.org/zap"
)

type RootConsumer interface {
	Start(ctx context.Context) error
}

type rootConsumer struct {
	downloadTaskCreatedHandler DownloadTaskCreated
	mqConsumer                 consumer.Consumer
	logger                     *zap.Logger
}

func NewRootConsumer(
	downloadTaskCreatedHandler DownloadTaskCreated,
	mqConsumer consumer.Consumer,
	logger *zap.Logger,
) RootConsumer {
	return &rootConsumer{
		downloadTaskCreatedHandler: downloadTaskCreatedHandler,
		mqConsumer:                 mqConsumer,
		logger:                     logger,
	}
}

// Start implements RootConsumer.
func (r rootConsumer) Start(ctx context.Context) error {
	r.mqConsumer.RegisterHandler(
		producer.MessageQueueDownloadTaskCreated,
		func(ctx context.Context, queueName string, payload []byte) error {
			var event producer.DownloadTaskCreated
			if err := json.Unmarshal(payload, &event); err != nil {
				return err
			}

			return r.downloadTaskCreatedHandler.Handle(ctx, event)
		},
	)

	return r.mqConsumer.Start(ctx)
}
