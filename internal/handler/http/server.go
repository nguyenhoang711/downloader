package http

import (
	"context"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nguyenhoang711/downloader/internal/configs"
	"github.com/nguyenhoang711/downloader/internal/generated/grpc/go_load"
	handlerGRPC "github.com/nguyenhoang711/downloader/internal/handler/grpc"
	"github.com/nguyenhoang711/downloader/internal/handler/http/servemuxoptions"
	"github.com/nguyenhoang711/downloader/internal/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	//nolint:gosec // This is just to specify the cookie name
	AuthTokenCookieName = "GOLOAD_AUTH"
)

type Server interface {
	Start(ctx context.Context) error
}

type server struct {
	grpcConfig configs.GRPC
	httpConfig configs.HTTP
	authConfig configs.Auth
	logger     *zap.Logger
}

func NewServer(
	grpcConfig configs.GRPC,
	httpConfig configs.HTTP,
	authConfig configs.Auth,
	logger *zap.Logger,
) Server {
	return &server{
		grpcConfig: grpcConfig,
		httpConfig: httpConfig,
		authConfig: authConfig,
		logger:     logger,
	}
}

func (s *server) Start(ctx context.Context) error {
	logger := utils.LoggerWithContext(ctx, s.logger)
	tokenExpiresInDuration, err := s.authConfig.Token.GetExpiresInDuration()
	if err != nil {
		return err
	}

	grpcMux := runtime.NewServeMux(
		servemuxoptions.WithAuthCookieToAuthMetadata(AuthTokenCookieName, handlerGRPC.AuthTokenMetadataName),
		servemuxoptions.WithAuthMetadataToAuthCookie(
			handlerGRPC.AuthTokenMetadataName, AuthTokenCookieName, tokenExpiresInDuration),
	)
	err = go_load.RegisterGoLoadServiceHandlerFromEndpoint(
		ctx,
		grpcMux,
		s.grpcConfig.Address, // gRPC server address forward from HTTP
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		})
	if err != nil {
		return err
	}
	httpServer := http.Server{
		Addr:              s.httpConfig.Address,
		ReadHeaderTimeout: time.Minute,
		Handler:           grpcMux,
	}

	logger.With(zap.String("address", s.httpConfig.Address)).Info("starting http server")
	return httpServer.ListenAndServe()
}
