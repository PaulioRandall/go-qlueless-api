package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// When given a new Thing to clean, a cleaned Thing is returned
func TestCleanThing___1(t *testing.T) {
	thing := Thing{
		Description: "  description  ",
		State:       "  state  ",
		Additional:  "  additional  ",
		ChildrenIDs: []int{1, 0, -1, 3},
	}
	CleanThing(&thing)
	assert.Equal(t, "description", thing.Description)
	assert.Equal(t, "state", thing.State)
	assert.Equal(t, "additional", thing.Additional)
	assert.Equal(t, []int{1, 3}, thing.ChildrenIDs)
}

// When given an empty Thing to clean, an empty Thing is returned
func TestCleanThing___2(t *testing.T) {
	thing := Thing{}
	CleanThing(&thing)
	assert.Equal(t, Thing{}, thing)
}

// When given a Thing with nothing to clean, nothing is cleaned
func TestCleanThing___3(t *testing.T) {
	thing := Thing{
		Self:   "  self  ",
		ID:     1,
		IsDead: true,
	}
	CleanThing(&thing)
	assert.Equal(t, "", thing.Description)
	assert.Equal(t, "", thing.State)
	assert.Equal(t, "", thing.Additional)
	assert.Empty(t, thing.ChildrenIDs)
	assert.Equal(t, "  self  ", thing.Self)
	assert.Equal(t, 1, thing.ID)
	assert.Equal(t, true, thing.IsDead)
}

// When given a new Thing to clean, its child IDs are cleaned
func TestCleanThingsChildIDs___1(t *testing.T) {
	thing := Thing{
		ChildrenIDs: []int{1, 0, -1, 3},
	}
	cleanThingsChildIDs(&thing)
	assert.Equal(t, []int{1, 3}, thing.ChildrenIDs)
}

// When given a new Thing with no child IDs, nothing changes
func TestCleanThingsChildIDs___2(t *testing.T) {
	thing := Thing{}
	cleanThingsChildIDs(&thing)
	var exp []int
	assert.Equal(t, exp, thing.ChildrenIDs)
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

// When given a positive integer, the message is NOT appended
func TestAppendIfNotPositive___1(t *testing.T) {
	act := appendIfNotPositive(5, []string{}, "abc")
	assert.Len(t, act, 0)
}

// When given zero, the message is appended
func TestAppendIfNotPositive___2(t *testing.T) {
	act := appendIfNotPositive(0, []string{}, "abc")
	assert.Len(t, act, 1)
	assert.Contains(t, act, "abc")
}

// When given a negative number, the message is appended
func TestAppendIfNotPositive___3(t *testing.T) {
	act := appendIfNotPositive(-5, []string{}, "abc")
	assert.Len(t, act, 1)
	assert.Contains(t, act, "abc")
}

// When given a multiple non-positive inputs, all messages are appended
func TestAppendIfNotPositive___4(t *testing.T) {
	act := []string{}
	act = appendIfNotPositive(-1, act, "abc")
	act = appendIfNotPositive(-1, act, "efg")
	act = appendIfNotPositive(-1, act, "hij")
	assert.Len(t, act, 3)
	assert.Contains(t, act, "abc")
	assert.Contains(t, act, "efg")
	assert.Contains(t, act, "hij")
}

// When given a new valid Thing, an empty slice is returned
func TestValidateThing___1(t *testing.T) {
	thing := Thing{
		Description: "description",
		State:       "state",
		ChildrenIDs: []int{2, 3},
	}
	act := ValidateThing(&thing, true)
	assert.Empty(t, act)
}

// When given a non-new valid Thing, an empty slice is returned
func TestValidateThing___2(t *testing.T) {
	thing := Thing{
		Description: "description",
		State:       "state",
		ChildrenIDs: []int{2, 3},
		ID:          1,
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
		ID:          1,
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
		ChildrenIDs: []int{0, -9000},
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
		ChildrenIDs: []int{0, -9000},
		ID:          0,
		Self:        "",
	}
	act := ValidateThing(&thing, false)
	assert.Len(t, act, 6)
}
