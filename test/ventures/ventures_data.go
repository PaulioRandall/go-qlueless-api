package ventures

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/PaulioRandall/go-cookies/toastify"
	"github.com/PaulioRandall/go-qlueless-api/api/database"
	"github.com/PaulioRandall/go-qlueless-api/api/server"
	"github.com/PaulioRandall/go-qlueless-api/api/ventures"
	"github.com/PaulioRandall/go-qlueless-api/shared/wrapped"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var dbPath string = ""

// SetupEmptyTest is run at the start of a test to setup the server but does
// not inject any test data.
func SetupEmptyTest() {
	dbPath = getDbPath()
	ResetDatabase()
	server.StartUp(true)
}

// SetupTest is run at the start of every test to setup the server and inject
// the test data.
func SetupTest() {
	dbPath = getDbPath()
	ResetDatabase()
	DBInjectLiving()
	DBInjectDead()
	server.StartUp(true)
}

// TearDown should be deferred straight after SetupTest() is run to close
// resources at the end of every test.
func TearDown() {
	server.Shutdown()
	CloseDatabase()
}

// ResetDatabase will reset the database by closing and deleting it then
// creating a new one.
func ResetDatabase() {
	CloseDatabase()
	deleteIfExists(dbPath)

	database.Open()

	err := ventures.CreateTables()
	if err != nil {
		panic(err)
	}
}

// CloseDatabase closes the test database.
func CloseDatabase() {
	database.Close()
}

// InjectAll injects a slice of Ventures into the database.
func InjectAll(new []ventures.NewVenture) []ventures.Venture {
	result := make([]ventures.Venture, len(new))
	for i, v := range new {
		result[i] = *Inject(v)
	}
	return result
}

// Inject injects a Venture into the database.
func Inject(new ventures.NewVenture) *ventures.Venture {
	ven, ok := new.Insert()
	if !ok {
		panic("Already printed above!")
	}
	return ven
}

// DBInjectLiving injects a default set of living Ventures into the database
func DBInjectLiving() {
	Inject(ventures.NewVenture{
		Description: "White wizard",
		State:       "Not started",
		Extra:       "colour: white; power: 9000",
	})
	Inject(ventures.NewVenture{
		Description: "Green lizard",
		State:       "In progress",
	})
	Inject(ventures.NewVenture{
		Description: "Pink gizzard",
		State:       "Finished",
	})
	Inject(ventures.NewVenture{
		Description: "Eddie Izzard",
		State:       "In Progress",
	})
	Inject(ventures.NewVenture{
		Description: "The Count of Tuscany",
		State:       "In Progress",
	})
}

// DBInjectDead injects a default set of dead Ventures into the database
func DBInjectDead() {
	s := []ventures.Venture{
		*Inject(ventures.NewVenture{
			Description: "Rose",
			State:       "Finised",
		}),
		*Inject(ventures.NewVenture{
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
		err := ven.Update()
		if err != nil {
			panic(err)
		}
	}
}

// DBQueryAll queries the database for all living ventures
func DBQueryAll() []ventures.Venture {
	rows, err := database.Get().Query(`
		SELECT id, last_modified, description, order_ids, state, extra
		FROM ql_venture
	`)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		panic(err)
	}

	return mapRows(rows)
}

// DBQueryMany queries the database for Ventures with the specified IDs
func DBQueryMany(ids string) []ventures.Venture {
	rows, err := database.Get().Query(fmt.Sprintf(`
		SELECT id, last_modified, description, order_ids, state, extra
		FROM ql_venture
		WHERE id IN (%s)`, ids))

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		panic(err)
	}

	return mapRows(rows)
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

// getDbPath gets the path to the database or panics if there is an error.
func getDbPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd + "/qlueless.db"
}

// deleteIfExists deletes the file at the path specified if it exist.
func deleteIfExists(path string) {
	err := os.Remove(path)
	if err != nil && os.IsExist(err) {
		panic(err)
	}
}

// mapRows is a file private function that maps rows from a database query into
// a slice of Ventures.
func mapRows(rows *sql.Rows) []ventures.Venture {
	vens := []ventures.Venture{}

	for rows.Next() {
		vens = append(vens, *mapRow(rows))
	}

	return vens
}

// mapRow is a file private function that maps a single row from a database
// query into a Venture.
func mapRow(rows *sql.Rows) *ventures.Venture {
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
