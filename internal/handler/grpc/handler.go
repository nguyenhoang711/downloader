package grpc

import (
	"context"

	"github.com/nguyenhoang711/downloader/internal/generated/grpc/go_load"
	"github.com/nguyenhoang711/downloader/internal/logic"
)

type Handler struct {
	go_load.UnimplementedGoLoadServiceServer
	accountLogic logic.Account
}

func NewHandler(
	accountLogic logic.Account,
) go_load.GoLoadServiceServer {
	return &Handler{
		accountLogic: accountLogic,
	}
}

// CreateAccount implements go_load.GoLoadServiceServer.
func (a Handler) CreateAccount(
	ctx context.Context,
	request *go_load.CreateAccountRequest,
) (*go_load.CreateAccountResponse, error) {
	output, err := a.accountLogic.CreateAccount(ctx, logic.CreateAccountParams{
		AccountName: request.GetAccountName(),
		Password:    request.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &go_load.CreateAccountResponse{
		AccountId: output.ID,
	}, nil
}

// CreateDownloadTask implements go_load.GoLoadServiceServer.
func (a *Handler) CreateDownloadTask(
	ctx context.Context,
	request *go_load.CreateDownloadTaskRequest,
) (*go_load.CreateDownloadTaskResponse, error) {
	panic("unimplemented")
}

// CreateSession implements go_load.GoLoadServiceServer.
func (a *Handler) CreateSession(
	ctx context.Context,
	request *go_load.CreateSessionRequest,
) (*go_load.CreateSessionResponse, error) {
	panic("unimplemented")
}

// DeleteDownloadTask implements go_load.GoLoadServiceServer.
func (a *Handler) DeleteDownloadTask(ctx context.Context,
	delDownloadTaskReq *go_load.DeleteDownloadTaskRequest,
) (*go_load.DeleteDownloadTaskResponse, error) {
	panic("unimplemented")
}

// GetDownloadTaskFile implements go_load.GoLoadServiceServer.
func (a *Handler) GetDownloadTaskFile(
	*go_load.GetDownloadTaskFileRequest,
	go_load.GoLoadService_GetDownloadTaskFileServer,
) error {
	panic("unimplemented")
}

// GetDownloadTaskList implements go_load.GoLoadServiceServer.
func (a *Handler) GetDownloadTaskList(
	ctx context.Context,
	getListDownloadReq *go_load.GetDownloadTaskListRequest,
) (*go_load.GetDownloadTaskListResponse, error) {
	panic("unimplemented")
}

// UpdateDownloadTask implements go_load.GoLoadServiceServer.
func (a *Handler) UpdateDownloadTask(
	ctx context.Context,
	updateDownloadTask *go_load.UpdateDownloadTaskRequest,
) (*go_load.UpdateDownloadTaskResponse, error) {
	panic("unimplemented")
}

// mustEmbedUnimplementedGoLoadServiceServer implements go_load.GoLoadServiceServer.
