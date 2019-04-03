package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// OpenSQLiteDatabase opens a SQLite database, creating it if it doesn't already
// exist
//
// @UNTESTED
func OpenSQLiteDatabase(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		if db != nil {
			db.Close()
		}

		return nil, err
	}

	return db, nil
}
