package pkg

import (
	"fmt"
	"sync"
)

// A ThingStore provides synchronisation for accessing Things
type ThingStore struct {
	mutex  *sync.RWMutex
	things map[int]Thing
}

// NewThingStore creates a new ThingStore
func NewThingStore() ThingStore {
	return ThingStore{
		mutex:  &sync.RWMutex{},
		things: map[int]Thing{},
	}
}

// GetAll returns the map of all Things currently held within the data store
func (ts ThingStore) GetAll() map[int]Thing {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()
	return ts.things
}

// GetAllAlive returns a slice of all Things which are not dead
func (ts ThingStore) GetAllAlive() []Thing {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()

	r := []Thing{}
	for _, v := range ts.things {
		if !v.IsDead {
			r = append(r, v)
		}
	}
	return r
}

// Get returns a specific Thing or nil if the Thing does not exist
func (ts ThingStore) Get(id int) Thing {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()
	return ts.things[id]
}

// Add adds a Thing to the data store assigning an unused ID
func (ts ThingStore) Add(t Thing) Thing {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	t.ID = ts.genNewID()
	t.Self = fmt.Sprintf("/things/%d", t.ID)

	ts.things[t.ID] = t
	return t
}

// genNewID generates a new, unused, Thing ID
func (ts ThingStore) genNewID() int {
	ID := 0
	for k, _ := range ts.things {
		if k > ID {
			ID = k
		}
	}

	ID++
	return ID
}
