package ventures

import (
	"strconv"
	"sync"
)

type VentureStore struct {
	mutex *sync.RWMutex
	items map[string]Venture
}

// NewVentureStore creates a new VentureStore
func NewVentureStore() VentureStore {
	return VentureStore{
		mutex: &sync.RWMutex{},
		items: map[string]Venture{},
	}
}

// GetAll returns a slice of all Ventures currently held within the data store
func (v *VentureStore) GetAll() []Venture {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	r := []Venture{}
	for _, v := range v.items {
		r = append(r, v)
	}

	return r
}

// GetAllAlive returns a slice of all Ventures currently held within the data store
func (v *VentureStore) GetAllAlive() []Venture {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	r := []Venture{}
	for _, v := range v.items {
		if v.IsAlive {
			r = append(r, v)
		}
	}

	return r
}

// Get returns a specific Venture if found else the bool result will be false
func (v *VentureStore) Get(id string) (Venture, bool) {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	r, ok := v.items[id]
	return r, ok
}

// Add adds a Venture to the data store assigning an unused ID
func (v *VentureStore) Add(new Venture) Venture {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	new.ID = v.genNewID()
	v.items[new.ID] = new
	return new
}

// genNewID generates a new, unused, Venture ID
func (v *VentureStore) genNewID() string {
	ID := 1
	var r string

	for {
		r = strconv.Itoa(ID)
		_, ok := v.items[r]
		if !ok {
			break
		}
		ID++
	}

	return r
}
