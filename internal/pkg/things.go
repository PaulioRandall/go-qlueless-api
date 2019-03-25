package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

// A Thing represents a... err... Thing
type Thing struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	ChildIDs    string `json:"child_ids,omitempty"`
	State       string `json:"state"`
	IsDead      bool   `json:"-"`
	Additional  string `json:"additional,omitempty"`
	Self        string `json:"self"`
}

// SplitChildIDs returns the IDs of the children as a slice
func (t *Thing) SplitChildIDs() []string {
	return strings.Split(t.ChildIDs, ",")
}

// SetChildIDs returns the IDs of the children as a slice
func (t *Thing) SetChildIDs(ids []string) {
	t.ChildIDs = strings.Join(ids, ",")
}

// CleanThing cleans up a Things contents, e.g. triming whitespace from its
// description
func (t *Thing) CleanThing() {
	t.Description = strings.TrimSpace(t.Description)
	t.Additional = strings.TrimSpace(t.Additional)
	t.State = strings.TrimSpace(t.State)
	t.cleanThingsChildIDs()
}

// cleanThingsChildIDs cleans the child IDs within a Thing
func (t *Thing) cleanThingsChildIDs() {
	if t.ChildIDs != "" {
		ids := t.SplitChildIDs()
		for i, l := 0, len(ids); i < l; {
			c, err := strconv.Atoi(ids[i])

			if err != nil || c < 1 {
				ids = DeleteStr(ids, i)
				l--
			} else {
				i++
			}
		}
		t.SetChildIDs(ids)
	}
}

// ValidateThing validates a Thing contains the required and valid content. The
// result will be an slice of strings each being a readable description of a
// violation. The result may be supplied to the client
func (t *Thing) ValidateThing(isNew bool) []string {
	var r []string

	r = appendIfEmpty((*t).Description, r, "'Description' must not be empty.")
	r = appendIfEmpty((*t).State, r, "'State' must not be empty.")

	if (*t).ChildIDs != "" {
		for _, c := range (*t).SplitChildIDs() {
			r = appendIfNotPositive(c, r,
				fmt.Sprintf("'ChildrenIDs:%s' must be defined and a positive integer.", c))
		}
	}

	if !isNew {
		r = appendIfNotPositive((*t).ID, r,
			fmt.Sprintf("The ID '%s' is not allowed to be zero or negative.", (*t).ID))
		r = appendIfEmpty((*t).Self, r, "The 'Self' must be present.")
	}

	return r
}

// appendIfEmpty appends 'm' to 'r' if 's' is empty
func appendIfEmpty(s string, r []string, m string) []string {
	if s == "" {
		return append(r, m)
	}
	return r
}

// appendIfNotPositive appends 'm' to 'r' if 's' is not a positive integer
func appendIfNotPositive(s string, r []string, m string) []string {
	i, err := strconv.Atoi(s)
	if err != nil || i < 1 {
		return append(r, m)
	}
	return r
}

// CreateDummyThings creates some dummy things for testing during these initial
// phases of development
func CreateDummyThings() {
	Things.Add(Thing{
		Description: "# Outline the saga\nCreate a rough outline of the new saga.",
		ID:          "1",
		ChildIDs:    "2,3,4",
		State:       "in_progress",
		Self:        "/things/1",
	})
	Things.Add(Thing{
		Description: "# Name the saga\nThink of a name for the saga.",
		ID:          "2",
		State:       "potential",
		Self:        "/things/2",
	})
	Things.Add(Thing{
		Description: "# Outline the first chapter",
		ID:          "3",
		State:       "delivered",
		Additional:  "archive_note:Done but not a compelling start",
		Self:        "/things/3",
	})
	Things.Add(Thing{
		Description: "# Outline the second chapter",
		ID:          "4",
		State:       "in_progress",
		Self:        "/things/4",
	})
}
