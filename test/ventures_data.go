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

var defaultVens []v.Venture = []v.Venture{}

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

// venDBInjectDefaults injects a default set of Ventures into the database
func venDBInjectDefaults() {
	defaultVens = append(defaultVens, *venDBInject(v.NewVenture{
		Description: "White wizard",
		State:       "Not started",
		Extra:       "colour: white; power: 9000",
	}))
	defaultVens = append(defaultVens, *venDBInject(v.NewVenture{
		Description: "Green lizard",
		State:       "In progress",
		OrderIDs:    "4,5,6,7,8",
	}))
	defaultVens = append(defaultVens, *venDBInject(v.NewVenture{
		Description: "Pink gizzard",
		State:       "Finished",
		OrderIDs:    "1,2,3",
	}))
	defaultVens = append(defaultVens, *venDBInject(v.NewVenture{
		Description: "Eddie Izzard",
		State:       "In Progress",
		OrderIDs:    "4,5,6",
	}))
	defaultVens = append(defaultVens, *venDBInject(v.NewVenture{
		Description: "The Count of Tuscany",
		State:       "In Progress",
		OrderIDs:    "4,5,6",
	}))
}
