package ventures

import (
	"database/sql"

	d "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/database"
	u "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/utils"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
	test "github.com/PaulioRandall/go-qlueless-assembly-api/test"
)

var ventureHttpMethods = []string{"GET", "POST", "PUT", "OPTIONS"}
var dbPath string = "../../bin/qlueless.db"
var venDB *sql.DB = nil

var livingVens map[string]v.Venture = nil
var deadVens map[string]v.Venture = nil

// beginVenTest is run at the start of every test to setup the server and
// inject the test data.
func beginVenTest() {
	venDBReset()
	venDBInjectLivingVentures()
	venDBInjectDeadVentures()
	test.StartServer("../../bin")
}

// endVenTest should be deferred straight after _beginVenTest() is run to
// close resources at the end of every test.
func endVenTest() {
	test.StopServer()
	venDBClose()
}

// venDBReset will reset the database by closing and deleting it then
// creating a new one.
func venDBReset() {
	venDBClose()
	u.DeleteIfExists(dbPath)

	var err error
	venDB, err = d.OpenSQLiteDatabase(dbPath)
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
func venDBInject(data map[string]v.Venture, new v.NewVenture) {
	ven, err := new.Insert(venDB)
	if err != nil {
		panic(err)
	}
	data[ven.ID] = *ven
}

// venDBInjectLivingVentures injects a default set of living Ventures into the
// database
func venDBInjectLivingVentures() {
	livingVens = map[string]v.Venture{}
	venDBInject(livingVens, v.NewVenture{
		Description: "White wizard",
		State:       "Not started",
		Extra:       "colour: white; power: 9000",
	})
	venDBInject(livingVens, v.NewVenture{
		Description: "Green lizard",
		State:       "In progress",
	})
	venDBInject(livingVens, v.NewVenture{
		Description: "Pink gizzard",
		State:       "Finished",
	})
	venDBInject(livingVens, v.NewVenture{
		Description: "Eddie Izzard",
		State:       "In Progress",
	})
	venDBInject(livingVens, v.NewVenture{
		Description: "The Count of Tuscany",
		State:       "In Progress",
	})
}

// venDBInjectDeadVentures injects a default set of dead Ventures into the
// database
func venDBInjectDeadVentures() {
	deadVens = map[string]v.Venture{}
	venDBInject(deadVens, v.NewVenture{
		Description: "Rose",
		State:       "Finised",
	})
	venDBInject(deadVens, v.NewVenture{
		Description: "Lily",
		State:       "Closed",
	})

	mod := v.ModVenture{
		Props: "is_dead",
		Values: v.Venture{
			IsDead: true,
		},
	}

	for _, ven := range deadVens {
		mod.ApplyMod(&ven)
		err := ven.Update(venDB)
		if err != nil {
			panic(err)
		}
	}
}
