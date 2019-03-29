package things

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// findThing finds the Thing with the specified ID
func findThing(id string, res *http.ResponseWriter, req *http.Request) (Thing, bool) {
	t := Things.Get(id)
	if t.ID == "" || t.IsDead {
		r := WrappedReply{
			Message: fmt.Sprintf("Thing '%s' not found", id),
		}
		Write4XXReply(res, req, 404, r)
		return Thing{}, false
	}
	return t, true
}

// decodeThing decodes a Thing from a Request.Body
func decodeThing(res *http.ResponseWriter, req *http.Request) (Thing, bool) {
	var t Thing
	d := json.NewDecoder(req.Body)
	err := d.Decode(&t)
	if err != nil {
		r := WrappedReply{
			Message: "Unable to decode request body into a Thing",
		}
		Write4XXReply(res, req, 400, r)
		return Thing{}, false
	}
	return t, true
}

// checkThing cleans and validates the Thing
func checkThing(t Thing, res *http.ResponseWriter, req *http.Request) (Thing, bool) {
	t.Clean()
	e := t.Validate(true)
	if e != nil {
		r := WrappedReply{
			Message: strings.Join(e, " "),
		}
		Write4XXReply(res, req, 400, r)
		return Thing{}, false
	}
	return t, true
}
