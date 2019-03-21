package pkg

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

// A ThingStore provides synchronisation for accessing Things
type ThingStore struct {
	mutex  sync.RWMutex
	things map[string]Thing
}

// NewThingStore creates a new ThingStore
func NewThingStore() ThingStore {
	ts := ThingStore{
		mutex:  sync.RWMutex{},
		things: map[string]Thing{},
	}
	return ts
}

// GetAll returns the map of all Things currently held within the data store
func (ts ThingStore) GetAll() map[string]Thing {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()
	return ts.things
}

// Get returns a specific Thing or nil if the Thing does not exist
func (ts ThingStore) Get(id string) Thing {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()
	return ts.things[id]
}

// Add adds a Thing to the data store assigning an unused ID
func (ts ThingStore) Add(t Thing) (Thing, error) {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	ID, err := ts.genNewID()
	if err != nil {
		return Thing{}, err
	}

	t.ID = ID
	t.Self = fmt.Sprintf("/things/%s", t.ID)

	ts.things[ID] = t
	return t, nil
}

// genNewID generates a new, unused, Thing ID
func (ts ThingStore) genNewID() (string, error) {
	newID := 0
	for k, _ := range ts.things {
		ID, err := strconv.Atoi(k)
		if LogIfErr(err) {
			return "", errors.New("[BUG] An unparsable ID exists within the data store")
		}
		if ID > newID {
			newID = ID
		}
	}
	newID++
	return strconv.Itoa(newID), nil
}
