package ventures

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	q "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/qserver"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
	test "github.com/PaulioRandall/go-qlueless-assembly-api/test"
)

var ventureHttpMethods = []string{"GET", "POST", "PUT", "OPTIONS"}
var dbPath string = "../../bin/qlueless.db"
var venDB *sql.DB = nil

// beginVenTest is run at the start of every test to setup the server and
// inject the test data.
func beginVenTest() {
	venDBReset()
	venDBInjectLiving()
	venDBInjectDead()
	test.StartServer("../../bin")
}

// endVenTest should be deferred straight after _beginVenTest() is run to
// close resources at the end of every test.
func endVenTest() {
	test.StopServer()
	venDBClose()
}

// _deleteIfExists deletes the file at the path specified if it exist.
func _deleteIfExists(path string) {
	err := os.Remove(path)
	switch {
	case err == nil, os.IsNotExist(err):
	default:
		log.Fatal(err)
	}
}

// venDBReset will reset the database by closing and deleting it then
// creating a new one.
func venDBReset() {
	venDBClose()
	_deleteIfExists(dbPath)

	var err error
	venDB, err = q.OpenSQLiteDatabase(dbPath)
	if err != nil {
		panic(err)
	}

	err = v.CreateTables(venDB)
	if err != nil {
		panic(err)
	}
}

// venDBClose closes the test database.
func venDBClose() {
	if venDB != nil {
		err := venDB.Close()
		if err != nil {
			panic(err)
		}
	}
	venDB = nil
}

// venDBInject injects a Venture into the database.
func venDBInject(new v.NewVenture) *v.Venture {
	ven, ok := new.Insert(venDB)
	if !ok {
		panic("Already printed above!")
	}
	return ven
}

// venDBInjectLiving injects a default set of living Ventures into the database
func venDBInjectLiving() {
	venDBInject(v.NewVenture{
		Description: "White wizard",
		State:       "Not started",
		Extra:       "colour: white; power: 9000",
	})
	venDBInject(v.NewVenture{
		Description: "Green lizard",
		State:       "In progress",
	})
	venDBInject(v.NewVenture{
		Description: "Pink gizzard",
		State:       "Finished",
	})
	venDBInject(v.NewVenture{
		Description: "Eddie Izzard",
		State:       "In Progress",
	})
	venDBInject(v.NewVenture{
		Description: "The Count of Tuscany",
		State:       "In Progress",
	})
}

// venDBInjectDead injects a default set of dead Ventures into the database
func venDBInjectDead() {
	s := []v.Venture{
		*venDBInject(v.NewVenture{
			Description: "Rose",
			State:       "Finised",
		}),
		*venDBInject(v.NewVenture{
			Description: "Lily",
			State:       "Closed",
		}),
	}

	mod := v.ModVenture{
		Props: "is_dead",
		Values: v.Venture{
			IsDead: true,
		},
	}

	for _, ven := range s {
		mod.ApplyMod(&ven)
		err := ven.Update(venDB)
		if err != nil {
			panic(err)
		}
	}
}

// venDBQueryAll queries the database for all living ventures
func venDBQueryAll() []v.Venture {
	rows, err := venDB.Query(`
		SELECT id, last_modified, description, order_ids, state, extra
		FROM ql_venture
	`)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		panic(err)
	}

	return _mapRows(rows)
}

// venDBQueryMany queries the database for Ventures with the specified IDs
func venDBQueryMany(ids string) []v.Venture {
	rows, err := venDB.Query(fmt.Sprintf(`
		SELECT id, last_modified, description, order_ids, state, extra
		FROM ql_venture
		WHERE id IN (%s)
	`, ids))

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		panic(err)
	}

	return _mapRows(rows)
}

// venDBQueryOne queries the database for a specific Venture
func venDBQueryOne(id string) v.Venture {
	vens := venDBQueryMany(id)
	if len(vens) != 1 {
		panic("Expected a single venture from query")
	}
	return vens[0]
}

// _mapRows is a file private function that maps rows from a database query into
// a slice of Ventures.
func _mapRows(rows *sql.Rows) []v.Venture {
	vens := []v.Venture{}

	for rows.Next() {
		vens = append(vens, *_mapRow(rows))
	}

	return vens
}

// _mapRow is a file private function that maps a single row from a database
// query into a Venture.
func _mapRow(rows *sql.Rows) *v.Venture {
	ven := v.Venture{}
	err := rows.Scan(&ven.ID,
		&ven.LastModified,
		&ven.Description,
		&ven.OrderIDs,
		&ven.State,
		&ven.Extra)

	if err != nil {
		panic(err)
	}
	return &ven
}
