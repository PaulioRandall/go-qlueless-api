package things

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var things = map[string]Thing{}

// cleanThing cleans up a Things contents, e.g. triming whitespace from its
// description
func cleanThing(t *Thing) {
	t.Description = strings.TrimSpace(t.Description)
	t.Additional = strings.TrimSpace(t.Additional)
	t.State = strings.TrimSpace(t.State)

	if t.ChildrenIDs == nil {
		return
	}

	for i, l := 0, len(t.ChildrenIDs); i < l; {
		c := t.ChildrenIDs[i]
		c = strings.TrimSpace(c)

		if c == "" {
			t.ChildrenIDs = DeleteStr(t.ChildrenIDs, i)
			l--
		} else {
			t.ChildrenIDs[i] = c
			i++
		}
	}
}

// appendIfEmpty appends 'm' to 'r' if 's' is empty
func appendIfEmpty(s string, r []string, m string) []string {
	if s == "" {
		return append(r, m)
	}
	return r
}

// validateThing validates a Thing contains the required and valid content. The
// result will be an slice of strings each being a readable description of a
// violation. The result may be supplied to the client
func validateThing(t *Thing, isNew bool) []string {
	var r []string

	r = appendIfEmpty((*t).Description, r, "'Description' must not be empty.")
	r = appendIfEmpty((*t).State, r, "'State' must not be empty.")

	for _, c := range (*t).ChildrenIDs {
		if !IsInt(c) {
			r = append(r, fmt.Sprintf("'ChildrenIDs:%s' is not an integer.", c))
		}
	}

	if !isNew {
		r = appendIfEmpty((*t).ID, r, "The 'ID' must be present.")
		r = appendIfEmpty((*t).Self, r, "The 'Self' must be present.")
	}

	return r
}

// AddThing adds a new thing to the data store returning the newly assigned ID
func addThing(t Thing) (*Thing, error) {
	next := 1
	for k, _ := range things {
		ID, err := strconv.Atoi(k)
		if LogIfErr(err) {
			return nil, errors.New("[BUG] An unparsable ID exists within the data store")
		}

		if ID > next {
			next = ID
		}
	}

	next++
	t.ID = strconv.Itoa(next)
	t.Self = fmt.Sprintf("/things/%s", t.ID)
	things[t.ID] = t
	return &t, nil
}

// CreateDummyThings creates some dummy things for testing during these initial
// phases of development
func CreateDummyThings() {
	things["1"] = Thing{
		Description: "# Outline the saga\nCreate a rough outline of the new saga.",
		ID:          "1",
		ChildrenIDs: []string{
			"2",
			"3",
			"4",
		},
		State: "in_progress",
		Self:  "/things/1",
	}
	things["2"] = Thing{
		Description: "# Name the saga\nThink of a name for the saga.",
		ID:          "2",
		State:       "potential",
		Self:        "/things/2",
	}
	things["3"] = Thing{
		Description: "# Outline the first chapter",
		ID:          "3",
		State:       "delivered",
		Additional:  "archive_note:Done but not a compelling start",
		Self:        "/things/3",
	}
	things["4"] = Thing{
		Description: "# Outline the second chapter",
		ID:          "4",
		State:       "in_progress",
		Self:        "/things/4",
	}
}
