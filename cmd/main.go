package main

import (
	"log/slog"
	"os"
)

// used to compose the dependencies
func main() {

	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	// structured logging - a very powerful logging solution
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	// Mounting the API - ie setting middlewares and routing
	h := api.mount()

	// Running the API
	if err := api.run(h); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
