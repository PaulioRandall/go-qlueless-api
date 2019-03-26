package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// When given Thing with mutliple child IDs, they are split from their string
// property and returned
func TestSplitChildIDs___1(t *testing.T) {
	thing := Thing{
		ChildIDs: "0,1,2",
	}
	act := thing.SplitChildIDs()
	assert.Len(t, act, 3)
	assert.Equal(t, "0", act[0])
	assert.Equal(t, "1", act[1])
	assert.Equal(t, "2", act[2])
}

// When given Thing with no child IDs, they are split from their string
// property and returned
func TestSplitChildIDs___2(t *testing.T) {
	thing := Thing{}
	act := thing.SplitChildIDs()
	assert.Len(t, act, 0)
}

// When given Thing with one child IDs, it is returned as the only entry in the
// resultant slice
func TestSplitChildIDs___3(t *testing.T) {
	thing := Thing{
		ChildIDs: "9",
	}
	act := thing.SplitChildIDs()
	assert.Len(t, act, 1)
	assert.Equal(t, "9", act[0])
}

// When given a slice with multiple IDs, a CSV of those IDs is set in the Thing
func TestSetChildIDs___1(t *testing.T) {
	thing := Thing{}
	IDs := []string{"1", "2", "3"}

	thing.SetChildIDs(IDs)
	assert.Equal(t, "1,2,3", thing.ChildIDs)
}

// When given an empty slice, an empty is set
func TestSetChildIDs___2(t *testing.T) {
	thing := Thing{}
	IDs := []string{}

	thing.SetChildIDs(IDs)
	assert.Equal(t, "", thing.ChildIDs)
}

// When given a slice with a single ID, the ID is set
func TestSetChildIDs___3(t *testing.T) {
	thing := Thing{}
	IDs := []string{"9"}

	thing.SetChildIDs(IDs)
	assert.Equal(t, "9", thing.ChildIDs)
}

// When given Thing with mutliple parent IDs, they are split from their string
// property and returned
func TestSplitParentIDs___1(t *testing.T) {
	thing := Thing{
		ParentIDs: "0,1,2",
	}
	act := thing.SplitParentIDs()
	assert.Len(t, act, 3)
	assert.Equal(t, "0", act[0])
	assert.Equal(t, "1", act[1])
	assert.Equal(t, "2", act[2])
}

// When given Thing with no parent IDs, they are split from their string
// property and returned
func TestSplitParentIDs___2(t *testing.T) {
	thing := Thing{}
	act := thing.SplitParentIDs()
	assert.Len(t, act, 0)
}

// When given Thing with one parent ID, it is returned as the only entry in the
// resultant slice
func TestSplitParentIDs___3(t *testing.T) {
	thing := Thing{
		ParentIDs: "9",
	}
	act := thing.SplitParentIDs()
	assert.Len(t, act, 1)
	assert.Equal(t, "9", act[0])
}

// When given a slice with multiple IDs, a CSV of those IDs is set in the Thing
func TestSetParentIDs___1(t *testing.T) {
	thing := Thing{}
	IDs := []string{"1", "2", "3"}

	thing.SetParentIDs(IDs)
	assert.Equal(t, "1,2,3", thing.ParentIDs)
}

// When given an empty slice, an empty is set
func TestSetParentIDs___2(t *testing.T) {
	thing := Thing{}
	IDs := []string{}

	thing.SetParentIDs(IDs)
	assert.Equal(t, "", thing.ParentIDs)
}

// When given a slice with a single ID, the ID is set
func TestSetParentIDs___3(t *testing.T) {
	thing := Thing{}
	IDs := []string{"9"}

	thing.SetParentIDs(IDs)
	assert.Equal(t, "9", thing.ParentIDs)
}

