package db

import (
	"GOprogect/confige"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func NewPort(ctx context.Context, attempts int, portnames confige.PortConfig) (pool *pgxpool.Pool, err error) {
	port := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		portnames.User,
		portnames.Password,
		portnames.Host,
		portnames.Port,
		portnames.DBName,
		portnames.SSlMode)
	err = TryAttempt(func() error {
		ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
		defer cancel()
		pool, err = pgxpool.New(ctx, port)
		if err != nil {
			return err
		}

		if err := pool.Ping(ctx); err != nil {
			log.Print(err)
			return err
		}
		return nil
	}, attempts, 5)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
