package config

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenDB(dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}
