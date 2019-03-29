package ventures

import (
	"strings"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// Venture represents a Venture, aka, project.
type Venture struct {
	Description string `json:"description"`
	ID          string `json:"venture_id,omitempty"`
	OrderIDs    string `json:"order_ids,omitempty"`
	State       string `json:"state"`
	IsAlive     bool   `json:"is_alive"`
	Extra       string `json:"extra,omitempty"`
}

// Clean removes redundent whitespace from property values within a Venture
// except where whitespace is allowable.
func (ven *Venture) Clean() {
	ven.Description = strings.TrimSpace(ven.Description)
	ven.ID = strings.TrimSpace(ven.ID)
	ven.OrderIDs = StripWhitespace(ven.OrderIDs)
	ven.State = strings.TrimSpace(ven.State)
}

// Validate checks each field contains valid content returning a non-empty
// slice of human readable error messages detailing the violations found or an
// empty slice if all is well. These messages are suitable for returning to
// clients.
func (ven *Venture) Validate(isNew bool) []string {
	errMsgs := []string{}

	errMsgs = AppendIfEmpty(ven.Description, errMsgs,
		"Ventures must have a description")

	if !isNew {
		errMsgs = AppendIfNotPositiveInt(ven.ID, errMsgs,
			"Ventures must have a positive integer ID")
	}

	if ven.OrderIDs != "" {
		errMsgs = AppendIfNotPositiveIntCSV(ven.OrderIDs, errMsgs,
			"Child OrderIDs within a Venture must all be positive integers")
	}

	errMsgs = AppendIfEmpty(ven.State, errMsgs, "Ventures must have a state")
	return errMsgs
}

// SplitOrderIDs returns the IDs of the Orders as a slice.
func (ven *Venture) SplitOrderIDs() []string {
	if ven.OrderIDs == "" {
		return []string{}
	}
	return strings.Split(ven.OrderIDs, ",")
}

// SetOrderIDs sets the OrderIDs CSV from a slice of Order IDs.
func (ven *Venture) SetOrderIDs(ids []string) {
	ven.OrderIDs = strings.Join(ids, ",")
}
