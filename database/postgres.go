package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/mhshahin/cool-service-go/config"
)

var db *sql.DB

func InitDB(cfg *config.AppConfig) (*sql.DB, error) {
	dbURL := cfg.DBConnectionString()
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(cfg.Postgres.MaxOpenConnections)
	conn.SetMaxIdleConns(cfg.Postgres.MaxIdleConnections)

	db = conn
	return db, nil
}
