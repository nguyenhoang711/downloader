package producer

import (
	"context"

	"go.uber.org/zap"
)

type Client interface {
	Produce(ctx context.Context, queueName string, payload []byte) error
}

type client struct {
	logger *zap.Logger
}

func NewClient(
	logger *zap.Logger,
) Client {
	return &client{
		logger: logger,
	}
}

// Produce implements Client.
func (c *client) Produce(ctx context.Context, queueName string, payload []byte) error {
	panic("unimplemented")
}
