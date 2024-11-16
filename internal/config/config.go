package config

import (
	"fmt"

	"github.com/kkyr/fig"
)

var (
	config cfg
)

type cfg struct {
	Port             string `fig:"port"`
	ConnectionString string `fig:"connection_string"`
}

func Parse() error {
	err := fig.Load(&config, fig.File("config/config.yaml"))
	if err != nil {
		return fmt.Errorf("an error occurred when trying to parse config file %w", err)
	}
	return nil
}

func Get() cfg {
	return config
}
