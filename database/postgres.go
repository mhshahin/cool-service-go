package database

import (
	"database/sql"

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

	db = conn
	return db, nil
}
