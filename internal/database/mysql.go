package database

import (
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
)

type Config struct {
    DSN          string
    MaxOpenConns int
    MaxIdleConns int
}

func NewMySQL(cfg Config) (*sql.DB, error) {
    db, err := sql.Open("mysql", cfg.DSN)
    if err != nil {
        return nil, err
    }

    db.SetMaxOpenConns(cfg.MaxOpenConns)
    db.SetMaxIdleConns(cfg.MaxIdleConns)

    if err := db.Ping(); err != nil {
        db.Close()
        return nil, err
    }

    return db, nil
}
