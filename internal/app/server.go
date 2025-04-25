package app

import (
	"context"
	"syscall"

	"github.com/nguyenhoang711/downloader/internal/dataaccess/database"
	"github.com/nguyenhoang711/downloader/internal/handler/consumers"
	"github.com/nguyenhoang711/downloader/internal/handler/grpc"
	"github.com/nguyenhoang711/downloader/internal/handler/http"
	"github.com/nguyenhoang711/downloader/internal/utils"
	"go.uber.org/zap"
)

type Server struct {
	grpcServer       grpc.Server
	httpServer       http.Server
	rootConsumer     consumers.RootConsumer
	logger           *zap.Logger
	databaseMigrator database.Migrator
}

func NewServer(
	grpcServer grpc.Server,
	httpServer http.Server,
	rootConsumer consumers.RootConsumer,
	databaseMigrator database.Migrator,
	logger *zap.Logger,
) *Server {
	return &Server{
		grpcServer:       grpcServer,
		httpServer:       httpServer,
		rootConsumer:     rootConsumer,
		logger:           logger,
		databaseMigrator: databaseMigrator,
	}
}

func (s Server) Start() error {
	if err := s.databaseMigrator.Up(context.Background()); err != nil {
		s.logger.With(zap.Error(err)).Error("failed to execute database up migration")
		return err
	}
	go func() {
		err := s.grpcServer.Start(context.Background())
		s.logger.With(zap.Error(err)).Info("grpc server stopped")
	}()

	go func() {
		err := s.httpServer.Start(context.Background())
		s.logger.With(zap.Error(err)).Info("http server stopped")
	}()

	go func() {
		err := s.rootConsumer.Start(context.Background())
		s.logger.With(zap.Error(err)).Info("message queue consumer stopped")
	}()

	utils.BlockUntilSignal(syscall.SIGINT, syscall.SIGTERM)
	return nil
}
