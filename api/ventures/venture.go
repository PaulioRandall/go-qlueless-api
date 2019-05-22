package ventures

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/PaulioRandall/go-cookies/cookies"
	"github.com/PaulioRandall/go-cookies/strlist"
	"github.com/PaulioRandall/go-qlueless-api/api/database"
)

// Venture represents a Venture, aka, project.
type Venture struct {
	ID           string `json:"id,omitempty"`
	LastModified int64  `json:"last_modified"`
	Description  string `json:"description"`
	Orders       string `json:"orders,omitempty"`
	State        string `json:"state"`
	Dead         bool   `json:"dead,omitempty"`
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
	ven.Orders = cookies.StripWhitespace(ven.Orders)
	ven.State = strings.TrimSpace(ven.State)
}

// Validate checks each field contains valid content returning a non-empty
// slice of human readable error messages detailing the violations found or an
// empty slice if all is well. These messages are suitable for returning to
// clients.
func (ven *Venture) Validate(isNew bool) []string {
	r := strlist.StrList{}

	if ven.Description == "" {
		r.Add("Ventures must have a description.")
	}

	if !isNew {
		if !cookies.IsUint(ven.ID) {
			r.Add("Ventures must have a positive integer ID.")
		}

		if ven.LastModified < 1 {
			r.Add("Ventures must have a last modified Unix date in milliseconds")
		}
	}

	if ven.Orders != "" {
		if !cookies.IsUintCSV(ven.Orders) {
			r.Add("Child Orders within a Venture must all be positive integers.")
		}
	}

	if ven.State == "" {
		r.Add("Ventures must have a state.")
	}

	return r.Slice()
}

// SplitOrders returns the IDs of the Orders as a slice.
func (ven *Venture) SplitOrders() []string {
	if ven.Orders == "" {
		return []string{}
	}
	return strings.Split(ven.Orders, ",")
}

// SetOrders sets the Orders CSV from a slice of Order IDs.
func (ven *Venture) SetOrders(ids []string) {
	ven.Orders = strings.Join(ids, ",")
}

// Update updates the Venture within the database.
func (ven *Venture) Update() error {
	stmt, err := database.Get().Prepare(`INSERT INTO venture (
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
		ven.Orders,
		ven.State,
		ven.Dead,
		ven.Extra)

	return err
}

// ByVenID is a slice of Ventures
type ByVenID []Venture

// Len implements from sort.Interface
func (bv ByVenID) Len() int {
	return len(bv)
}

// Swap implements from sort.Interface
func (bv ByVenID) Swap(i, j int) {
	bv[i], bv[j] = bv[j], bv[i]
}

// Less implements from sort.Interface
func (bv ByVenID) Less(i, j int) bool {
	return bv[i].ID < bv[j].ID
}
