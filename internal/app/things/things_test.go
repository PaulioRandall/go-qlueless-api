package things

import (
	"testing"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

func createDummyThing() Thing {
	return Thing{
		Description: "# Outline the saga\nCreate a rough outline of the new saga.",
		State:       "in_progress",
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
	cleanThing(&thing)
	assert.Equal(t, "description", thing.Description)
	assert.Equal(t, "state", thing.State)
	assert.Equal(t, "additional", thing.Additional)
	assert.Equal(t, []string{"1", "3"}, thing.ChildrenIDs)
}

// When given an empty Thing to clean, an empty Thing is returned
func TestCleanThing___2(t *testing.T) {
	thing := Thing{}
	cleanThing(&thing)
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
	cleanThing(&thing)
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
	act := validateThing(&thing, true)
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
	act := validateThing(&thing, false)
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
	act := validateThing(&thing, false)
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
	act := validateThing(&thing, true)
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
	act := validateThing(&thing, false)
	assert.Len(t, act, 6)
}

// When given an thing, returns an thing ID
func TestAddThing___1(t *testing.T) {
	o := createDummyThing()
	act, err := addThing(o)
	assert.Nil(t, err)
	assert.NotNil(t, act)

	stored, ok := things[act.ID]
	assert.True(t, ok)
	assert.Equal(t, *act, stored)
}

// When invoked twice with the same thing, returns differnt IDs each time
func TestAddThing___2(t *testing.T) {
	a := createDummyThing()
	w1, err := addThing(a)
	assert.Nil(t, err)
	a.ID = w1.ID
	a.Self = "/things/" + a.ID

	b := createDummyThing()
	w2, err := addThing(b)
	assert.Nil(t, err)
	b.ID = w2.ID
	b.Self = "/things/" + b.ID

	assert.NotEqual(t, w1, w2)
	assert.NotEqual(t, *w1, *w2)

	stored, ok := things[w1.ID]
	assert.True(t, ok)
	assert.Equal(t, a, stored)

	stored, ok = things[w2.ID]
	assert.True(t, ok)
	assert.Equal(t, b, stored)
}
