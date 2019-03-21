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
		GetThing(&res, req)
	case "HEAD":
		WriteEmptyReply(&res)
	case "OPTIONS":
		WriteEmptyReply(&res)
	default:
		MethodNotAllowed(&res, req)
	}
}

// GetThing generates responses for requests for a single Thing
func GetThing(res *http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	t := Things.Get(id)

	if t.ID == "" || t.IsDead {
		r := Reply4XX{
			Res:     res,
			Req:     req,
			Message: fmt.Sprintf("Thing %v not found", id),
		}
		Write4XXReply(404, &r)
		return
	}

	m := fmt.Sprintf("Found Thing with ID %s", id)
	data := PrepResponseData(req, t, m)
	WriteReply(res, req, data)
}
