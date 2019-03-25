package things

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// ThingHandler handles requests to do with a specific Thing
func ThingHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	switch req.Method {
	case "GET":
		get_Thing(&res, req)
	case "HEAD":
		WriteEmptyJSONReply(&res, "")
	case "OPTIONS":
		WriteEmptyJSONReply(&res, "")
	default:
		MethodNotAllowed(&res, req)
	}
}

// get_Thing generates responses for requests for a single Thing
func get_Thing(res *http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	t, ok := findThing(id, res, req)
	if !ok {
		return
	}

	m := fmt.Sprintf("Found Thing with ID %s", id)
	data := PrepResponseData(req, t, m)
	WriteJSONReply(res, req, data, "")
}
