package configs

import (
	"fmt"
	"os"

	"github.com/nguyenhoang711/downloader/configs"
	"gopkg.in/yaml.v2"
)

type ConfigFilePath string

type Config struct {
	GRPC     GRPC     `yaml:"grpc"`
	HTTP     HTTP     `yaml:"http"`
	Log      Log      `yaml:"log"`
	Auth     Auth     `yaml:"auth"`
	Database Database `yaml:"database"`
	Cache    Cache    `yaml:"cache"`
	MQ       MQ       `yaml:"mq"`
	Download Download `yaml:"download"`
	Cron     Cron     `yaml:"cron"`
}

func NewConfig(filePath ConfigFilePath) (Config, error) {
	var (
		configBytes = configs.DefaultConfigBytes
		config      = Config{}
		err         error
	)

	if filePath != "" {
		configBytes, err = os.ReadFile(string(filePath))
		if err != nil {
			return Config{}, fmt.Errorf("failed to read YAML file: %w", err)
		}
	}
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return config, nil
}
