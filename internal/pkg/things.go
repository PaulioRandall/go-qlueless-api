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
}

// SplitChildIDs returns the IDs of the children as a slice
func (t *Thing) SplitChildIDs() []string {
	return strings.Split(t.ChildIDs, ",")
}

// SetChildIDs returns the IDs of the children as a slice
func (t *Thing) SetChildIDs(ids []string) {
	t.ChildIDs = strings.Join(ids, ",")
}

// Clean cleans up a Things contents, e.g. triming whitespace from its
// description, state, etc
func (t *Thing) Clean() {
	t.Description = strings.TrimSpace(t.Description)
	t.Additional = strings.TrimSpace(t.Additional)
	t.State = strings.TrimSpace(t.State)
	t.ChildIDs = StripWhitespace(t.ChildIDs)
}

// Validate validates a Thing contains the required and valid content. The
// result will be an slice of strings each being a readable description of a
// violation. The result may be supplied to the client
func (t *Thing) Validate(isNew bool) []string {
	var r []string

	r = appendIfEmpty((*t).Description, r, "'Description' must not be empty.")
	r = appendIfEmpty((*t).State, r, "'State' must not be empty.")

	if (*t).ChildIDs != "" {
		for i, c := range (*t).SplitChildIDs() {
			r = appendIfNotPositiveInt(c, r,
				fmt.Sprintf("'ChildIDs[%d]:%s' must be a positive integer.", i, c))
		}
	}

	if !isNew {
		r = appendIfNotPositiveInt((*t).ID, r,
			fmt.Sprintf("The ID '%s' must be a positive integer.", (*t).ID))
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

// appendIfNotPositiveInt appends 'm' to 'r' if 's' is not a positive integer
func appendIfNotPositiveInt(s string, r []string, m string) []string {
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
	})
	Things.Add(Thing{
		Description: "# Name the saga\nThink of a name for the saga.",
		ID:          "2",
		State:       "potential",
	})
	Things.Add(Thing{
		Description: "# Outline the first chapter",
		ID:          "3",
		State:       "delivered",
		Additional:  "archive_note:Done but not a compelling start",
	})
	Things.Add(Thing{
		Description: "# Outline the second chapter",
		ID:          "4",
		State:       "in_progress",
	})
}
