package grpc

import (
	"context"
	"errors"
	"io"

	"github.com/nguyenhoang711/downloader/internal/configs"
	go_load "github.com/nguyenhoang711/downloader/internal/generated/go_load/v1"
	"github.com/nguyenhoang711/downloader/internal/logic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	//nolint:gosec // This is just to specify the metadata name
	AuthTokenMetadataName = "GOLOAD_AUTH"
)

type Handler struct {
	go_load.UnimplementedGoLoadServiceServer
	accountLogic                                 logic.Account
	downloadTaskLogic                            logic.DownloadTaskLogic
	getDownloadTaskFileResponseBufferSizeInBytes uint64
}

func NewHandler(
	accountLogic logic.Account,
	downloadTaskLogic logic.DownloadTaskLogic,
	grpcConfig configs.GRPC,
) (go_load.GoLoadServiceServer, error) {
	getDownloadTaskFileResponseBufferSizeBytes, err := grpcConfig.GetDownloadTaskFile.GetResponseBufferSizeInBytes()
	if err != nil {
		return nil, err
	}
	return &Handler{
		accountLogic:      accountLogic,
		downloadTaskLogic: downloadTaskLogic,
		getDownloadTaskFileResponseBufferSizeInBytes: getDownloadTaskFileResponseBufferSizeBytes,
	}, nil
}

func (a Handler) getAuthTokenMetadata(ctx context.Context) string {
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	metadataValues := metadata.Get(AuthTokenMetadataName)
	if len(metadataValues) == 0 {
		return ""
	}

	return metadataValues[0]
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
func (a Handler) CreateDownloadTask(
	ctx context.Context,
	request *go_load.CreateDownloadTaskRequest,
) (*go_load.CreateDownloadTaskResponse, error) {
	output, err := a.downloadTaskLogic.CreateDownloadTask(ctx, logic.CreateDownloadTaskParams{
		Token:        a.getAuthTokenMetadata(ctx),
		DownloadType: request.GetDownloadType(),
		URL:          request.GetUrl(),
	})
	if err != nil {
		return nil, err
	}

	return &go_load.CreateDownloadTaskResponse{
		DownloadTask: output.DownloadTask,
	}, nil
}

// CreateSession implements go_load.GoLoadServiceServer.
func (a Handler) CreateSession(
	ctx context.Context,
	request *go_load.CreateSessionRequest,
) (*go_load.CreateSessionResponse, error) {
	output, err := a.accountLogic.CreateSession(ctx, logic.CreateSessionParams{
		AccountName: request.GetAccountName(),
		Password:    request.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	err = grpc.SetHeader(ctx, metadata.Pairs(AuthTokenMetadataName, output.Token))
	if err != nil {
		return nil, err
	}

	return &go_load.CreateSessionResponse{
		Account: output.Account,
	}, nil
}

// DeleteDownloadTask implements go_load.GoLoadServiceServer.
func (a Handler) DeleteDownloadTask(
	ctx context.Context,
	request *go_load.DeleteDownloadTaskRequest,
) (*go_load.DeleteDownloadTaskResponse, error) {
	if err := a.downloadTaskLogic.DeleteDownloadTask(ctx, logic.DeleteDownloadTaskParams{
		Token:          a.getAuthTokenMetadata(ctx),
		DownloadTaskID: request.GetDownloadTaskId(),
	}); err != nil {
		return nil, err
	}

	return &go_load.DeleteDownloadTaskResponse{}, nil
}

// GetDownloadTaskFile implements go_load.GoLoadServiceServer.
func (a Handler) GetDownloadTaskFile(
	request *go_load.GetDownloadTaskFileRequest,
	server go_load.GoLoadService_GetDownloadTaskFileServer,
) error {
	outputReader, err := a.downloadTaskLogic.GetDownloadTaskFile(server.Context(), logic.GetDownloadTaskFileParams{
		Token:          a.getAuthTokenMetadata(server.Context()),
		DownloadTaskID: request.GetDownloadTaskId(),
	})
	if err != nil {
		return err
	}

	defer outputReader.Close()

	for {
		dataBuffer := make([]byte, a.getDownloadTaskFileResponseBufferSizeInBytes)
		readByteCount, readErr := outputReader.Read(dataBuffer)

		if readByteCount > 0 {
			sendErr := server.Send(&go_load.GetDownloadTaskFileResponse{
				Data: dataBuffer[:readByteCount],
			})
			if sendErr != nil {
				return sendErr
			}

			continue
		}
		if readErr != nil {
			if errors.Is(readErr, io.EOF) {
				break
			}

			return readErr
		}
	}

	return nil
}

// GetDownloadTaskList implements go_load.GoLoadServiceServer.
func (a *Handler) GetDownloadTaskList(
	ctx context.Context,
	getListDownloadReq *go_load.GetDownloadTaskListRequest,
) (*go_load.GetDownloadTaskListResponse, error) {
	output, err := a.downloadTaskLogic.GetDownloadTaskList(ctx, logic.GetDownloadTaskListParams{
		Token:  a.getAuthTokenMetadata(ctx),
		Offset: getListDownloadReq.GetOffset(),
		Limit:  getListDownloadReq.GetLimit(),
	})
	if err != nil {
		return nil, err
	}

	return &go_load.GetDownloadTaskListResponse{
		DownloadTaskList:       output.DownloadTaskList,
		TotalDownloadTaskCount: output.TotalDownloadTaskCount,
	}, nil
}

// UpdateDownloadTask implements go_load.GoLoadServiceServer.
func (a *Handler) UpdateDownloadTask(
	ctx context.Context,
	request *go_load.UpdateDownloadTaskRequest,
) (*go_load.UpdateDownloadTaskResponse, error) {
	output, err := a.downloadTaskLogic.UpdateDownloadTask(ctx, logic.UpdateDownloadTaskParams{
		Token:          a.getAuthTokenMetadata(ctx),
		DownloadTaskID: request.GetDownloadTaskId(),
		URL:            request.GetUrl(),
	})
	if err != nil {
		return nil, err
	}

	return &go_load.UpdateDownloadTaskResponse{
		DownloadTask: output.UpdatedDownloadTask,
	}, nil
}

// mustEmbedUnimplementedGoLoadServiceServer implements go_load.GoLoadServiceServer.
