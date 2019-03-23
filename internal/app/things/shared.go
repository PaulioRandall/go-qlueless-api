package things

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// parseThingID parses the ID in the request
func parseThingID(idStr string, res *http.ResponseWriter, req *http.Request) (int, bool) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		r := ReplyMeta{
			Message: fmt.Sprintf("ID '%s' could not be parsed to an integer", idStr),
		}
		Write4XXReply(res, req, 400, r)
		return 0, false
	}
	return id, true
}

// findThing finds the Thing with the specified ID
func findThing(id int, res *http.ResponseWriter, req *http.Request) (Thing, bool) {
	t := Things.Get(id)
	if t.ID < 1 || t.IsDead {
		r := ReplyMeta{
			Message: fmt.Sprintf("Thing %d not found", id),
		}
		Write4XXReply(res, req, 404, r)
		return Thing{}, false
	}
	return t, true
}

// decodeThing decodes a thing from a Request.Body
func decodeThing(res *http.ResponseWriter, req *http.Request) (Thing, bool) {
	var t Thing
	d := json.NewDecoder(req.Body)
	err := d.Decode(&t)
	if err != nil {
		r := ReplyMeta{
			Message: "Unable to decode request body into a Thing",
		}
		Write4XXReply(res, req, 400, r)
		return Thing{}, false
	}
	return t, true
}

// checkThing cleans and validates the Thing
func checkThing(t Thing, res *http.ResponseWriter, req *http.Request) (Thing, bool) {
	CleanThing(&t)
	e := ValidateThing(&t, true)
	if e != nil {
		r := ReplyMeta{
			Message: strings.Join(e, " "),
		}
		Write4XXReply(res, req, 400, r)
		return Thing{}, false
	}
	return t, true
}
