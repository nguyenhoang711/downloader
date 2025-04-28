package logic

import (
	"context"
	"io"
	"net/http"

	"github.com/nguyenhoang711/downloader/internal/utils"
	"go.uber.org/zap"
)

const (
	HTTPResponseHeaderContentType = "Content-Type"
	HTTPMetadataKeyContentType    = "content-type"
)

type Downloader interface {
	Download(ctx context.Context, writer io.Writer) (map[string]any, error)
}

type HTTPDownloader struct {
	url    string
	logger *zap.Logger
}

func NewHTTPDownloader(
	url string,
	logger *zap.Logger,
) Downloader {
	return &HTTPDownloader{
		url:    url,
		logger: logger,
	}
}

func (h HTTPDownloader) Download(ctx context.Context, writer io.Writer) (map[string]any, error) {
	logger := utils.LoggerWithContext(ctx, h.logger)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, h.url, http.NoBody)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to create http get request")
		return nil, err
	}

	// thuc hien request download file
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to make http get request")
		return nil, err
	}

	defer response.Body.Close()

	// copy tu bo doc sang bo ghi
	_, err = io.Copy(writer, response.Body)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to read response and write to writer")
		return nil, err
	}

	metadata := map[string]any{
		HTTPMetadataKeyContentType: response.Header.Get(HTTPResponseHeaderContentType),
	}

	return metadata, nil
}
