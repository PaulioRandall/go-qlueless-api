package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// When given an thing, returns an thing ID
func TestAddThing___1(t *testing.T) {
	o := createDummyThing()
	act, err := AddThing(o)
	assert.Nil(t, err)
	assert.NotNil(t, act)

	stored, ok := ThingSlice[act.ID]
	assert.True(t, ok)
	assert.Equal(t, *act, stored)
}

// When invoked twice with the same thing, returns differnt IDs each time
func TestAddThing___2(t *testing.T) {
	a := createDummyThing()
	w1, err := AddThing(a)
	assert.Nil(t, err)
	a.ID = w1.ID
	a.Self = "/things/" + a.ID

	b := createDummyThing()
	w2, err := AddThing(b)
	assert.Nil(t, err)
	b.ID = w2.ID
	b.Self = "/things/" + b.ID

	assert.NotEqual(t, w1, w2)
	assert.NotEqual(t, *w1, *w2)

	stored, ok := ThingSlice[w1.ID]
	assert.True(t, ok)
	assert.Equal(t, a, stored)

	stored, ok = ThingSlice[w2.ID]
	assert.True(t, ok)
	assert.Equal(t, b, stored)
}
