package ventures

import (
	"database/sql"
	"encoding/json"
	"io"
	"strings"

	u "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/utils"
)

// Venture represents a Venture, aka, project.
type Venture struct {
	ID           string `json:"venture_id,omitempty"`
	LastModified int64  `json:"last_modified"`
	Description  string `json:"description"`
	OrderIDs     string `json:"order_ids,omitempty"`
	State        string `json:"state"`
	IsDead       bool   `json:"is_dead,omitempty"`
	Extra        string `json:"extra,omitempty"`
}

// DecodeVenture decodes a Venture from data obtained via a Reader.
func DecodeVenture(r io.Reader) (Venture, error) {
	var v Venture
	d := json.NewDecoder(r)
	err := d.Decode(&v)
	return v, err
}

// DecodeVentureSlice decodes a slice of Ventures from data obtained via a
// Reader.
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
	ven.OrderIDs = u.StripWhitespace(ven.OrderIDs)
	ven.State = strings.TrimSpace(ven.State)
}

// Validate checks each field contains valid content returning a non-empty
// slice of human readable error messages detailing the violations found or an
// empty slice if all is well. These messages are suitable for returning to
// clients.
func (ven *Venture) Validate(isNew bool) []string {
	errMsgs := []string{}

	errMsgs = u.AppendIfEmpty(ven.Description, errMsgs,
		"Ventures must have a description.")

	if !isNew {
		errMsgs = u.AppendIfNotPositiveInt(ven.ID, errMsgs,
			"Ventures must have a positive integer ID.")

		if ven.LastModified < 1 {
			errMsgs = append(errMsgs, "Ventures must have a last modified Unix date in milliseconds")
		}
	}

	if ven.OrderIDs != "" {
		errMsgs = u.AppendIfNotPositiveIntCSV(ven.OrderIDs, errMsgs,
			"Child OrderIDs within a Venture must all be positive integers.")
	}

	errMsgs = u.AppendIfEmpty(ven.State, errMsgs, "Ventures must have a state.")
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

// Update updates the Venture within the database.
//
// @UNTESTED
func (ven *Venture) Update(db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO venture (
		id, description, order_ids, state, is_dead, extra
	) VALUES (
		?, ?, ?, ?, ?, ?
	);`)

	if stmt != nil {
		defer stmt.Close()
	}

	if err != nil {
		return err
	}

	_, err = stmt.Exec(ven.ID,
		ven.Description,
		ven.OrderIDs,
		ven.State,
		ven.IsDead,
		ven.Extra)

	return err
}
