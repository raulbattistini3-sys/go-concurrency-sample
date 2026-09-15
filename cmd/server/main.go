package main

import (
	"context"
	"log"
	"net/http"

	"go-concurrency-sample/internal/api"
	"go-concurrency-sample/internal/config"
	"go-concurrency-sample/internal/database"
	"go-concurrency-sample/internal/workers"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	cfg := config.Get()

	db, err := database.NewMySQL(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := database.NewUserRepository(db)

	pool := workers.NewPool(
		userRepo,
		cfg.Workers.Count,
		cfg.Workers.Queue,
	)

	ctx := context.Background()

	// Start workers.
	go pool.Run(ctx)

	handler := api.NewHandler(pool)

	mux := http.NewServeMux()

	mux.HandleFunc(
		"POST /users/generate",
		handler.GenerateUsers,
	)

	log.Printf("HTTP server listening on :5555")

	if err := http.ListenAndServe(":5555", mux); err != nil {
		log.Fatal(err)
	}
}
