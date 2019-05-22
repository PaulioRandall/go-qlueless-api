package ventures

import (
	"database/sql"
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"github.com/PaulioRandall/go-cookies/cookies"
	"github.com/PaulioRandall/go-cookies/strlist"
	"github.com/PaulioRandall/go-qlueless-api/api/database"
)

// NewVenture represents a new Venture.
type NewVenture struct {
	Description string `json:"description"`
	Orders      string `json:"orders"`
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
	nv.Orders = cookies.StripWhitespace(nv.Orders)
	nv.State = strings.TrimSpace(nv.State)
}

// Validate checks each field contains valid content returning a non-empty
// slice of human readable error messages detailing the violations found or an
// empty slice if all is well. These messages are suitable for returning to
// clients.
func (nv *NewVenture) Validate() []string {
	r := strlist.StrList{}

	if nv.Description == "" {
		r.Add("Ventures must have a description.")
	}

	if nv.Orders != "" && !cookies.IsUintCSV(nv.Orders) {
		r.Add("Child OrderIDs within a Venture must all be positive integers.")
	}

	if nv.State == "" {
		r.Add("Ventures must have a state.")
	}

	return r.Slice()
}

// Insert inserts the NewVenture into the database
func (nv *NewVenture) Insert() (ven *Venture, ok bool) {
	ok = false

	id, err := findNextID()
	if cookies.LogIfErr(err) {
		return
	}

	stmt, err := database.Get().Prepare(`INSERT INTO venture (
		id, description, order_ids, state, extra
	) VALUES (
		?, ?, ?, ?, ?
	);`)

	if stmt != nil {
		defer stmt.Close()
	}

	if cookies.LogIfErr(err) {
		return
	}

	_, err = nv.execInsert(id, stmt)
	if cookies.LogIfErr(err) {
		return
	}

	ven, err = QueryFor(id)
	if cookies.LogIfErr(err) {
		return
	}

	ok = true
	return
}

// findNextID returns the next free Venture ID.
func findNextID() (result string, err error) {
	result = ""
	stmt, err := database.Get().Prepare(`SELECT COALESCE(MAX(id), 0) FROM venture;`)

	if stmt != nil {
		defer stmt.Close()
	}

	if err != nil {
		return
	}

	var id int64
	err = stmt.QueryRow().Scan(&id)
	if err != nil {
		return
	}

	id++
	result = strconv.FormatInt(id, 10)
	return
}

// execInsert is a file private function that executes the supplied insert
// statement
func (nv *NewVenture) execInsert(id string, stmt *sql.Stmt) (ven *Venture, err error) {
	ven = &Venture{
		ID:          id,
		Description: nv.Description,
		Orders:      nv.Orders,
		State:       nv.State,
		Extra:       nv.Extra,
	}

	_, err = stmt.Exec(ven.ID,
		ven.Description,
		ven.Orders,
		ven.State,
		ven.Extra)

	if err != nil {
		ven = nil
	}

	return
}
