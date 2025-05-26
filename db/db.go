package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func GetConnection() *sql.DB {
	open, err := sql.Open("sqlite3", "store.sqlite")
	if err != nil {
		log.Println("Could not open database")
	}
	return open
}
