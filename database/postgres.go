package database

import (
	"database/sql"
	"time"

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

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxIdleTime(10 * time.Minute)
	conn.SetConnMaxLifetime(30 * time.Minute)

	db = conn
	return db, nil
}
