package std

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB is the applications shared database.
var DB *Database = &Database{}

// Database represents the server resources.
type Database struct {
	SQL    *sql.DB
	isOpen bool
}

// Start initialises the server resources.
func (db *Database) Open() {
	if db.isOpen {
		return
	}

	var err error
	db.SQL, err = openSQLiteDatabase("./qlueless.db")
	if err != nil {
		panic(err)
	}
}

// Close closes resources used by the server.
func (db *Database) Close() {
	if !db.isOpen {
		return
	}

	db.isOpen = false
	err := db.SQL.Close()
	db.SQL = nil

	if err != nil {
		panic(err)
	}
}

// openSQLiteDatabase opens a SQLite database, creating it if it doesn't already
// exist.
func openSQLiteDatabase(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		if db != nil {
			db.Close()
		}

		return nil, err
	}

	return db, nil
}
