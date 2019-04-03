package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

// OpenDatabase opens a database, creating it if it doesn't already exist
func OpenDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	_fatalClose(db, err)

	err = v.CreateTable(db)
	_fatalClose(db, err)

	return db
}

// _fatalClose is a file private function that performs log.Fatal(err) after
// closing the supplied database
func _fatalClose(db *sql.DB, err error) {
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
}
