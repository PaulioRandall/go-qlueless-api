//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"fmt"
	"log"
	"strings"

	d "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/database"
	u "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/utils"
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

	u.DeleteIfExists(path)
	db, err := d.OpenSQLiteDatabase(path)
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

		ven, err := nv.Insert(db)
		if err != nil {
			log.Fatal(err)
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
}
