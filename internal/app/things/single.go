package things

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// SingleThingHandler handles requests for a specific things currently
// within the service
func SingleThingHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	id := mux.Vars(req)["id"]
	o, ok := things[id]

	if !ok || o.IsDead {
		r := Reply4XX{
			Res:     &res,
			Req:     req,
			Message: fmt.Sprintf("Thing %v not found", id),
		}
		Write4XXReply(404, &r)
		return
	}

	data := PrepResponseData(req, o, "Found thing")
	WriteReply(&res, req, data)
}
