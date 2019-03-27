package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// When init with some Things, returns a map of the init Things
func TestThingStore___GetAll___1(t *testing.T) {
	ts := NewThingStore()

	a := NewDummyThing()
	a.ID = "1"
	ts.things[a.ID] = a

	b := NewDummyThing()
	b.ID = "2"
	ts.things[b.ID] = b

	act := ts.GetAll()
	assert.NotEmpty(t, act)
	assert.Equal(t, a, act["1"])
	assert.Equal(t, b, act["2"])
	assert.Len(t, act, 2)
}

// When init with no Things, returns a map of no Things
func TestThingStore___GetAll___2(t *testing.T) {
	ts := NewThingStore()
	act := ts.GetAll()
	assert.Empty(t, act)
}

// When init with some Things, returns a slice of only alive things
func TestThingStore___GetAllAlive___1(t *testing.T) {
	ts := NewThingStore()

	a := NewDummyThing()
	a.ID = "1"
	a.IsDead = true
	ts.things[a.ID] = a

	b := NewDummyThing()
	b.ID = "2"
	b.IsDead = false
	ts.things[b.ID] = b

	act := ts.GetAllAlive()
	assert.Len(t, act, 1)
	assert.Equal(t, b, act[0])
}

// When init with no Things, returns a slice with no Things in it
func TestThingStore___GetAllAlive___2(t *testing.T) {
	ts := NewThingStore()
	act := ts.GetAllAlive()
	assert.Empty(t, act)
}

// When requesting an existing Thing, it is returned
func TestThingStore___Get___1(t *testing.T) {
	ts := NewThingStore()
	a := DummyThing()

	ts.things["1"] = a
	act := ts.Get("1")
	assert.Equal(t, a, act)
}

// When requesting a non-existing Thing, nil is returned
func TestThingStore___Get___2(t *testing.T) {
	ts := NewThingStore()
	a := DummyThing()

	ts.things["1"] = a
	act := ts.Get("99999")
	assert.Empty(t, act)
}

// When given a new Thing, an ID is assigned and the Self set
func TestThingStore___Add___1(t *testing.T) {
	ts := NewThingStore()
	a := NewDummyThing()

	act := ts.Add(a)
	assert.Equal(t, "1", act.ID)
	assert.Equal(t, a.Description, act.Description)
	assert.Equal(t, a.State, act.State)
}

// When given a new Thing, the returned Thing and stored Thing are equal
func TestThingStore___Add___2(t *testing.T) {
	ts := NewThingStore()
	a := NewDummyThing()

	exp := ts.Add(a)
	act, ok := ts.things["1"]

	require.True(t, ok)
	assert.Equal(t, exp, act)
}

// When given a Thing to update, it is updated
func TestThingStore___Update___1(t *testing.T) {
	ts := NewThingStore()
	a := DummyThing()
	a.Description = "new"
	ts.things["1"] = a

	a2 := DummyThing()
	a2.ID = "1"
	a2.Description = "updated"
	err := ts.Update(a2)
	require.Nil(t, err)

	act, ok := ts.things["1"]
	require.True(t, ok)
	assert.Equal(t, a.State, act.State)
	assert.Equal(t, a2.Description, act.Description)
}

// When given a Thing to update without an ID, an error is returned and no
// changes are made to the ThingStore
func TestThingStore___Update___2(t *testing.T) {
	ts := NewThingStore()
	a := DummyThing()
	a.Description = "new"
	ts.things["1"] = a

	a2 := DummyThing()
	a2.ID = ""
	a2.Description = "updated"
	err := ts.Update(a2)
	require.NotNil(t, err)

	act, ok := ts.things["1"]
	require.True(t, ok)
	assert.Equal(t, a.State, act.State)
	assert.Equal(t, a.Description, act.Description)
}
