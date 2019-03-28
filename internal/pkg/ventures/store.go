package ventures

import "sync"

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
