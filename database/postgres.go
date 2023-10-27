package database

import (
	"database/sql"
	"time"

	"github.com/cool-service-go/config"
	_ "github.com/lib/pq"
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

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	db = conn
	return db, nil
}
