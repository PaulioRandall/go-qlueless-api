//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

var data []v.Venture = []v.Venture{
	v.Venture{
		Description: "Sunshine, lollipops and rainbows" +
			"\nEverything that's wonderful is what I feel when we're together",
		OrderIDs: "1,2,3,4",
		State:    "In Progress",
	},
	v.Venture{
		Description: "description",
		State:       "state",
		Extra:       "extra",
	},
}

// Entry point for the script
func main() {
	path := "../bin/qlueless.db"

	deleteIfExists(path)
	db := openDatabase(path)
	defer db.Close()

	for i, ven := range data {
		ven.Clean()
		errMsgs := ven.Validate(true)
		if len(errMsgs) > 0 {
			m := strings.Join(errMsgs, " ")
			log.Fatal(m)
		}

		data[i] = insertVenture(db, ven)
	}

	vens := queryVentures(db)
	for i, ven := range vens {
		fmt.Printf("%d: %v\n", i, ven)
	}
}

// deleteIfExists deletes the file at the path specified if it exists
func deleteIfExists(path string) {
	err := os.Remove(path)
	switch {
	case err == nil, os.IsNotExist(err):
	default:
		log.Fatal(err)
	}
}

// openDatabase opens a database, creating it if it doesn't already exist
func openDatabase(path string) *sql.DB {
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

// insertVenture inserts a new Venture into the database
func insertVenture(db *sql.DB, ven v.Venture) v.Venture {
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

// queryVentures queries the database for all Ventures
func queryVentures(db *sql.DB) []v.Venture {
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
