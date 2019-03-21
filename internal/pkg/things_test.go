package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createDummyThing() Thing {
	return Thing{
		Description: "# Outline the saga\nCreate a rough outline of the new saga.",
		State:       "in_progress",
		IsDead:      false,
	}
}

// When given a new Thing to clean, a cleaned Thing is returned
func TestCleanThing___1(t *testing.T) {
	thing := Thing{
		Description: "  description  ",
		State:       "  state  ",
		Additional:  "  additional  ",
		ChildrenIDs: []string{"  1  ", "   ", "  3  "},
	}
	CleanThing(&thing)
	assert.Equal(t, "description", thing.Description)
	assert.Equal(t, "state", thing.State)
	assert.Equal(t, "additional", thing.Additional)
	assert.Equal(t, []string{"1", "3"}, thing.ChildrenIDs)
}

// When given an empty Thing to clean, an empty Thing is returned
func TestCleanThing___2(t *testing.T) {
	thing := Thing{}
	CleanThing(&thing)
	assert.Equal(t, "", thing.Description)
	assert.Equal(t, "", thing.State)
	assert.Equal(t, "", thing.Additional)
	assert.Empty(t, thing.ChildrenIDs)
}

// When given a Thing with nothing to clean, nothing is cleaned
func TestCleanThing___3(t *testing.T) {
	thing := Thing{
		Self:   "  self  ",
		ID:     "  id  ",
		IsDead: true,
	}
	CleanThing(&thing)
	assert.Equal(t, "", thing.Description)
	assert.Equal(t, "", thing.State)
	assert.Equal(t, "", thing.Additional)
	assert.Empty(t, thing.ChildrenIDs)
	assert.Equal(t, "  self  ", thing.Self)
	assert.Equal(t, "  id  ", thing.ID)
	assert.Equal(t, true, thing.IsDead)
}

// When given an empty string, the message is appended
func TestAppendIfEmpty___1(t *testing.T) {
	act := appendIfEmpty("", []string{}, "abc")
	assert.Len(t, act, 1)
	assert.Contains(t, act, "abc")
}

// When given an empty string, the message is appended
func TestAppendIfEmpty___2(t *testing.T) {
	act := appendIfEmpty("", []string{"xyz"}, "abc")
	assert.Len(t, act, 2)
	assert.Contains(t, act, "xyz")
	assert.Contains(t, act, "abc")
}

// When given a non-empty string, no appending occurs
func TestAppendIfEmpty___3(t *testing.T) {
	act := appendIfEmpty("NOT-EMPTY", []string{}, "abc")
	assert.Len(t, act, 0)
}

// When given a new valid Thing, an empty slice is returned
func TestValidateThing___1(t *testing.T) {
	thing := Thing{
		Description: "description",
		State:       "state",
		ChildrenIDs: []string{"1", "2"},
	}
	act := ValidateThing(&thing, true)
	assert.Empty(t, act)
}

// When given a non-new valid Thing, an empty slice is returned
func TestValidateThing___2(t *testing.T) {
	thing := Thing{
		Description: "description",
		State:       "state",
		ChildrenIDs: []string{"1", "2"},
		ID:          "1",
		Self:        "/self",
	}
	act := ValidateThing(&thing, false)
	assert.Empty(t, act)
}

// When given a non-new valid Thing without children, an empty slice is
// returned
func TestValidateThing___3(t *testing.T) {
	thing := Thing{
		Description: "description",
		State:       "state",
		ID:          "1",
		Self:        "/self",
	}
	act := ValidateThing(&thing, false)
	assert.Empty(t, act)
}

// When given a new Thing with invalid property values, a slice containing all
// expected errors is returned
func TestValidateThing___4(t *testing.T) {
	thing := Thing{
		Description: "",
		State:       "",
		ChildrenIDs: []string{"abc", "efg"},
	}
	act := ValidateThing(&thing, true)
	assert.Len(t, act, 4)
}

// When given a non-new Thing with invalid property values, a slice containing
// all expected errors is returned
func TestValidateThing___5(t *testing.T) {
	thing := Thing{
		Description: "",
		State:       "",
		ChildrenIDs: []string{"abc", "efg"},
		ID:          "",
		Self:        "",
	}
	act := ValidateThing(&thing, false)
	assert.Len(t, act, 6)
}
