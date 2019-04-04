package ventures

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	u "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/utils"
)

// ModVenture represents an update to a Venture.
type ModVenture struct {
	IDs    string  `json:"ids"`
	Props  string  `json:"set"`
	Values Venture `json:"values"`
}

// DecodeModVenture decodes a ModVenture from data obtained via a Reader.
func DecodeModVenture(r io.Reader) (ModVenture, error) {
	var mv ModVenture
	d := json.NewDecoder(r)
	err := d.Decode(&mv)
	return mv, err
}

// SplitProp returns the property names of the properties to update.
func (mv *ModVenture) SplitProps() []string {
	if mv.Props == "" {
		return []string{}
	}
	return strings.Split(mv.Props, ",")
}

// Clean cleans up the ModVenture by removing whitespace where applicable.
func (mv *ModVenture) Clean() {
	mv.Props = u.StripWhitespace(mv.Props)
	mv.Values.Clean()
}

// _validateProps is a private function that checks the properties declared for
// change are valid and the property value for each is valid. Returned is the
// input slice of human readable error messages with the violations found
// appended to it. These messages are suitable for returning to clients.
func (mv *ModVenture) _validateProps(errMsgs []string) []string {
	for _, prop := range mv.SplitProps() {
		switch prop {
		case "is_alive", "extra":
		case "description":
			errMsgs = u.AppendIfEmpty(mv.Values.Description, errMsgs,
				"Ventures must have a description.")
		case "state":
			errMsgs = u.AppendIfEmpty(mv.Values.State, errMsgs,
				"Ventures must have a state.")
		case "order_ids":
			errMsgs = u.AppendIfNotPositiveIntCSV(mv.Values.OrderIDs, errMsgs,
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
func (mv *ModVenture) Validate() []string {
	errMsgs := []string{}

	errMsgs = u.AppendIfEmpty(mv.Values.ID, errMsgs,
		"'values.venture_id' must be supplied so I know which Venture to update.")

	errMsgs = u.AppendIfEmpty(mv.Props, errMsgs,
		"Some properties must be 'set' for any updating to take place.")

	errMsgs = mv._validateProps(errMsgs)
	return errMsgs
}

// ApplyMod applies the modifications to the supplied Venture only touching
// those properties the user has specified
//
// @UNTESTED
func (mv *ModVenture) ApplyMod(ven *Venture) {
	mod := mv.Values
	for _, p := range mv.SplitProps() {
		switch p {
		case "description":
			ven.Description = mod.Description
		case "order_ids":
			ven.OrderIDs = mod.OrderIDs
		case "state":
			ven.State = mod.State
		case "is_alive":
			ven.IsAlive = mod.IsAlive
		case "extra":
			ven.Extra = mod.Extra
		}
	}
}
