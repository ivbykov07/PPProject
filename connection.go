package main

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"go.uber.org/fx"
)

func NewConnection(lc fx.Lifecycle, cfg *Config) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.DBHost, cfg.DBPort),
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Database: cfg.DBName,
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.Ping(ctx)
		},
		OnStop: func(context.Context) error {
			return db.Close()
		},
	})
	return db
}
