package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	Address     string `env:"ADDRESS" envDefault:"localhost"`
	Port        int    `env:"PORT" envDefault:"8000"`
	DataBaseURL string `env:"DATABASE_URL_GO"`
}

func LoadFromEnvironment() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return cfg, nil
}
