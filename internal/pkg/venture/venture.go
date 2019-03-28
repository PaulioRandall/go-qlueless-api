package venture

import (
	"strings"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

type Venture struct {
	Description string `json:"description"`
	VentureID   string `json:"venture_id,omitempty"`
	OrderIDs    string `json:"order_ids,omitempty"`
	State       string `json:"state"`
	IsAlive     bool   `json:"is_alive"`
	Extra       string `json:"extra,omitempty"`
}

// Clean removes redundent whitespace from property values within a Venture
// except where whitespace is allowable
func (ven *Venture) Clean() {
	ven.Description = strings.TrimSpace(ven.Description)
	ven.VentureID = strings.TrimSpace(ven.VentureID)
	ven.OrderIDs = StripWhitespace(ven.OrderIDs)
	ven.State = strings.TrimSpace(ven.State)
}
