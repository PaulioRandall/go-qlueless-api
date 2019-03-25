package pkg

import (
	"strconv"
	"sync"
)

// A ThingStore provides synchronisation for accessing Things
type ThingStore struct {
	mutex  *sync.RWMutex
	things map[string]Thing
}

// NewThingStore creates a new ThingStore
func NewThingStore() ThingStore {
	return ThingStore{
		mutex:  &sync.RWMutex{},
		things: map[string]Thing{},
	}
}

// GetAll returns the map of all Things currently held within the data store
func (ts ThingStore) GetAll() map[string]Thing {
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
func (ts ThingStore) Get(id string) Thing {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()
	return ts.things[id]
}

// Add adds a Thing to the data store assigning an unused ID
func (ts ThingStore) Add(t Thing) Thing {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	t.ID = ts.genNewID()

	ts.things[t.ID] = t
	return t
}

// genNewID generates a new, unused, Thing ID
func (ts ThingStore) genNewID() string {
	ID := 1
	var r string

	for {
		r = strconv.Itoa(ID)
		_, ok := ts.things[r]
		if !ok {
			break
		}
		ID++
	}

	return r
}
