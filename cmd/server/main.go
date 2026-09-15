package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go-concurrency-sample/internal/config"
	"go-concurrency-sample/internal/database"
	"go-concurrency-sample/internal/workers"
)

func main() {

  if err := config.Load(); err != nil { 
    log.Fatal(err) 
  } 
  cfg := config.Get() 
  
  log.Printf( "starting server with %d workers", cfg.Workers.Count, ) 
  log.Printf( "database: %s:%s/%s", cfg.Database.Host, cfg.Database.Port, cfg.Database.Name, )
  userRepo := database.NewUserRepository(config.GetDatabase())
    
  if userRepo == nil {
      err := errors.New("error user repo nil")
      log.Fatal(fmt.Sprintf("error user repo nil %w", err))
  }

    ctx := context.Background()

    pool := workers.NewPool(
        userRepo,
        cfg.Workers,
    )

    if err := pool.Run(ctx); err != nil {
        log.Fatal(err)
    }
}
