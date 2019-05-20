package std

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Sev QServer = QServer{}

// QServer represents the server resources.
type QServer struct {
	DB *sql.DB
}

// Init initialises the server resources.
func (s *QServer) Init() {
	var err error
	s.DB, err = OpenSQLiteDatabase("./qlueless.db")
	if err != nil {
		log.Fatal(err)
	}
}

// Close closes resources used by the server.
func (s *QServer) Close() {
	if s.DB != nil {
		log.Fatal(s.DB.Close())
		s.DB = nil
	}
}

// OpenSQLiteDatabase opens a SQLite database, creating it if it doesn't already
// exist.
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
