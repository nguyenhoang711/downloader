//go:build wireinject
// +build wireinject

package wiring

import (
	"github.com/google/wire"
	"github.com/nguyenhoang711/downloader/internal/app"
	"github.com/nguyenhoang711/downloader/internal/configs"
	"github.com/nguyenhoang711/downloader/internal/dataaccess"
	"github.com/nguyenhoang711/downloader/internal/handler"
	"github.com/nguyenhoang711/downloader/internal/logic"
	"github.com/nguyenhoang711/downloader/internal/utils"
)

var wireSet = wire.NewSet(
	configs.WireSet,
	dataaccess.WireSet,
	handler.WireSet,
	logic.WireSet,
	utils.WireSet,
	app.WireSet,
)

func InitializeStandaloneServer(configFilePath configs.ConfigFilePath) (*app.StandaloneServer, func(), error) {
	wire.Build(wireSet)

	return nil, nil, nil
}
