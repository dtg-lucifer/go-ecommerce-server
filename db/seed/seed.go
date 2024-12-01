package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/go-sql-driver/mysql"

	"github.com/dtg-lucifer/go-backend/config"
	"github.com/dtg-lucifer/go-backend/db"
)

func main() {
	seedCount := 100

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

	fmt.Printf("Processing the seed of %d rows\n", seedCount)

	for n := range seedCount {
		_, err := db.Exec(
			"INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)",
			"seed_user_",
			strconv.Itoa(n),
			fmt.Sprintf("see_user_%d@gmail.com", n),
			"sUpER_Str@ng__PaSSwoRD",
		)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Error: Connecting to the database", err)
	}

	log.Println("DB: Connected to the database")
}
