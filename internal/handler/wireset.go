package handler

import (
	"github.com/google/wire"
	"github.com/nguyenhoang711/downloader/internal/handler/grpc"
	"github.com/nguyenhoang711/downloader/internal/handler/http"
)

var WireSet = wire.NewSet(
	http.WireSet,
	grpc.WireSet,
)
