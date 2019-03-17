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
	r := Reply{
		Req: req,
		Res: &res,
	}

	orders := LoadOrders()
	if orders == nil {
		Http_500(&r)
		return
	}

	id := mux.Vars(req)["order_id"]
	o, ok := orders[id]

	if !ok {
		r.Message = Str(fmt.Sprintf("Order %v not found", id))
		Http_4xx(&r, 404)
		return
	}

	m := Str(fmt.Sprintf("Found order %v", id))
	WriteJsonReply(&r, m, o, nil)
}
