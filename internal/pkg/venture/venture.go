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
// except where whitespace is allowable.
func (ven *Venture) Clean() {
	ven.Description = strings.TrimSpace(ven.Description)
	ven.VentureID = strings.TrimSpace(ven.VentureID)
	ven.OrderIDs = StripWhitespace(ven.OrderIDs)
	ven.State = strings.TrimSpace(ven.State)
}

// Validate checks each field contains valid content returning a non-empty
// slice of human readable error messages detailing the violations found or an
// empty slice if all is well. These messages are suitable for returning to
// clients.
func (ven *Venture) Validate() []string {
	errMsgs := []string{}

	errMsgs = AppendIfEmpty(ven.Description, errMsgs,
		"Ventures must have a description")
	errMsgs = AppendIfNotPositiveInt(ven.VentureID, errMsgs,
		"Ventures must have a positive integer ID")

	if ven.OrderIDs != "" {
		errMsgs = AppendIfNotPositiveIntCSV(ven.OrderIDs, errMsgs,
			"Child OrderIDs within a Venture must all be positive integers")
	}

	errMsgs = AppendIfEmpty(ven.State, errMsgs, "Ventures must have a state")
	return errMsgs
}
