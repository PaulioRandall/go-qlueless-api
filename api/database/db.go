package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// db is the applications shared database.
var db *sql.DB = nil

// Get returns the applications shared database.
func Get() *sql.DB {
	return db
}

// IsOpen returns true if the shared database is open.
func IsOpen() bool {
	if db == nil {
		return false
	}
	return true
}

// Open opens the database if not already open.
func Open() *sql.DB {
	if !IsOpen() {
		var err error
		db, err = openSQLiteDatabase("./qlueless.db")
		if err != nil {
			panic(err)
		}
	}

	return db
}

// Close closes the database if open.
func Close() {
	if IsOpen() {
		err := db.Close()
		db = nil

		if err != nil {
			panic(err)
		}
	}
}

// openSQLiteDatabase opens a SQLite database, creating it if it doesn't already
// exist.
func openSQLiteDatabase(path string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", path)

	if err != nil {
		if db != nil {
			db.Close()
		}
	}

	return
}
