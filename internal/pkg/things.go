package pkg

// A Thing represents a... err... Thing
type Thing struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	ParentID    string `json:"parent_id,omitempty"`
	State       string `json:"state"`
	Additional  string `json:"additional,omitempty"`
}

// A ThingStore provides synchronisation for accessing Things
type ThingStore struct {
	setChan chan Thing
	getChan chan map[string]Thing
}

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
				t[k] = v
			}
		}
	}
}

// Get gets all Things
func (wis ThingStore) Get() map[string]Thing {
	return <-wis.getChan
}

// Set sets a Thing
func (wis ThingStore) Set(w Thing) {
	wis.setChan <- w
}
