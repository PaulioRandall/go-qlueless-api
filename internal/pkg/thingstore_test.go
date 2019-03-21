package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// When init with some Things, returns a map of the init Things
func TestThingStore___GetAll___1(t *testing.T) {
	ts := NewThingStore()

	a := createDummyThing()
	a.ID = "1"
	ts.things[a.ID] = a

	b := createDummyThing()
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

// When requesting an existing Thing, it is returned
func TestThingStore___Get___1(t *testing.T) {
	ts := NewThingStore()
	a := createDummyThing()
	a.ID = "1"
	a.Self = "/things/1"

	ts.things["1"] = a
	act := ts.Get("1")
	assert.Equal(t, a, act)
}

// When requesting a non-existing Thing, nil is returned
func TestThingStore___Get___2(t *testing.T) {
	ts := NewThingStore()
	a := createDummyThing()
	a.ID = "1"
	a.Self = "/things/1"

	ts.things["1"] = a
	act := ts.Get("99999")
	assert.Empty(t, act)
}

// When given a new Thing, an ID is assigned and the Self set
func TestThingStore___Add___1(t *testing.T) {
	ts := NewThingStore()
	a := createDummyThing()

	act, err := ts.Add(a)
	assert.Nil(t, err)
	assert.Equal(t, "1", act.ID)
	assert.Equal(t, "/things/1", act.Self)
	assert.Equal(t, a.Description, act.Description)
	assert.Equal(t, a.State, act.State)
}

// When given a new Thing, the returned Thing and stored Thing are equal
func TestThingStore___Add___2(t *testing.T) {
	ts := NewThingStore()
	a := createDummyThing()

	exp, err := ts.Add(a)
	assert.Nil(t, err)
	act, ok := ts.things["1"]

	assert.True(t, ok)
	assert.Equal(t, exp, act)
}

// When invalid internal state, an error is returned when calling Add
func TestThingStore___Add___3(t *testing.T) {
	ts := NewThingStore()
	ts.things["abc"] = createDummyThing()

	a := createDummyThing()
	_, err := ts.Add(a)
	assert.NotNil(t, err)
}
