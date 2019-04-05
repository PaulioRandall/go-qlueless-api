package qserver

import (
	"database/sql"
	"log"

	d "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/database"
)

var Sev QServer = QServer{}

// QServer represents the server resources.
type QServer struct {
	DB *sql.DB
}

// Init initialises the server resources.
func (s *QServer) Init() {
	var err error
	s.DB, err = d.OpenSQLiteDatabase("./qlueless.db")
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
