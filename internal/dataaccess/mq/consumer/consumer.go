package consumer

import "context"

type HandlerFunc func(ctx context.Context, queueName string, payload []byte) error

type Consumer interface {
	RegisterHandler(queueName string, handlerFunc HandlerFunc) error
	Start(ctx context.Context) error
}
