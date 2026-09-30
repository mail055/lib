package tui

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func initDB() {
	var err error

	// Establishes the databases connection
	db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	// Tests the database connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}
