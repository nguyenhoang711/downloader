package app

import (
	"context"
	"syscall"

	"github.com/go-co-op/gocron/v2"
	"github.com/nguyenhoang711/downloader/internal/configs"
	"github.com/nguyenhoang711/downloader/internal/handler/consumers"
	"github.com/nguyenhoang711/downloader/internal/handler/grpc"
	"github.com/nguyenhoang711/downloader/internal/handler/http"
	"github.com/nguyenhoang711/downloader/internal/handler/jobs"
	"github.com/nguyenhoang711/downloader/internal/utils"
	"go.uber.org/zap"
)

type StandaloneServer struct {
	grpcServer                                               grpc.Server
	httpServer                                               http.Server
	rootConsumer                                             consumers.RootConsumer
	executeAllPendingDownloadTaskJob                         jobs.ExecuteAllPendingDownloadTask
	updateDownloadingAndFailedDownloadTaskStatusToPendingJob jobs.UpdateDownloadingAndFailedDownloadTaskStatusToPending
	logger                                                   *zap.Logger
	cronConfig                                               configs.Cron
}

func NewStandaloneServer(
	grpcServer grpc.Server,
	httpServer http.Server,
	rootConsumer consumers.RootConsumer,
	executeAllPendingDownloadTaskJob jobs.ExecuteAllPendingDownloadTask,
	updateDownloadingAndFailedDownloadTaskStatusToPendingJob jobs.UpdateDownloadingAndFailedDownloadTaskStatusToPending,
	logger *zap.Logger,
	cronConfig configs.Cron,
) *StandaloneServer {
	return &StandaloneServer{
		grpcServer:                       grpcServer,
		httpServer:                       httpServer,
		rootConsumer:                     rootConsumer,
		executeAllPendingDownloadTaskJob: executeAllPendingDownloadTaskJob,
		updateDownloadingAndFailedDownloadTaskStatusToPendingJob: updateDownloadingAndFailedDownloadTaskStatusToPendingJob,
		logger:     logger,
		cronConfig: cronConfig,
	}
}

func (s StandaloneServer) scheduleCronJobs(scheduler gocron.Scheduler) error {
	if _, err := scheduler.NewJob(
		gocron.CronJob(s.cronConfig.ExecuteAllPendingDownloadTask.Schedule, true),
		gocron.NewTask(func() {
			if err := s.executeAllPendingDownloadTaskJob.Run(context.Background()); err != nil {
				s.logger.With(zap.Error(err)).Error("failed to run execute all pending download task job")
			}
		}),
	); err != nil {
		s.logger.With(zap.Error(err)).Error("failed to schedule execute all pending download task job")
		return err
	}

	if _, err := scheduler.NewJob(
		gocron.CronJob(s.cronConfig.UpdateDownloadingAndFailedDownloadTaskStatusToPending.Schedule, true),
		gocron.NewTask(func() {
			if err := s.executeAllPendingDownloadTaskJob.Run(context.Background()); err != nil {
				s.logger.With(zap.Error(err)).
					Error("failed to run update downloading and failed download task status to pending job")
			}
		}),
	); err != nil {
		s.logger.With(zap.Error(err)).
			Error("failed to schedule update downloading and failed download task status to pending job")
		return err
	}

	return nil
}

func (s StandaloneServer) Start() error {
	// reset trạng thái các tasks failed hoặc downloading về pending
	if err := s.updateDownloadingAndFailedDownloadTaskStatusToPendingJob.Run(context.Background()); err != nil {
		return err
	}

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		s.logger.With(zap.Error(err)).Error("failed to initialize scheduler")
		return err
	}

	defer func() {
		if shutdownErr := scheduler.Shutdown(); shutdownErr != nil {
			s.logger.With(zap.Error(shutdownErr)).Error("failed to shutdown scheduler")
		}
	}()

	err = s.scheduleCronJobs(scheduler)
	if err != nil {
		return err
	}

	go func() {
		grpcStartErr := s.grpcServer.Start(context.Background())
		s.logger.With(zap.Error(grpcStartErr)).Info("grpc server stopped")
	}()

	go func() {
		httpStartErr := s.httpServer.Start(context.Background())
		s.logger.With(zap.Error(httpStartErr)).Info("http server stopped")
	}()

	go func() {
		consumerStartErr := s.rootConsumer.Start(context.Background())
		s.logger.With(zap.Error(consumerStartErr)).Info("message queue consumer stopped")
	}()

	utils.BlockUntilSignal(syscall.SIGINT, syscall.SIGTERM)
	return nil
}
