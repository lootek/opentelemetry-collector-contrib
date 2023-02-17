package broker

import (
	"github.com/fhmq/hmq/broker"
)

type Config struct {
	*broker.Config
}

func DefaultConfig() *Config {
	cfg := &Config{
		broker.DefaultConfig,
	}

	cfg.Debug = true

	return cfg
}
