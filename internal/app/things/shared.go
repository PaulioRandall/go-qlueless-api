package things

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var things = map[string]Thing{}

// methodNotAllowed handles cases where a HTTP method has been used but is not
// handled by this particular endpoint
func methodNotAllowed(res *http.ResponseWriter, req *http.Request) {
	reply := Reply4XX{
		Res:     res,
		Req:     req,
		Message: fmt.Sprintf("Method not allowed for this endpoint (%s)", req.Method),
	}
	Write4XXReply(405, &reply)
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
