package config

import (
	"github.com/kelseyhightower/envconfig"
)

const EnvPrefix = "Xls"

type Config struct {
	Port       int    `envconfig:"SERVER_PORT" default:"8080"`
	DevMode    bool   `envconfig:"DEV_MODE"`
	Index      string `envconfig:"INDEX" default:"test"`
	LogLevel  string `envconfig:"LOG_LEVEL" default:"debug"`
	SentryDSN string `envconfig:"SENTRY_DSN"`
	PathToFiles string `envconfig:"FILES_PATH"`
	PathToBackupFolder string `envconfig:"FOLDER_PATH"`
	PathToClearDir string `envconfig:"FOLDER_CLEAN_PATH"`
}

func NewConfig() (*Config, error) {
	conf := &Config{}
	err := conf.overrideWithEnvVars()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func (c *Config) overrideWithEnvVars() error {
	return envconfig.Process(EnvPrefix, c)
}
