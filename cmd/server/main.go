package main

import (
    "context"
    "log"

    "internal/config"
    "internal/database"
    "internal/workers"
)

func main() {
    cfg := config.Load()

    db, err := database.NewMySQL(cfg.Database)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    ctx := context.Background()

    pool := workers.NewPool(
        db,
        cfg.Workers,
    )

    if err := pool.Run(ctx); err != nil {
        log.Fatal(err)
    }
}