// When given a new Thing to clean, a cleaned Thing is returned
func TestClean___1(t *testing.T) {
	thing := Thing{
		Description: "  description  ",
		State:       "  state  ",
		Additional:  "  additional  ",
		ChildIDs:    "  1,3  ",
		ParentIDs:   "  2,4  ",
	}
	thing.Clean()
	assert.Equal(t, "description", thing.Description)
	assert.Equal(t, "state", thing.State)
	assert.Equal(t, "additional", thing.Additional)
	assert.Equal(t, "1,3", thing.ChildIDs)
	assert.Equal(t, "2,4", thing.ParentIDs)
}

// When given an empty Thing to clean, an empty Thing is returned
func TestClean___2(t *testing.T) {
	thing := Thing{}
	thing.Clean()
	assert.Equal(t, Thing{}, thing)
}

// When given a Thing with nothing to clean, nothing is cleaned
func TestClean___3(t *testing.T) {
	thing := Thing{
		ID:     "1",
		IsDead: true,
	}
	thing.Clean()
	assert.Equal(t, "", thing.Description)
	assert.Equal(t, "", thing.State)
	assert.Equal(t, "", thing.Additional)
	assert.Empty(t, thing.ChildIDs)
	assert.Empty(t, thing.ParentIDs)
	assert.Equal(t, "1", thing.ID)
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

// When given a positive integer, the message is NOT appended
func TestAppendIfNotPositiveInt___1(t *testing.T) {
	act := appendIfNotPositiveInt("5", []string{}, "abc")
	assert.Len(t, act, 0)
}

// When given zero, the message is appended
func TestAppendIfNotPositiveInt___2(t *testing.T) {
	act := appendIfNotPositiveInt("0", []string{}, "abc")
	assert.Len(t, act, 1)
	assert.Contains(t, act, "abc")
}

// When given a negative number, the message is appended
func TestAppendIfNotPositiveInt___3(t *testing.T) {
	act := appendIfNotPositiveInt("-5", []string{}, "abc")
	assert.Len(t, act, 1)
	assert.Contains(t, act, "abc")
}

// When given a multiple non-positive inputs, all messages are appended
func TestAppendIfNotPositiveInt___4(t *testing.T) {
	act := []string{}
	act = appendIfNotPositiveInt("-1", act, "abc")
	act = appendIfNotPositiveInt("-1", act, "efg")
	act = appendIfNotPositiveInt("-1", act, "hij")
	assert.Len(t, act, 3)
	assert.Contains(t, act, "abc")
	assert.Contains(t, act, "efg")
	assert.Contains(t, act, "hij")
}

// When given a new valid Thing, an empty slice is returned
func TestValidate___1(t *testing.T) {
	thing := Thing{
		Description: "description",
		State:       "state",
		ChildIDs:    "2,3",
		ParentIDs:   "4,5",
	}
	act := thing.Validate(true)
	assert.Empty(t, act)
}

// When given a non-new valid Thing, an empty slice is returned
func TestValidate___2(t *testing.T) {
	thing := Thing{
		Description: "description",
		State:       "state",
		ChildIDs:    "2,3",
		ParentIDs:   "4,5",
		ID:          "1",
	}
	act := thing.Validate(false)
	assert.Empty(t, act)
}

// When given a non-new valid Thing without children, an empty slice is
// returned
func TestValidate___3(t *testing.T) {
	thing := Thing{
		Description: "description",
		State:       "state",
		ID:          "1",
	}
	act := thing.Validate(false)
	assert.Empty(t, act)
}

// When given a new Thing with invalid property values, a slice containing all
// expected errors is returned
func TestValidate___4(t *testing.T) {
	thing := Thing{
		Description: "",
		State:       "",
		ChildIDs:    "0,-9000",
		ParentIDs:   "0,-9000",
		ID:          "0",
	}
	act := thing.Validate(true)
	assert.Len(t, act, 6)
}

// When given a new Thing with invalid property values, a slice containing all
// expected errors is returned
func TestValidate___5(t *testing.T) {
	thing := Thing{
		Description: "",
		State:       "",
		ChildIDs:    "abc,!!!",
		ParentIDs:   "abc,!!!",
		ID:          "0",
	}
	act := thing.Validate(false)
	assert.Len(t, act, 7)
}
