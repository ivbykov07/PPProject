package main

import (
	"github.com/caarlos0/env"
)

type Config struct {
	Port int64 `env:"API_APP_PORT" envDefault:"3000"`

	DBHost        string `env:"API_PG_HOST" envDefault:"localhost"`
	DBPort        int64  `env:"API_PG_PORT" envDefault:"5432"`
	DBUser        string `env:"API_DB_USER" envDefault:"postgres"`
	DBPassword    string `env:"API_DB_PASSWORD" envDefault:"postgres"`
	DBName        string `env:"API_DB_NAME" envDefault:"postgres"`
	PrometheusURL string `env:"API_PROM_URL"`

	//JWTSecret string `env:"API_JWT_SECRET"`
	//NATSURL   string `env:"API_NATS_URL"`
	//RedisURL  string `env:"API_REDIS_URL"`
}

func NewConfig() (*Config, error) {
	cfg := new(Config)
	if err := cfg.Load(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) Load() error {
	if err := env.Parse(cfg); err != nil {
		return err
	}
	return nil
}
