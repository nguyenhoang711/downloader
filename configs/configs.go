package configs

import (
	_ "embed"
)

// name of default config file is local.yaml
//
//go:embed local.yaml
var DefaultConfigBytes []byte
