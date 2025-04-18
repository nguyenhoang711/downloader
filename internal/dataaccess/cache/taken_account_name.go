package cache

import (
	"context"

	"github.com/nguyenhoang711/downloader/internal/utils"
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
	logger := utils.LoggerWithContext(ctx, t.logger).With(zap.String("account_name", accountName))
	if err := t.client.AddToSet(ctx, setKeyNameTakenAccountName, accountName); err != nil {
		logger.With(zap.Error(err)).Error("failed to add account name to set in cache")
		return err
	}
	return nil
}

// Has implements TakenAccountName.
func (t takenAccountName) Has(ctx context.Context, accountName string) (bool, error) {
	xLogger := utils.LoggerWithContext(ctx, t.logger).With(zap.String("account_name", accountName))
	res, err := t.client.IsDataInSet(ctx, setKeyNameTakenAccountName, accountName)
	if err != nil {
		xLogger.With(zap.Error(err)).Error("failed to fetch account name in redis")
		return false, err
	}
	return res, nil
}
