package ventures

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	u "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/utils"
)

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
	vu.Props = u.StripWhitespace(vu.Props)
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
			errMsgs = u.AppendIfEmpty(vu.Values.Description, errMsgs,
				"Ventures must have a description.")
		case "state":
			errMsgs = u.AppendIfEmpty(vu.Values.State, errMsgs,
				"Ventures must have a state.")
		case "order_ids":
			errMsgs = u.AppendIfNotPositiveIntCSV(vu.Values.OrderIDs, errMsgs,
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

	errMsgs = u.AppendIfEmpty(vu.Values.ID, errMsgs,
		"'values.venture_id' must be supplied so I know which Venture to update.")

	errMsgs = u.AppendIfEmpty(vu.Props, errMsgs,
		"Some properties must be 'set' for any updating to take place.")

	errMsgs = vu._validateProps(errMsgs)
	return errMsgs
}
