//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	d "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/database"
	u "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/utils"
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

	u.DeleteIfExists(path)
	db := d.OpenDatabase(path)
	defer db.Close()

	for i, ven := range data {
		ven.Clean()
		errMsgs := ven.Validate(true)
		if len(errMsgs) > 0 {
			m := strings.Join(errMsgs, " ")
			log.Fatal(m)
		}

		data[i] = d.InsertVenture(db, ven)
	}

	vens := d.QueryVentures(db)
	for i, ven := range vens {
		fmt.Printf("%d: %v\n", i, ven)
	}
}
