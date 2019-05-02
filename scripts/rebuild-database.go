//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	q "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/qserver"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

var input []v.NewVenture = []v.NewVenture{
	v.NewVenture{
		Description: "Sunshine, lollipops and rainbows" +
			"\nEverything that's wonderful is what I feel when we're together",
		OrderIDs: "1,2,3,4",
		State:    "In Progress",
	},
	v.NewVenture{
		Description: "description",
		State:       "state",
		Extra:       "extra",
	},
}

// Entry point for the script
func main() {
	path := "../bin/qlueless.db"

	_deleteIfExists(path)
	db, err := q.OpenSQLiteDatabase(path)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = v.CreateTables(db)
	if err != nil {
		log.Fatal(err)
	}

	output := []v.Venture{}

	for _, nv := range input {
		nv.Clean()
		errMsgs := nv.Validate()
		if len(errMsgs) > 0 {
			m := strings.Join(errMsgs, " ")
			log.Fatal(m)
		}

		ven, ok := nv.Insert(db)
		if !ok {
			log.Fatal("Error printed above!")
		}

		output = append(output, *ven)
	}

	vens, err := v.QueryAll(db)
	if err != nil {
		log.Fatal(err)
	}

	for i, ven := range vens {
		fmt.Printf("%d: %v\n", i, ven)
	}

	vens[1].Extra = "AAAAAGGGGGHHHHH"
	err = vens[1].Update(db)
	if err != nil {
		log.Fatal(err)
	}

	ven, err := v.QueryFor(db, vens[1].ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("UPDATED: %v\n", ven)
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
