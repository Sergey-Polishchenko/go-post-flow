package config

import (
	"fmt"
	"sync"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string `env:"PORT" envDefault:"8080"`
}

var (
	cfg    *Config
	cfgErr error
	once   sync.Once
)

func loadConfig() {
	if err := godotenv.Load(); err != nil {
		cfgErr = fmt.Errorf("failed to load .env file: %v", err)
	}
	cfg = &Config{}
	if err := env.Parse(cfg); err != nil {
		cfgErr = fmt.Errorf("failed to parse environment variables: %v", err)
	}
}

func GetConfig() (*Config, error) {
	once.Do(loadConfig)
	return cfg, cfgErr
}
