package things

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

// When given an ID to an existing Thing, it is returned
func TestFindThing___1(t *testing.T) {
	req, res, _ := SetupRequest("/")
	exp := DummyThing()
	Things.Add(exp)

	act, ok := findThing("1", res, req)
	assert.True(t, ok)
	assert.Equal(t, exp, act)
}

// When given an ID to a non-existing Thing, an 400 error is set
func TestFindThing___2(t *testing.T) {
	req, res, rec := SetupRequest("/")
	exp := NewDummyThing()
	Things.Add(exp)

	_, ok := findThing("999", res, req)
	assert.False(t, ok)
	assert.Equal(t, 404, rec.Code)

	assert.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m["message"])
}

// When given a valid Thing to decode, it is correctly decoded
func TestDecodeThing___1(t *testing.T) {
	req, res, _ := SetupRequest("/")
	json := `{
		"description": "description",
		"state": "state",
		"ID": "1",
		"self": "/things/1",
		"is_dead": false,
		"additional": "colour:red"
	}`

	req.Body = ioutil.NopCloser(strings.NewReader(json))
	act, ok := decodeThing(res, req)
	assert.True(t, ok)
	assert.Equal(t, "description", act.Description)
	assert.Equal(t, "state", act.State)
	assert.Equal(t, "1", act.ID)
	assert.Equal(t, "/things/1", act.Self)
	assert.Equal(t, false, act.IsDead)
	assert.Equal(t, "colour:red", act.Additional)
}

// When given invalid JSON to decode, false is returned
func TestDecodeThing___2(t *testing.T) {
	req, res, _ := SetupRequest("/")
	json := `[
		"Avantasia",
		"Dream Theater"
	]`

	req.Body = ioutil.NopCloser(strings.NewReader(json))
	act, ok := decodeThing(res, req)
	assert.False(t, ok)
	assert.Equal(t, Thing{}, act)
}

// When given an empty Thing to decode, an empty Thing is returned
func TestDecodeThing___3(t *testing.T) {
	req, res, _ := SetupRequest("/")
	req.Body = ioutil.NopCloser(strings.NewReader("{}"))
	act, ok := decodeThing(res, req)
	assert.True(t, ok)
	assert.Equal(t, Thing{}, act)
}

// When given a valid Thing, it validates successfully
func TestCheckThing___1(t *testing.T) {
	req, res, _ := SetupRequest("/")
	exp := Thing{
		Description: "description",
		State:       "state",
		IsDead:      false,
		Self:        "/things/1",
		ID:          "1",
	}

	act, ok := checkThing(exp, res, req)
	assert.True(t, ok)
	assert.Equal(t, exp, act)
}

// When given a valid Thing, it is cleaned successfully
func TestCheckThing___2(t *testing.T) {
	req, res, _ := SetupRequest("/")
	exp := Thing{
		Description: "   description   ",
		State:       "   state   ",
		ChildrenIDs: []string{"1", "0", "-1"},
		IsDead:      false,
		Self:        "/things/1",
		ID:          "1",
	}

	act, _ := checkThing(exp, res, req)
	exp.ChildrenIDs = []string{"1"}
	exp.Description = "description"
	exp.State = "state"
	assert.Equal(t, exp, act)
}

// When given an invalid Thing, validation fails
func TestCheckThing___3(t *testing.T) {
	req, res, rec := SetupRequest("/")
	exp := Thing{
		Description: "",
		State:       "",
		IsDead:      false,
		Self:        "/things/1",
		ID:          "1",
	}

	_, ok := checkThing(exp, res, req)
	assert.False(t, ok)
	assert.Equal(t, 400, rec.Code)

	assert.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m["message"])
}
