package ventures

import (
	"database/sql"
	"encoding/json"
	"io"
	"strconv"
	"strings"

	u "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/utils"
)

// NewVenture represents a new Venture.
type NewVenture struct {
	Description string `json:"description"`
	OrderIDs    string `json:"order_ids"`
	State       string `json:"state"`
	Extra       string `json:"extra"`
}

// DecodeNewVenture decodes a NewVenture from data obtained via a Reader
func DecodeNewVenture(r io.Reader) (NewVenture, error) {
	var v NewVenture
	d := json.NewDecoder(r)
	err := d.Decode(&v)
	return v, err
}

// Clean removes redundent whitespace from property values within a Venture
// except where whitespace is allowable.
func (ven *NewVenture) Clean() {
	ven.Description = strings.TrimSpace(ven.Description)
	ven.OrderIDs = u.StripWhitespace(ven.OrderIDs)
	ven.State = strings.TrimSpace(ven.State)
}

// Validate checks each field contains valid content returning a non-empty
// slice of human readable error messages detailing the violations found or an
// empty slice if all is well. These messages are suitable for returning to
// clients.
func (ven *NewVenture) Validate() []string {
	errMsgs := []string{}

	errMsgs = u.AppendIfEmpty(ven.Description, errMsgs,
		"Ventures must have a description.")

	if ven.OrderIDs != "" {
		errMsgs = u.AppendIfNotPositiveIntCSV(ven.OrderIDs, errMsgs,
			"Child OrderIDs within a Venture must all be positive integers.")
	}

	errMsgs = u.AppendIfEmpty(ven.State, errMsgs, "Ventures must have a state.")
	return errMsgs
}

// Insert inserts the NewVenture into the database
//
// @UNTESTED
func (nv *NewVenture) Insert(db *sql.DB) (*Venture, error) {
	stmt, err := db.Prepare(`INSERT INTO venture (
		description, order_ids, state, is_alive, extra
	) VALUES (
		?, ?, ?, ?, ?
	);`)

	if stmt != nil {
		defer stmt.Close()
	}

	if err != nil {
		return nil, err
	}

	ven := Venture{
		Description: nv.Description,
		OrderIDs:    nv.OrderIDs,
		State:       nv.State,
		IsAlive:     true,
		Extra:       nv.Extra,
	}

	res, err := stmt.Exec(ven.Description,
		ven.OrderIDs,
		ven.State,
		ven.IsAlive,
		ven.Extra)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	ven.ID = strconv.FormatInt(id, 10)
	return &ven, nil
}
