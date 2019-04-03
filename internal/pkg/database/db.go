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

	err = v.CreateVentureTable(db)
	_fatalClose(db, err)

	return db
}

// QueryVentures queries the database for all Ventures
func QueryVentures(db *sql.DB) []v.Venture {
	rows, err := db.Query(`SELECT
		id,
		description,
		order_ids,
		state,
		is_alive,
		extra
	FROM venture`)
	_fatalClose(db, err)

	r := []v.Venture{}

	for rows.Next() {
		ven := v.Venture{}
		err = rows.Scan(&ven.ID,
			&ven.Description,
			&ven.OrderIDs,
			&ven.State,
			&ven.IsAlive,
			&ven.Extra)
		_fatalClose(db, err)

		r = append(r, ven)
	}

	return r
}

// _fatalClose is a file private function that performs log.Fatal(err) after
// closing the supplied database
func _fatalClose(db *sql.DB, err error) {
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
}
