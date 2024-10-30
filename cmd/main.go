package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/dtg-lucifer/go-backend/cmd/api"
	"github.com/dtg-lucifer/go-backend/config"
	"github.com/dtg-lucifer/go-backend/db"
)

func main() {
	db, err := db.NewSQLStorage(mysql.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPassword,
		Addr:                 config.Env.DBAddress,
		DBName:               config.Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal("Error: Connecting to the database", err)
	}

	initStorage(db)

	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal("Error: Running the server", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Error: Connecting to the database", err)
	}

	log.Println("DB: Connected to the database")
}
