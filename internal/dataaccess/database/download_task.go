package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/nguyenhoang711/downloader/internal/generated/grpc/go_load"
	"go.uber.org/zap"
)

type DownloadTask struct {
	ID             uint64                 `db:"id"`
	OfAccountID    uint64                 `db:"of_account_id"`
	DownloadType   go_load.DownloadType   `db:"download_type"`
	URL            string                 `db:"url"`
	DownloadStatus go_load.DownloadStatus `db:"download_status"`
	Metadata       string                 `db:"metadata"`
}

type DownloadTaskDataAccessor interface {
	CreateDownloadTask(ctx context.Context, task DownloadTask) (uint64, error)
	GetDownloadTaskListOfUser(ctx context.Context, userID, offset, limit uint64) ([]DownloadTask, error)
	GetDownloadTaskCountOfUser(ctx context.Context, userID uint64) (uint64, error)
	UpdateDownloadTask(ctx context.Context, task DownloadTask) error
	DeleteDownloadTask(ctx context.Context, id uint64) error
	WithDatabase(database Database) DownloadTaskDataAccessor
}

type downloadTaskDataAccessor struct {
	database Database
	logger   *zap.Logger
}

func NewDownloadTaskDataAccessor(
	database *goqu.Database,
	logger *zap.Logger,
) DownloadTaskDataAccessor {
	return &downloadTaskDataAccessor{
		database: database,
		logger:   logger,
	}
}

// CreateDownloadTask implements DownloadTaskDataAccessor.
func (d *downloadTaskDataAccessor) CreateDownloadTask(ctx context.Context, task DownloadTask) (uint64, error) {
	panic("unimplemented")
}

// DeleteDownloadTask implements DownloadTaskDataAccessor.
func (d *downloadTaskDataAccessor) DeleteDownloadTask(ctx context.Context, id uint64) error {
	panic("unimplemented")
}

// GetDownloadTaskCountOfUser implements DownloadTaskDataAccessor.
func (d *downloadTaskDataAccessor) GetDownloadTaskCountOfUser(ctx context.Context, userID uint64) (uint64, error) {
	panic("unimplemented")
}

// GetDownloadTaskListOfUser implements DownloadTaskDataAccessor.
func (d *downloadTaskDataAccessor) GetDownloadTaskListOfUser(ctx context.Context, userID uint64, offset uint64, limit uint64) ([]DownloadTask, error) {
	panic("unimplemented")
}

// UpdateDownloadTask implements DownloadTaskDataAccessor.
func (d *downloadTaskDataAccessor) UpdateDownloadTask(ctx context.Context, task DownloadTask) error {
	panic("unimplemented")
}

// WithDatabase implements DownloadTaskDataAccessor.
func (d *downloadTaskDataAccessor) WithDatabase(database Database) DownloadTaskDataAccessor {
	panic("unimplemented")
}
