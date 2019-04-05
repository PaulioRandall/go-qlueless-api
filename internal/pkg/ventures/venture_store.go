package ventures

import (
	"strconv"
	"sync"
)

// VentureStore represents a data store of Ventures
type VentureStore struct {
	mutex *sync.RWMutex
	items map[string]Venture
}

// NewVentureStore creates a new VentureStore.
func NewVentureStore() VentureStore {
	return VentureStore{
		mutex: &sync.RWMutex{},
		items: map[string]Venture{},
	}
}

// GetAll returns a slice of all Ventures currently held within the data store.
func (vs *VentureStore) GetAll() []Venture {
	vs.mutex.RLock()
	defer vs.mutex.RUnlock()

	r := []Venture{}
	for _, v := range vs.items {
		r = append(r, v)
	}

	return r
}

// GetAllAlive returns a slice of all Ventures currently held within the data
// store.
func (vs *VentureStore) GetAllAlive() []Venture {
	vs.mutex.RLock()
	defer vs.mutex.RUnlock()

	r := []Venture{}
	for _, v := range vs.items {
		if v.IsAlive {
			r = append(r, v)
		}
	}

	return r
}

// Get returns a specific Venture if found else the bool result will be false.
func (vs *VentureStore) Get(id string) (Venture, bool) {
	vs.mutex.RLock()
	defer vs.mutex.RUnlock()

	r, ok := vs.items[id]
	return r, ok
}

// Add adds a Venture to the data store assigning an unused ID.
func (vs *VentureStore) Add(new NewVenture) Venture {
	vs.mutex.Lock()
	defer vs.mutex.Unlock()

	ven := Venture{
		ID:          vs._genNewID(),
		Description: new.Description,
		OrderIDs:    new.OrderIDs,
		State:       new.State,
		IsAlive:     true,
		Extra:       new.Extra,
	}

	vs.items[ven.ID] = ven
	return ven
}

// Update updates Ventures within the data store. Only the Ventures in the slice
// returned were actually updated.
func (vs *VentureStore) Update(mv *ModVenture) []Venture {
	vs.mutex.Lock()
	defer vs.mutex.Unlock()

	r := []Venture{}

	for _, id := range mv.SplitIDs() {
		v, ok := vs.items[id]
		if !ok {
			continue
		}

		mv.ApplyMod(&v)
		vs.items[id] = v
		r = append(r, v)
	}

	return r
}

// _genNewID generates a new, unused, Venture ID
func (vs *VentureStore) _genNewID() string {
	ID := 1
	var r string

	for {
		r = strconv.Itoa(ID)
		_, ok := vs.items[r]
		if !ok {
			break
		}
		ID++
	}

	return r
}
