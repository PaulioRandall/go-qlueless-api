package database

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"

	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

// OpenDatabase opens a database, creating it if it doesn't already exist
func OpenDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	_fatalClose(db, err)

	stmt, err := db.Prepare(`CREATE TABLE venture (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		order_ids TEXT NOT NULL,
		state TEXT NOT NULL,
		is_alive BOOL DEFAULT TRUE,
		extra TEXT DEFAULT NULL
	);`)
	_fatalClose(db, err)

	_, err = stmt.Exec()
	_fatalClose(db, err)

	return db
}

// InsertVenture inserts a new Venture into the database
func InsertVenture(db *sql.DB, ven v.Venture) v.Venture {
	stmt, err := db.Prepare(`INSERT INTO venture (
		description, order_ids, state, extra
	) VALUES (
		?, ?, ?, ?
	);`)
	_fatalClose(db, err)

	res, err := stmt.Exec(ven.Description, ven.OrderIDs, ven.State, ven.Extra)
	_fatalClose(db, err)

	id, err := res.LastInsertId()
	_fatalClose(db, err)

	ven.ID = strconv.FormatInt(id, 10)
	return ven
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
