package pkg

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

// When given nil, does nothing
func TestCheck___1(t *testing.T) {
	Check(nil)
}

// When given an error, a panics ensues
func TestCheck___2(t *testing.T) {
	CheckPanic(t, func() {
		err := errors.New("Computer says no!")
		Check(err)
	})
}

// When given nil, returns false and doesn't print anything
func TestLogIfErr___1(t *testing.T) {
	act := LogIfErr(nil)
	assert.False(t, act)
	// Output:
	//
}

// When given an error, returns true and prints the error message
func TestLogIfErr___2(t *testing.T) {
	err := errors.New("Computer says no!")
	act := LogIfErr(err)
	assert.True(t, act)
	// Output:
	// Computer says no!
}

func createRequest() (*http.Request, *http.ResponseWriter, *httptest.ResponseRecorder) {
	req, err := http.NewRequest("GET", "http://example.com/angels", nil)
	if err != nil {
		panic(err)
	}
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	return req, &res, rec
}

// When invoked, sets 500 status code
func TestReply500___1(t *testing.T) {
	req, res, rec := createRequest()

	Reply500(res, req)
	assert.Equal(t, 500, rec.Code)
}

// When invoked, writes JSON headers
func TestReply500___2(t *testing.T) {
	req, res, _ := createRequest()

	Reply500(res, req)
	CheckJSONResponseHeaders(t, (*res).Header())
}

// When invoked, writes JSON headers
func TestReply500___3(t *testing.T) {
	req, res, rec := createRequest()

	Reply500(res, req)

	assert.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	assert.Nil(t, err)
	assert.Contains(t, m, "message")
	assert.Contains(t, m, "self")
}
