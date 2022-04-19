package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func openDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Print("DBConnectionError:", err)
	}
	return db
}

func getManufacturers() {
	db := openDatabase()
	rows, err := db.Query(`SELECT name FROM manufacturer`)
	if err != nil {
		return nil, err
	}
}
