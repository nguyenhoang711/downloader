//go:build wireinject
// +build wireinject

package wiring

import (
	"github.com/google/wire"
	"github.com/nguyenhoang711/downloader/internal/configs"
	"github.com/nguyenhoang711/downloader/internal/dataaccess"
	"github.com/nguyenhoang711/downloader/internal/handler"
	"github.com/nguyenhoang711/downloader/internal/handler/grpc"
	"github.com/nguyenhoang711/downloader/internal/logic"
)

var wireSet = wire.NewSet(
	configs.WireSet,
	dataaccess.WireSet,
	handler.WireSet,
	logic.WireSet,
)

func InitializeGRPCServer(configFilePath configs.ConfigFilePath) (grpc.Server, func(), error) {
	wire.Build(wireSet)

	return nil, nil, nil
}
