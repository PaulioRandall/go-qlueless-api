package ventures

import (
	"fmt"
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
		"Ventures must have a description.")

	if !isNew {
		errMsgs = AppendIfNotPositiveInt(ven.ID, errMsgs,
			"Ventures must have a positive integer ID.")
	}

	if ven.OrderIDs != "" {
		errMsgs = AppendIfNotPositiveIntCSV(ven.OrderIDs, errMsgs,
			"Child OrderIDs within a Venture must all be positive integers.")
	}

	errMsgs = AppendIfEmpty(ven.State, errMsgs, "Ventures must have a state.")
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

// VentureUpdate represents an update to a Venture.
type VentureUpdate struct {
	Props  string  `json:"set"`
	Values Venture `json:"values"`
}

// SplitProp returns the property names of the properties to update.
func (vu *VentureUpdate) SplitProps() []string {
	if vu.Props == "" {
		return []string{}
	}
	return strings.Split(vu.Props, ",")
}

// _validateProps is a private function that checks the properties declared for
// change are valid and the property value for each is valid. Returned is the
// input slice of human readable error messages with the violations found
// appended to it. These messages are suitable for returning to clients.
func (vu *VentureUpdate) _validateProps(errMsgs []string) []string {
	for _, p := range vu.SplitProps() {
		switch p {
		case "is_alive", "extra":
		case "description":
			errMsgs = AppendIfEmpty(vu.Values.Description, errMsgs,
				"Ventures must have a description.")
		case "state":
			errMsgs = AppendIfEmpty(vu.Values.State, errMsgs,
				"Ventures must have a state.")
		case "order_ids":
			errMsgs = AppendIfNotPositiveIntCSV(vu.Values.OrderIDs, errMsgs,
				"The list of Order IDs within a Venture must be an integer CSV")
		default:
			errMsgs = append(errMsgs,
				fmt.Sprintf("Can't update unknown or immutable property '%s'.", p))
		}
	}

	return errMsgs
}

// Validate checks each field contains valid content returning a non-empty
// slice of human readable error messages detailing the violations found or an
// empty slice if all is well. These messages are suitable for returning to
// clients.
func (vu *VentureUpdate) Validate() []string {
	errMsgs := []string{}

	errMsgs = AppendIfEmpty(vu.Values.ID, errMsgs,
		"'values.venture_id' must be supplied so I know which Venture to update.")

	errMsgs = AppendIfEmpty(vu.Props, errMsgs,
		"Some properties must be 'set' for any updating to take place.")

	errMsgs = vu._validateProps(errMsgs)
	return errMsgs
}
