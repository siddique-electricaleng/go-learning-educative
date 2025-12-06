package main

import (
	"context"
	"ecom/internal/env"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

// used to compose the dependencies
func main() {
	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost port=5432 user=postgres password=Localpg12W1 dbname=products sslmode=disable"),
		},
	}

	// structured logging - a very powerful logging solution
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("Connected to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     conn,
	}

	// Mounting the API - ie setting middlewares and routing
	h := api.mount()

	// Running the API
	if err := api.run(h); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
