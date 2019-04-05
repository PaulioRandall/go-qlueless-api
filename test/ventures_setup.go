package test

import (
	"database/sql"

	d "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/database"
	u "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/utils"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

var ventureHttpMethods = []string{"GET", "POST", "PUT", "OPTIONS"}
var dbPath string = "../bin/qlueless.db"
var venDB *sql.DB

// venDBReset will reset the database by closing and deleting it then
// creating a new one.
func venDBReset() {
	venDBClose()

	venDB = nil
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
}

// venDBInject injects a dummy Venture into the database.
func venDBInject(new v.NewVenture) *v.Venture {
	ven, err := new.Insert(venDB)
	if err != nil {
		panic(err)
	}
	return ven
}
