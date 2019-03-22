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
	idStr := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		r := Reply4XX{
			Res:     res,
			Req:     req,
			Message: fmt.Sprintf("ID '%s' could not be parsed to an integer", idStr),
		}
		Write4XXReply(400, &r)
		return
	}

	t := Things.Get(id)
	if t.ID < 1 || t.IsDead {
		r := Reply4XX{
			Res:     res,
			Req:     req,
			Message: fmt.Sprintf("Thing %d not found", id),
		}
		Write4XXReply(404, &r)
		return
	}

	m := fmt.Sprintf("Found Thing with ID %d", id)
	data := PrepResponseData(req, t, m)
	WriteReply(res, req, data)
}
