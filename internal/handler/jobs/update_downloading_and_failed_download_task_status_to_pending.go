package jobs

import (
	"context"

	"github.com/nguyenhoang711/downloader/internal/logic"
)

type UpdateDownloadingAndFailedDownloadTaskStatusToPending interface {
	Run(context.Context) error
}

type updateDownloadingAndFailedDownloadTaskStatusToPending struct {
	downloadTaskLogic logic.DownloadTaskLogic
}

func NewUpdateDownloadingAndFailedDownloadTaskStatusToPending(
	downloadTaskLogic logic.DownloadTaskLogic,
) UpdateDownloadingAndFailedDownloadTaskStatusToPending {
	return &updateDownloadingAndFailedDownloadTaskStatusToPending{
		downloadTaskLogic: downloadTaskLogic,
	}
}

func (u updateDownloadingAndFailedDownloadTaskStatusToPending) Run(ctx context.Context) error {
	return u.downloadTaskLogic.UpdateDownloadingAndFailedDownloadTaskStatusToPending(ctx)
}
