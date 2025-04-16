package dataaccess

import (
	"github.com/google/wire"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/database"
)

var WireSet = wire.NewSet(
	database.WireSet,
)
