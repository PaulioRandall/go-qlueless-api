package test

import (
	"database/sql"

	d "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/database"
	u "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/utils"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

var ventureHttpMethods = []string{"GET", "POST", "PUT", "OPTIONS"}
var dbPath string = "../bin/qlueless.db"
var venDB *sql.DB = nil

var livingVens []v.Venture = []v.Venture{}
var deadVens []v.Venture = []v.Venture{}

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
func venDBInject(new v.NewVenture) *v.Venture {
	ven, err := new.Insert(venDB)
	if err != nil {
		panic(err)
	}
	return ven
}

// venDBInjectLivingVentures injects a default set of living Ventures into the
// database
func venDBInjectLivingVentures() {
	livingVens = append(livingVens, *venDBInject(v.NewVenture{
		Description: "White wizard",
		State:       "Not started",
		Extra:       "colour: white; power: 9000",
	}))
	livingVens = append(livingVens, *venDBInject(v.NewVenture{
		Description: "Green lizard",
		State:       "In progress",
	}))
	livingVens = append(livingVens, *venDBInject(v.NewVenture{
		Description: "Pink gizzard",
		State:       "Finished",
	}))
	livingVens = append(livingVens, *venDBInject(v.NewVenture{
		Description: "Eddie Izzard",
		State:       "In Progress",
	}))
	livingVens = append(livingVens, *venDBInject(v.NewVenture{
		Description: "The Count of Tuscany",
		State:       "In Progress",
	}))
}

// venDBInjectDeadVentures injects a default set of dead Ventures into the
// database
func venDBInjectDeadVentures() {
	deadVens = append(deadVens, *venDBInject(v.NewVenture{
		Description: "Rose",
		State:       "Finised",
	}))
	deadVens = append(deadVens, *venDBInject(v.NewVenture{
		Description: "Lily",
		State:       "Closed",
	}))

	mod := v.ModVenture{
		Props: "is_dead",
		Values: v.Venture{
			IsDead: true,
		},
	}

	mod.ApplyMod(&deadVens[0])
	deadVens[0].Update(venDB)

	mod.ApplyMod(&deadVens[1])
	deadVens[1].Update(venDB)
}
