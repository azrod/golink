package config

import (
	"github.com/num30/config"

	"github.com/azrod/golink/pkg/sb"
)

type (
	Config struct {
		App     App
		Health  Health
		Storage sb.Config
	}

	App struct {
		Address string `default:"localhost"`
		Port    int    `default:"8081"`
	}

	Health struct {
		Address string `default:"localhost"`
		Port    int    `default:"8082"`
	}
)

// Read reads the config from the config file, environment variables and flags.
func Read() (cfg Config, err error) {
	cfgReader := config.
		NewConfReader("config").
		WithSearchDirs("/etc/golink", ".").
		WithPrefix("GOLINK")

	if err := cfgReader.Read(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
