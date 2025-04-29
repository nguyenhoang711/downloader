package jobs

import (
	"context"

	"github.com/nguyenhoang711/downloader/internal/logic"
)

type ExecuteAllPendingDownloadTask interface {
	Run(context.Context) error
}

type executeAllPendingDownloadTask struct {
	downloadTaskLogic logic.DownloadTaskLogic
}

func NewExecuteAllPendingDownloadTask(
	downloadTaskLogic logic.DownloadTaskLogic,
) ExecuteAllPendingDownloadTask {
	return &executeAllPendingDownloadTask{
		downloadTaskLogic: downloadTaskLogic,
	}
}

func (e executeAllPendingDownloadTask) Run(ctx context.Context) error {
	return e.downloadTaskLogic.ExecuteAllPendingDownloadTask(ctx)
}
