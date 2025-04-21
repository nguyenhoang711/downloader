package utils

import "github.com/google/wire"

// wireset config for wire package
var WireSet = wire.NewSet(
	InitializeLogger,
)
