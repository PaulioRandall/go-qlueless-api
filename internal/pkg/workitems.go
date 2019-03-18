package pkg

// A WorkItem represents and is a genralisation of orders and batches
type WorkItem struct {
	Description      string `json:"description"`
	WorkItemID       string `json:"work_item_id"`
	ParentWorkItemID string `json:"parent_work_item_id,omitempty"`
	TagID            string `json:"tag_id"`
	StatusID         string `json:"status_id"`
	Additional       string `json:"additional,omitempty"`
}

// A WorkItemStore provides synchronisation for accessing WorkItems
type WorkItemStore struct {
	setChan chan WorkItem
	getChan chan map[string]WorkItem
}

// NewWorkItemStore creates a new WorkItemStore
func NewWorkItemStore() WorkItemStore {
	wis := WorkItemStore{
		setChan: make(chan WorkItem),
		getChan: make(chan map[string]WorkItem),
	}
	go wis.mux()
	return wis
}

// mux should be invoked as a goroutine; it forever loops handling incoming
// channel communications sequentially
func (wis WorkItemStore) mux() {
	m := make(map[string]WorkItem)
	var r map[string]WorkItem
	for {
		select {
		case w := <-wis.setChan:
			m[w.WorkItemID] = w
		case wis.getChan <- r:
			t := make(map[string]WorkItem)
			for k, v := range m {
				t[k] = v
			}
		}
	}
}

// Get gets all WorkItems
func (wis WorkItemStore) Get() map[string]WorkItem {
	return <-wis.getChan
}

// Set sets a WorkItem
func (wis WorkItemStore) Set(w WorkItem) {
	wis.setChan <- w
}
