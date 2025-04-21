package mq

import (
	"github.com/google/wire"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/mq/producer"
)

var WireSet = wire.NewSet(
	producer.WireSet,
)
