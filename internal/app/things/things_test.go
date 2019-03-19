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

	b := createDummyThing()
	w2, err := addThing(b)
	assert.Nil(t, err)
	b.ID = w2.ID

	assert.NotEqual(t, w1, w2)
	assert.NotEqual(t, *w1, *w2)

	stored, ok := things[w1.ID]
	assert.True(t, ok)
	assert.Equal(t, a, stored)

	stored, ok = things[w2.ID]
	assert.True(t, ok)
	assert.Equal(t, b, stored)
}
