package db

import (
	"GOprogect/confige"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Connection struct{ pool *pgxpool.Pool }

func TryAttempt(f func() error, attempts int, sec time.Duration) error {
	for attempt := 0; attempt < attempts; attempt++ {
		err := f()
		if err == nil {
			log.Println("Connected to port")
			return nil
		}
		log.Printf("%d attempts remaining", attempt+1)
		if attempt < attempts-1 {
			time.Sleep(sec * time.Second)
		}
	}
	return errors.New("attempt exceeded")
}

func ProcessingPort(ctx context.Context, config string) Connection {
	data, err := os.ReadFile(config)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to read config file: %w", err))
	}
	var conf confige.PortConfig
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatal(fmt.Printf("failed to unmarshal config file: %w", err))
	}
	Pool, err := NewPort(ctx, 3, conf)
	if err != nil {
		log.Fatal(fmt.Printf("failed to connect to database: %w", err))
	}
	fmt.Printf("Connected to port %d\n", conf.Port)
	return Connection{pool: Pool}
}
