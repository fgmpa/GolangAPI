package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	carTable = "cars"
)

type Config struct {
	Host    string
	Port    string
	DBName  string
	SSLMode string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.DBName, "postgres", "1234", cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
