package things

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

func createRequest(path string) (*http.Request, *http.ResponseWriter, *httptest.ResponseRecorder) {
	req, err := http.NewRequest("GET", "http://example.com"+path, nil)
	if err != nil {
		panic(err)
	}
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	return req, &res, rec
}

func createDummyThing() Thing {
	return Thing{
		Description: "# Outline the saga\nCreate a rough outline of the new saga.",
		State:       "in_progress",
		IsDead:      false,
		Self:        "/things/1",
		ID:          1,
	}
}

// When given a parseable ID, it is parsed and returned
func TestParseID___1(t *testing.T) {
	req, res, _ := createRequest("/")
	act, ok := parseID("1", res, req)
	assert.True(t, ok)
	assert.Equal(t, 1, act)
}

// When given an unparseable ID, 400 error is set
func TestParseID___2(t *testing.T) {
	req, res, rec := createRequest("/")

	_, ok := parseID("abc", res, req)
	assert.False(t, ok)
	assert.Equal(t, 400, rec.Code)

	assert.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m["message"])
}

// When given an ID to an existing Thing, it is returned
func TestFindThing___1(t *testing.T) {
	req, res, _ := createRequest("/")
	exp := createDummyThing()
	Things.Add(exp)

	act, ok := findThing(1, res, req)
	assert.True(t, ok)
	assert.Equal(t, exp, act)
}

// When given an ID to a non-existing Thing, an 400 error is set
func TestFindThing___2(t *testing.T) {
	req, res, rec := createRequest("/")
	exp := createDummyThing()
	Things.Add(exp)

	_, ok := findThing(999, res, req)
	assert.False(t, ok)
	assert.Equal(t, 404, rec.Code)

	assert.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m["message"])
}
