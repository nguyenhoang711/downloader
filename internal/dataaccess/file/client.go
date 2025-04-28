package file

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/nguyenhoang711/downloader/internal/configs"
	"github.com/nguyenhoang711/downloader/internal/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client interface {
	Read(ctx context.Context, filePath string) (io.ReadCloser, error)
	Write(ctx context.Context, filePath string) (io.WriteCloser, error)
}

// wrapper struct read bucket (streaming)
type bufferedFileReader struct {
	file           *os.File
	bufferedReader io.Reader
}

func newBufferedFileReader(
	file *os.File,
) io.ReadCloser {
	return &bufferedFileReader{
		file:           file,
		bufferedReader: bufio.NewReader(file),
	}
}
func (b bufferedFileReader) Close() error {
	return b.file.Close()
}

func (b bufferedFileReader) Read(p []byte) (int, error) {
	return b.bufferedReader.Read(p)
}

type LocalClient struct {
	downloadDirectory string
	logger            *zap.Logger
}

func NewLocalClient(
	downloadConfig configs.Download,
	logger *zap.Logger,
) (Client, error) {
	if err := os.MkdirAll(downloadConfig.DownloadDirectory, os.ModeDir); err != nil {
		if !errors.Is(err, os.ErrExist) {
			return nil, fmt.Errorf("failed to create download directory: %w", err)
		}
	}

	return &LocalClient{
		downloadDirectory: downloadConfig.DownloadDirectory,
		logger:            logger,
	}, nil
}

func (l LocalClient) Read(ctx context.Context, filePath string) (io.ReadCloser, error) {
	logger := utils.LoggerWithContext(ctx, l.logger).With(zap.String("file_path", filePath))

	absolutePath := path.Join(l.downloadDirectory, filePath)
	file, err := os.Open(absolutePath)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to open file")
		return nil, status.Error(codes.Internal, "failed to open file")
	}

	return newBufferedFileReader(file), nil
}

func (l *LocalClient) Write(ctx context.Context, filePath string) (io.WriteCloser, error) {
	logger := utils.LoggerWithContext(ctx, l.logger).With(zap.String("file_path", filePath))

	absolutePath := path.Join(l.downloadDirectory, filePath)
	file, err := os.Create(absolutePath)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to open file")
		return nil, status.Error(codes.Internal, "failed to open file")
	}

	return file, nil
}
