package app

import (
	"context"
	"syscall"

	"github.com/nguyenhoang711/downloader/internal/dataaccess/database"
	"github.com/nguyenhoang711/downloader/internal/handler/grpc"
	"github.com/nguyenhoang711/downloader/internal/handler/http"
	"github.com/nguyenhoang711/downloader/internal/utils"
	"go.uber.org/zap"
)

type Server struct {
	databaseMigrator database.Migrator
	grpcServer       grpc.Server
	httpServer       http.Server
	logger           *zap.Logger
}

func NewServer(
	grpcServer grpc.Server,
	httpServer http.Server,
	logger *zap.Logger,
	migrator database.Migrator,
) *Server {
	return &Server{
		grpcServer:       grpcServer,
		httpServer:       httpServer,
		logger:           logger,
		databaseMigrator: migrator,
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

	utils.BlockUntilSignal(syscall.SIGINT, syscall.SIGTERM)
	return nil
}
