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
func (nv *NewVenture) Clean() {
	nv.Description = strings.TrimSpace(nv.Description)
	nv.OrderIDs = u.StripWhitespace(nv.OrderIDs)
	nv.State = strings.TrimSpace(nv.State)
}

// Validate checks each field contains valid content returning a non-empty
// slice of human readable error messages detailing the violations found or an
// empty slice if all is well. These messages are suitable for returning to
// clients.
func (nv *NewVenture) Validate() []string {
	errMsgs := []string{}

	errMsgs = u.AppendIfEmpty(nv.Description, errMsgs,
		"Ventures must have a description.")

	if nv.OrderIDs != "" {
		errMsgs = u.AppendIfNotPositiveIntCSV(nv.OrderIDs, errMsgs,
			"Child OrderIDs within a Venture must all be positive integers.")
	}

	errMsgs = u.AppendIfEmpty(nv.State, errMsgs, "Ventures must have a state.")
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

	return nv._execInsert(stmt)
}

// _execInsert is a file private function that executes the supplied insert
// statement
func (nv *NewVenture) _execInsert(stmt *sql.Stmt) (*Venture, error) {
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
