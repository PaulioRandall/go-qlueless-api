package pkg

import (
	"errors"
	"net/http"
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

func dummyWorkItems() []WorkItem {
	return []WorkItem{
		WorkItem{
			Title:        "Outline the saga",
			Description:  "Create a rough outline of the new saga.",
			Work_item_id: "1",
			Tag_id:       "mid",
			Status_id:    "in_progress",
		},
		WorkItem{
			Title:               "Name the saga",
			Description:         "Think of a name for the saga.",
			Work_item_id:        "2",
			Parent_work_item_id: "1",
			Tag_id:              "mid",
			Status_id:           "potential",
		},
		WorkItem{
			Title:               "Outline the first chapter",
			Description:         "Outline the first chapter.",
			Work_item_id:        "3",
			Parent_work_item_id: "1",
			Tag_id:              "mid",
			Status_id:           "delivered",
			Additional:          "archive_note:Done but not a compelling start",
		},
		WorkItem{
			Title:               "Outline the second chapter",
			Description:         "Outline the second chapter.",
			Work_item_id:        "4",
			Parent_work_item_id: "1",
			Tag_id:              "mid",
			Status_id:           "in_progress",
		},
	}
}

// When given some WorkItems and the ID of an item that appears within those
// WorkItems, does NOT return nil
func TestFindWorkItem___1(t *testing.T) {
	items := dummyWorkItems()
	var w *WorkItem
	w = FindWorkItem(items, "1")
	assert.NotNil(t, w)
	w = FindWorkItem(items, "3")
	assert.NotNil(t, w)
}

// When given some batches and the ID of an item that does NOT appear within
// those batches, returns nil
func TestFindWorkItem___2(t *testing.T) {
	items := dummyWorkItems()
	id := "99999999999"
	w := FindWorkItem(items, id)
	assert.Nil(t, w)
}

// When given some batches and the ID of an item that appears within those
// batches, returns the expected batch
func TestFindWorkItem___3(t *testing.T) {
	items := dummyWorkItems()
	id := "2"
	w := FindWorkItem(items, id)
	assert.Equal(t, id, w.Work_item_id)
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

// When given an empty string, false is returned
func TestIsWrapperProp___1(t *testing.T) {
	act := isWrapperProp("")
	assert.False(t, act)
}

// When given an invalid wrapper property, false is returned
func TestIsWrapperProp___2(t *testing.T) {
	act := isWrapperProp("invalid")
	assert.False(t, act)
}

// When given a valid wrapper property, true is returned
func TestIsWrapperProp___3(t *testing.T) {
	act := isWrapperProp("message")
	assert.True(t, act)
}

// When given an empty string, an error is returned
func TestWrapWith___1(t *testing.T) {
	act, err := wrapWith("")
	assert.Nil(t, act)
	assert.NotNil(t, err)
}

// When given a dot separated list containing an empty value, an error is
// returned
func TestWrapWith___2(t *testing.T) {
	act, err := wrapWith("message..self")
	assert.Nil(t, act)
	assert.NotNil(t, err)
}

// When given a dot separated list containing an invalid value, an error is
// returned
func TestWrapWith___3(t *testing.T) {
	act, err := wrapWith("invalid")
	assert.Nil(t, act)
	assert.NotNil(t, err)
}

// When given a dot separated list containing an invalid value, an error is
// returned
func TestWrapWith___4(t *testing.T) {
	act, err := wrapWith("message.invalid")
	assert.Nil(t, act)
	assert.NotNil(t, err)
}

// When given a dot separated list containing a single valid value, the value
// is returned within an ordered array
func TestWrapWith___5(t *testing.T) {
	act, err := wrapWith("message")
	assert.NotNil(t, act)
	assert.Nil(t, err)
	assert.Len(t, act, 1)
	assert.Equal(t, "message", act[0])
}

// When given a dot separated list containing a multiple valid values, the
// values are returned within an ordered array
func TestWrapWith___6(t *testing.T) {
	act, err := wrapWith("message.self")
	assert.NotNil(t, act)
	assert.Nil(t, err)
	assert.Len(t, act, 2)
	assert.Equal(t, "message", act[0])
	assert.Equal(t, "self", act[1])
}

// When given a request without a 'wrap_with' query parameter, no values should
// be returned
func TestWrapData___1(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		panic(err)
	}
	r := Reply{
		Req: req,
	}
	act, err := wrapData(&r)
	assert.Nil(t, act)
	assert.Nil(t, err)
}

// When given a request with an invalid 'wrap_with' query parameter, the values
// should be returned
func TestWrapData___2(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/?wrap_with=abc.efg", nil)
	if err != nil {
		panic(err)
	}
	r := Reply{
		Req: req,
	}
	act, err := wrapData(&r)
	assert.Nil(t, act)
	assert.NotNil(t, err)
}

// When given a request with a valid 'wrap_with' query parameter, the values
// should be returned
func TestWrapData___3(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/?wrap_with=message.self", nil)
	if err != nil {
		panic(err)
	}
	r := Reply{
		Req: req,
	}
	act, err := wrapData(&r)
	assert.NotNil(t, act)
	assert.Nil(t, err)
	assert.Len(t, act, 2)
	assert.Equal(t, "message", act[0])
	assert.Equal(t, "self", act[1])
}
