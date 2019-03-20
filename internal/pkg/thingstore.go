package pkg

import (
	"errors"
	"fmt"
	"strconv"
)

// A ThingStore provides synchronisation for accessing Things
type ThingStore struct {
	setChan chan Thing
	getChan chan map[string]Thing
}

var ThingSlice = map[string]Thing{}

// NewThingStore creates a new ThingStore
func NewThingStore() ThingStore {
	wis := ThingStore{
		setChan: make(chan Thing),
		getChan: make(chan map[string]Thing),
	}
	go wis.mux()
	return wis
}

// mux should be invoked as a goroutine; it forever loops handling incoming
// channel communications sequentially
func (wis ThingStore) mux() {
	m := make(map[string]Thing)
	var r map[string]Thing
	for {
		select {
		case w := <-wis.setChan:
			m[w.ID] = w
		case wis.getChan <- r:
			t := make(map[string]Thing)
			for k, v := range m {
				if !v.IsDead {
					t[k] = v
				}
			}
		}
	}
}

// Get gets all Things
func (wis ThingStore) Get() map[string]Thing {
	return <-wis.getChan
}

// Set sets a Thing
func (wis ThingStore) Set(t Thing) {
	wis.setChan <- t
}

// AddThing adds a new thing to the data store returning the newly assigned ID
func AddThing(t Thing) (*Thing, error) {
	next := 1
	for k, _ := range ThingSlice {
		ID, err := strconv.Atoi(k)
		if LogIfErr(err) {
			return nil, errors.New("[BUG] An unparsable ID exists within the data store")
		}

		if ID > next {
			next = ID
		}
	}

	next++
	t.ID = strconv.Itoa(next)
	t.Self = fmt.Sprintf("/things/%s", t.ID)
	ThingSlice[t.ID] = t
	return &t, nil
}
