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

// _updateVenture is a file private function that updates a Venture with the
// changes defined within the supplied venture update structure.
func (vs *VentureStore) _updateVenture(v Venture, vu VentureUpdate) Venture {
	u := vu.Values
	for _, p := range vu.SplitProps() {
		switch p {
		case "description":
			v.Description = u.Description
		case "order_ids":
			v.OrderIDs = u.OrderIDs
		case "state":
			v.State = u.State
		case "is_alive":
			v.IsAlive = u.IsAlive
		case "extra":
			v.Extra = u.Extra
		}
	}
	return v
}

// Update updates a Venture within the data store. If false is returned then
// the item does not currently exist within the data store.
func (vs *VentureStore) Update(vu VentureUpdate) (Venture, bool) {

	v, ok := vs.Get(vu.Values.ID)
	if !ok {
		return Venture{}, false
	}

	v = vs._updateVenture(v, vu)

	vs.mutex.Lock()
	defer vs.mutex.Unlock()

	vs.items[v.ID] = v
	return v, true
}

// Delete removes a Venture from within the data store. If false is returned
// then the item does not currently exist.
func (vs *VentureStore) Delete(id string) (Venture, bool) {
	vs.mutex.Lock()
	defer vs.mutex.Unlock()

	v, ok := vs.items[id]
	if ok {
		delete(vs.items, id)
		return v, true
	}

	return Venture{}, false
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
