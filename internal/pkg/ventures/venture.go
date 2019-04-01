package ventures

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	p "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
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

// DecodeVenture decodes a Venture from data obtained via a Reader
func DecodeVenture(r io.Reader) (Venture, error) {
	var v Venture
	d := json.NewDecoder(r)
	err := d.Decode(&v)
	return v, err
}

// DecodeVentureSlice decodes a slice of Ventures from data obtained via a
// Reader
func DecodeVentureSlice(r io.Reader) ([]Venture, error) {
	var v []Venture
	d := json.NewDecoder(r)
	err := d.Decode(&v)
	return v, err
}

// Clean removes redundent whitespace from property values within a Venture
// except where whitespace is allowable.
func (ven *Venture) Clean() {
	ven.Description = strings.TrimSpace(ven.Description)
	ven.ID = strings.TrimSpace(ven.ID)
	ven.OrderIDs = p.StripWhitespace(ven.OrderIDs)
	ven.State = strings.TrimSpace(ven.State)
}

// Validate checks each field contains valid content returning a non-empty
// slice of human readable error messages detailing the violations found or an
// empty slice if all is well. These messages are suitable for returning to
// clients.
func (ven *Venture) Validate(isNew bool) []string {
	errMsgs := []string{}

	errMsgs = p.AppendIfEmpty(ven.Description, errMsgs,
		"Ventures must have a description.")

	if !isNew {
		errMsgs = p.AppendIfNotPositiveInt(ven.ID, errMsgs,
			"Ventures must have a positive integer ID.")
	}

	if ven.OrderIDs != "" {
		errMsgs = p.AppendIfNotPositiveIntCSV(ven.OrderIDs, errMsgs,
			"Child OrderIDs within a Venture must all be positive integers.")
	}

	errMsgs = p.AppendIfEmpty(ven.State, errMsgs, "Ventures must have a state.")
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

// DecodeVentureUpdate decodes a VentureUpdate from data obtained via a Reader
func DecodeVentureUpdate(r io.Reader) (VentureUpdate, error) {
	var vu VentureUpdate
	d := json.NewDecoder(r)
	err := d.Decode(&vu)
	return vu, err
}

// SplitProp returns the property names of the properties to update.
func (vu *VentureUpdate) SplitProps() []string {
	if vu.Props == "" {
		return []string{}
	}
	return strings.Split(vu.Props, ",")
}

// Clean cleans up the Venture update by removing whitespace where applicable
func (vu *VentureUpdate) Clean() {
	vu.Props = p.StripWhitespace(vu.Props)
	vu.Values.Clean()
}

// _validateProps is a private function that checks the properties declared for
// change are valid and the property value for each is valid. Returned is the
// input slice of human readable error messages with the violations found
// appended to it. These messages are suitable for returning to clients.
func (vu *VentureUpdate) _validateProps(errMsgs []string) []string {
	for _, prop := range vu.SplitProps() {
		switch prop {
		case "is_alive", "extra":
		case "description":
			errMsgs = p.AppendIfEmpty(vu.Values.Description, errMsgs,
				"Ventures must have a description.")
		case "state":
			errMsgs = p.AppendIfEmpty(vu.Values.State, errMsgs,
				"Ventures must have a state.")
		case "order_ids":
			errMsgs = p.AppendIfNotPositiveIntCSV(vu.Values.OrderIDs, errMsgs,
				"The list of Order IDs within a Venture must be an integer CSV")
		default:
			errMsgs = append(errMsgs,
				fmt.Sprintf("Can't update unknown or immutable property '%s'.", prop))
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

	errMsgs = p.AppendIfEmpty(vu.Values.ID, errMsgs,
		"'values.venture_id' must be supplied so I know which Venture to update.")

	errMsgs = p.AppendIfEmpty(vu.Props, errMsgs,
		"Some properties must be 'set' for any updating to take place.")

	errMsgs = vu._validateProps(errMsgs)
	return errMsgs
}
