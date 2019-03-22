package things

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// ThingHandler handles requests to do with a specific Thing
func ThingHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	switch req.Method {
	case "GET":
		getThing(&res, req)
	case "HEAD":
		WriteEmptyReply(&res)
	case "OPTIONS":
		WriteEmptyReply(&res)
	default:
		MethodNotAllowed(&res, req)
	}
}

// GetThing generates responses for requests for a single Thing
func getThing(res *http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, ok := parseID(idStr, res, req)
	if !ok {
		return
	}

	t, ok := findThing(id, res, req)
	if !ok {
		return
	}

	m := fmt.Sprintf("Found Thing with ID %d", id)
	data := PrepResponseData(req, t, m)
	WriteReply(res, req, data)
}

// parseID parses the ID in the request
func parseID(idStr string, res *http.ResponseWriter, req *http.Request) (int, bool) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		r := Reply4XX{
			Res:     res,
			Req:     req,
			Message: fmt.Sprintf("ID '%s' could not be parsed to an integer", idStr),
		}
		Write4XXReply(400, &r)
		return 0, false
	}
	return id, true
}

// findThing finds the Thing with the specified ID
func findThing(id int, res *http.ResponseWriter, req *http.Request) (Thing, bool) {
	t := Things.Get(id)
	if t.ID < 1 || t.IsDead {
		r := Reply4XX{
			Res:     res,
			Req:     req,
			Message: fmt.Sprintf("Thing %d not found", id),
		}
		Write4XXReply(404, &r)
		return Thing{}, false
	}
	return t, true
}
