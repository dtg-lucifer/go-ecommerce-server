package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewSQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal("Error: Opening the database", err)
	}

	return db, nil
}
