package things

import (
	"testing"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

// When a valid map is provided, a Thing is returned
func TestMapToThing___1(t *testing.T) {
	m := make(map[string]interface{})
	m["description"] = "description"

	act := mapToThing(m)
	assert.Equal(t, "description", act.Description)
	assert.Empty(t, act.ID)
	assert.Empty(t, act.ChildrenIDs)
	assert.Empty(t, act.State)
	assert.Empty(t, act.Additional)
}

// When a valid map is provided, a Thing is returned
func TestMapToThing___2(t *testing.T) {
	m := make(map[string]interface{})
	m["description"] = "description"
	m["id"] = "id"
	m["childrens_ids"] = []string{"2", "3"}
	m["state"] = "state"
	m["additional"] = "abc: xyz; colour: black"

	act := mapToThing(m)
	assert.Equal(t, "description", act.Description)
	assert.Equal(t, "id", act.ID)
	assert.Equal(t, []string{"2", "3"}, act.ChildrenIDs)
	assert.Equal(t, "state", act.State)
	assert.Equal(t, "abc: xyz; colour: black", act.Additional)
}

func createDummyThing() Thing {
	return Thing{
		Description: "# Outline the saga\nCreate a rough outline of the new saga.",
		State:       "in_progress",
	}
}

// When given an thing, returns an thing ID
func TestAddThing___1(t *testing.T) {
	o := createDummyThing()
	act, err := addThing(o)
	assert.Nil(t, err)
	assert.NotEmpty(t, act)
	o.ID = act

	stored, ok := things[o.ID]
	assert.True(t, ok)
	assert.Equal(t, o, stored)
}

// When invoked twice with the same thing, returns differnt IDs each time
func TestAddThing___2(t *testing.T) {
	a := createDummyThing()
	ID_1, err := addThing(a)
	assert.Nil(t, err)
	a.ID = ID_1

	b := createDummyThing()
	ID_2, err := addThing(b)
	assert.Nil(t, err)
	b.ID = ID_2

	assert.NotEqual(t, ID_1, ID_2)

	stored, ok := things[a.ID]
	assert.True(t, ok)
	assert.Equal(t, a, stored)

	stored, ok = things[b.ID]
	assert.True(t, ok)
	assert.Equal(t, b, stored)
}
