package orders

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// SingleOrderHandler handles requests for a specific orders currently
// within the service
func SingleOrderHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	id := mux.Vars(req)["order_id"]
	o, ok := orders[id]

	if !ok || o.IsDead {
		r := Reply4XX{
			Res:     &res,
			Req:     req,
			Message: fmt.Sprintf("Order %v not found", id),
		}
		Write4XXReply(404, &r)
		return
	}

	data := PrepResponseData(req, o, "Found order")
	WriteReply(&res, req, data)
}
