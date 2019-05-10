package ventures

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

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

// SplitIDs returns the IDs of the Ventures to update as a slice.
func (mv *ModVenture) SplitIDs() []string {
	if mv.IDs == "" {
		return []string{}
	}
	return strings.Split(mv.IDs, ",")
}

// SplitProps returns the property names of the properties to update.
func (mv *ModVenture) SplitProps() []string {
	if mv.Props == "" {
		return []string{}
	}
	return strings.Split(mv.Props, ",")
}

// Clean cleans up the ModVenture by removing whitespace where applicable.
func (mv *ModVenture) Clean() {
	mv.IDs = u.StripWhitespace(mv.IDs)
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
		case "dead", "extra":
		case "description":
			errMsgs = u.AppendIfEmpty(mv.Values.Description, errMsgs,
				"Ventures must have a description.")
		case "state":
			errMsgs = u.AppendIfEmpty(mv.Values.State, errMsgs,
				"Ventures must have a state.")
		case "orders":
			errMsgs = u.AppendIfNotPositiveIntCSV(mv.Values.Orders, errMsgs,
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

	errMsgs = u.AppendIfNotPositiveIntCSV(mv.IDs, errMsgs,
		"'ids' must be supplied so the Ventures to update can be determined.")

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
		ven.LastModified = u.ToUnixMilli(time.Now())
		switch p {
		case "description":
			ven.Description = mod.Description
		case "orders":
			ven.Orders = mod.Orders
		case "state":
			ven.State = mod.State
		case "dead":
			ven.Dead = mod.Dead
		case "extra":
			ven.Extra = mod.Extra
		}
	}
}

// Update pushes the modification of changes to the database.
//
// @UNTESTED
func (mv *ModVenture) Update(db *sql.DB) ([]Venture, bool) {

	ids := mv.SplitIDs()
	args := make([]interface{}, len(ids))
	for i, _ := range ids {
		args[i] = ids[i]
	}

	vens, err := QueryMany(db, args)
	if u.LogIfErr(err) {
		return nil, false
	}

	ok := mv._insertEach(db, vens)
	if !ok {
		return nil, false
	}

	return vens, true
}

// _insertEach is a file private function that performs the actual SQL operation
// of pushing modifications to the database.
func (mv *ModVenture) _insertEach(db *sql.DB, vens []Venture) bool {

	stmt, err := db.Prepare(`INSERT INTO venture
			(id, description, order_ids, state, is_dead, extra)
		VALUES
			(?, ?, ?, ?, ?, ?);`)

	if stmt != nil {
		defer stmt.Close()
	}

	if u.LogIfErr(err) {
		return false
	}

	return mv._execStmtForEach(stmt, vens)
}

// _execStmtForEach executes the insert statment provided for each Venture
// provided.
func (mv *ModVenture) _execStmtForEach(stmt *sql.Stmt, vens []Venture) bool {
	for i, _ := range vens {

		ven := &vens[i]
		mv.ApplyMod(ven)

		_, err := stmt.Exec(ven.ID,
			ven.Description,
			ven.Orders,
			ven.State,
			ven.Dead,
			ven.Extra)

		if u.LogIfErr(err) {
			return false
		}
	}

	return true
}
