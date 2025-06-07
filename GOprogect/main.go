package main

import (
	"GOprogect/db"
	"context"
)

func main() {
	confFile := "config.yml"
	ctx := context.Background()
	db.ProcessingPort(ctx, confFile)
}
