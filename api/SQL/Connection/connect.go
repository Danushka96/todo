package Connection

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB = nil
var err error = nil

func Initialize(path string) {
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
}

func GetInstance() *sql.DB {
	if db == nil {
		db, _ = sql.Open("sqlite3", "./database.db")
	}
	return db
}
