package pkg

import (
	"fmt"
	"strings"
)

// A Thing represents a... err... Thing
type Thing struct {
	Description string `json:"description"`
	ID          int    `json:"id"`
	ChildrenIDs []int  `json:"childrens_ids,omitempty"`
	State       string `json:"state"`
	IsDead      bool   `json:"-"`
	Additional  string `json:"additional,omitempty"`
	Self        string `json:"self"`
}

// CleanThing cleans up a Things contents, e.g. triming whitespace from its
// description
func CleanThing(t *Thing) {
	t.Description = strings.TrimSpace(t.Description)
	t.Additional = strings.TrimSpace(t.Additional)
	t.State = strings.TrimSpace(t.State)

	if t.ChildrenIDs == nil {
		return
	}

	for i, l := 0, len(t.ChildrenIDs); i < l; {
		c := t.ChildrenIDs[i]

		if c < 1 {
			t.ChildrenIDs = DeleteInt(t.ChildrenIDs, i)
			l--
		} else {
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

// appendIfNotPositive appends 'm' to 'r' if 'i' is not positive
func appendIfNotPositive(i int, r []string, m string) []string {
	if i < 1 {
		return append(r, m)
	}
	return r
}

// ValidateThing validates a Thing contains the required and valid content. The
// result will be an slice of strings each being a readable description of a
// violation. The result may be supplied to the client
func ValidateThing(t *Thing, isNew bool) []string {
	var r []string

	r = appendIfEmpty((*t).Description, r, "'Description' must not be empty.")
	r = appendIfEmpty((*t).State, r, "'State' must not be empty.")

	for _, c := range (*t).ChildrenIDs {
		r = appendIfNotPositive(c, r,
			fmt.Sprintf("'ChildrenIDs:%d' must be defined and a positive integer.", c))
	}

	if !isNew {
		r = appendIfNotPositive((*t).ID, r,
			fmt.Sprintf("The ID '%d' is not allowed to be zero or negative.", (*t).ID))
		r = appendIfEmpty((*t).Self, r, "The 'Self' must be present.")
	}

	return r
}

// CreateDummyThings creates some dummy things for testing during these initial
// phases of development
func CreateDummyThings() {
	Things.Add(Thing{
		Description: "# Outline the saga\nCreate a rough outline of the new saga.",
		ID:          1,
		ChildrenIDs: []int{
			2,
			3,
			4,
		},
		State: "in_progress",
		Self:  "/things/1",
	})
	Things.Add(Thing{
		Description: "# Name the saga\nThink of a name for the saga.",
		ID:          2,
		State:       "potential",
		Self:        "/things/2",
	})
	Things.Add(Thing{
		Description: "# Outline the first chapter",
		ID:          3,
		State:       "delivered",
		Additional:  "archive_note:Done but not a compelling start",
		Self:        "/things/3",
	})
	Things.Add(Thing{
		Description: "# Outline the second chapter",
		ID:          4,
		State:       "in_progress",
		Self:        "/things/4",
	})
}
