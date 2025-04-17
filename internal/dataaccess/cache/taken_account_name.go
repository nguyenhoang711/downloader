package cache

import (
	"context"

	"go.uber.org/zap"
)

const (
	setKeyNameTakenAccountName = "taken_account_name_set"
)

type TakenAccountName interface {
	Add(ctx context.Context, accountName string) error
	Has(ctx context.Context, accountName string) (bool, error)
}

type takenAccountName struct {
	client Client
	logger *zap.Logger
}

func NewTakenAccountName(
	client Client,
	logger *zap.Logger,
) TakenAccountName {
	return &takenAccountName{
		client: client,
		logger: logger,
	}
}

// Add implements TakenAccountName.
func (t takenAccountName) Add(ctx context.Context, accountName string) error {
	if err := t.client.AddToSet(ctx, setKeyNameTakenAccountName, accountName); err != nil {
		return err
	}
	return nil
}

// Has implements TakenAccountName.
func (t takenAccountName) Has(ctx context.Context, accountName string) (bool, error) {
	return t.client.IsDataInSet(ctx, setKeyNameTakenAccountName, accountName)
}
