package pkg

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogIfErr___1(t *testing.T) {
	act := LogIfErr(nil)
	assert.False(t, act)
	// Output:
	//
}

func TestLogIfErr___2(t *testing.T) {
	var err error = errors.New("Computer says no!")
	act := LogIfErr(err)
	assert.True(t, act)
	// Output:
	// Computer says no!
}

func dummyWorkItems() *[]WorkItem {
	return &[]WorkItem{
		WorkItem{
			Description: "# Outline the saga\nCreate a rough outline of the new saga.",
			WorkItemID:  "1",
			TagID:       "mid",
			StatusID:    "in_progress",
		},
		WorkItem{
			Description:      "# Name the saga\nThink of a name for the saga.",
			WorkItemID:       "2",
			ParentWorkItemID: "1",
			TagID:            "mid",
			StatusID:         "potential",
		},
		WorkItem{
			Description:      "# Outline the first chapter",
			WorkItemID:       "3",
			ParentWorkItemID: "1",
			TagID:            "mid",
			StatusID:         "delivered",
			Additional:       "archive_note:Done but not a compelling start",
		},
		WorkItem{
			Description:      "# Outline the second chapter",
			WorkItemID:       "4",
			ParentWorkItemID: "1",
			TagID:            "mid",
			StatusID:         "in_progress",
		},
	}
}

// When given an empty string, true is returned
func TestIsBlank___1(t *testing.T) {
	act := IsBlank("")
	assert.True(t, act)
}

// When given a string with whitespace, true is returned
func TestIsBlank___2(t *testing.T) {
	act := IsBlank("\r\n \t\f")
	assert.True(t, act)
}

// When given a string with no whitespaces, false is returned
func TestIsBlank___3(t *testing.T) {
	act := IsBlank("Captain Vimes")
	assert.False(t, act)
}

// When a value is present, it is returned
func TestValueOrEmpty___1(t *testing.T) {
	m := make(map[string]interface{})
	m["key"] = "value"
	act := ValueOrEmpty(m, "key")
	assert.Equal(t, "value", act)
}

// When a value is not present, empty string is returned
func TestValueOrEmpty___2(t *testing.T) {
	m := make(map[string]interface{})
	m["key"] = "value"
	act := ValueOrEmpty(m, "responsibilities")
	assert.Empty(t, act)
}
