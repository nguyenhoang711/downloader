package handler

import (
	"github.com/google/wire"
	"github.com/nguyenhoang711/downloader/internal/handler/consumers"
	"github.com/nguyenhoang711/downloader/internal/handler/grpc"
	"github.com/nguyenhoang711/downloader/internal/handler/http"
	"github.com/nguyenhoang711/downloader/internal/handler/jobs"
)

var WireSet = wire.NewSet(
	http.WireSet,
	grpc.WireSet,
	consumers.WireSet,
	jobs.WireSet,
)
