package std

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

//
// TODO: Refactor file as shared 'database' functionality
//
var Sev QServer = QServer{}

// QServer represents the server resources.
type QServer struct {
	DB *sql.DB
}

// Start initialises the server resources.
func (q *QServer) Start() {
	log.Println("[Go Qlueless API]: Starting application")

	var err error
	q.DB, err = OpenSQLiteDatabase("./qlueless.db")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[Go Qlueless API]: Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Close closes resources used by the server.
func (q *QServer) Close() {
	if q.DB != nil {
		log.Fatal(q.DB.Close())
		q.DB = nil
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
