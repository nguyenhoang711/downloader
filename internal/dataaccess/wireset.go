package dataaccess

import (
	"github.com/google/wire"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/cache"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/database"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/mq"
)

var WireSet = wire.NewSet(
	database.WireSet,
	cache.WireSet,
	mq.WireSet,
)
