package ventures

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"

	toastify "github.com/PaulioRandall/go-cookies/toastify"
	std "github.com/PaulioRandall/go-qlueless-api/api/std"
	ventures "github.com/PaulioRandall/go-qlueless-api/api/ventures"
	wrapped "github.com/PaulioRandall/go-qlueless-api/shared/wrapped"
	test "github.com/PaulioRandall/go-qlueless-api/test"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

var dbPath string = ""
var venDB *sql.DB = nil

// BeginEmptyTest is run at the start of a test to setup the server but does
// not inject any test data.
func BeginEmptyTest(relServerPath string) {
	dbPath = relServerPath + "/qlueless.db"
	DBReset()
	test.StartServer(relServerPath)
}

// BeginTest is run at the start of every test to setup the server and
// inject the test data.
func BeginTest(relServerPath string) {
	dbPath = relServerPath + "/qlueless.db"
	DBReset()
	DBInjectLiving()
	DBInjectDead()
	test.StartServer(relServerPath)
}

// EndTest should be deferred straight after BeginTest() is run to
// close resources at the end of every test.
func EndTest() {
	test.StopServer()
	DBClose()
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

// DBReset will reset the database by closing and deleting it then
// creating a new one.
func DBReset() {
	DBClose()
	_deleteIfExists(dbPath)

	var err error
	venDB, err = std.OpenSQLiteDatabase(dbPath)
	if err != nil {
		panic(err)
	}

	err = ventures.CreateTables(venDB)
	if err != nil {
		panic(err)
	}
}

// DBClose closes the test database.
func DBClose() {
	if venDB != nil {
		err := venDB.Close()
		if err != nil {
			panic(err)
		}
	}
	venDB = nil
}

// DBInject injects a Venture into the database.
func DBInject(new ventures.NewVenture) *ventures.Venture {
	ven, ok := new.Insert(venDB)
	if !ok {
		panic("Already printed above!")
	}
	return ven
}

// DBInjectLiving injects a default set of living Ventures into the database
func DBInjectLiving() {
	DBInject(ventures.NewVenture{
		Description: "White wizard",
		State:       "Not started",
		Extra:       "colour: white; power: 9000",
	})
	DBInject(ventures.NewVenture{
		Description: "Green lizard",
		State:       "In progress",
	})
	DBInject(ventures.NewVenture{
		Description: "Pink gizzard",
		State:       "Finished",
	})
	DBInject(ventures.NewVenture{
		Description: "Eddie Izzard",
		State:       "In Progress",
	})
	DBInject(ventures.NewVenture{
		Description: "The Count of Tuscany",
		State:       "In Progress",
	})
}

// DBInjectDead injects a default set of dead Ventures into the database
func DBInjectDead() {
	s := []ventures.Venture{
		*DBInject(ventures.NewVenture{
			Description: "Rose",
			State:       "Finised",
		}),
		*DBInject(ventures.NewVenture{
			Description: "Lily",
			State:       "Closed",
		}),
	}

	mod := ventures.ModVenture{
		Props: "is_dead",
		Values: ventures.Venture{
			Dead: true,
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

// DBQueryAll queries the database for all living ventures
func DBQueryAll() []ventures.Venture {
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

// DBQueryMany queries the database for Ventures with the specified IDs
func DBQueryMany(ids string) []ventures.Venture {
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

// DBQueryOne queries the database for a specific Venture
func DBQueryOne(id string) ventures.Venture {
	vens := DBQueryMany(id)
	if len(vens) != 1 {
		panic("Expected a single venture from query")
	}
	return vens[0]
}

// DBQueryFirst queries the database for the first Venture encountered
func DBQueryFirst() *ventures.Venture {
	vens := DBQueryAll()
	if len(vens) > 0 {
		return &vens[0]
	}
	return nil
}

// _mapRows is a file private function that maps rows from a database query into
// a slice of Ventures.
func _mapRows(rows *sql.Rows) []ventures.Venture {
	vens := []ventures.Venture{}

	for rows.Next() {
		vens = append(vens, *_mapRow(rows))
	}

	return vens
}

// _mapRow is a file private function that maps a single row from a database
// query into a Venture.
func _mapRow(rows *sql.Rows) *ventures.Venture {
	ven := ventures.Venture{}
	err := rows.Scan(&ven.ID,
		&ven.LastModified,
		&ven.Description,
		&ven.Orders,
		&ven.State,
		&ven.Extra)

	if err != nil {
		panic(err)
	}
	return &ven
}

// AssertHeaders asserts that the expected headers in 'h' have been supplied.
func AssertHeaders(t *testing.T, h http.Header) {
	toastify.AssertHeadersEqual(t, h, map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "*",
		"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
		"Content-Type":                 "application/json; charset=utf-8",
	})
}

// AssertGenericReply asserts that the response 'body' contains a generic reply.
func AssertGenericReply(t *testing.T, body io.Reader) {
	gr, err := wrapped.DecodeFromReader(body)
	require.Nil(t, err)
	assert.NotEmpty(t, gr.Message)
	assert.NotEmpty(t, gr.Self)
	assert.Empty(t, gr.Data)
}
