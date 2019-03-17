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

	if !ok {
		r := Reply4XX{
			Res:     &res,
			Req:     req,
			Message: fmt.Sprintf("Order %v not found", id),
		}
		Http_4XX(404, &r)
		return
	}

	data := prepOrderData(req, o)
	WriteReply(&res, req, data)
}

// prepOrderData prepares the data by wrapping it up if the client has
// requested
func prepOrderData(req *http.Request, data interface{}) interface{} {
	if WrapUpReply(req) {
		return ReplyWrapped{
			Message: "Found order",
			Self:    req.URL.String(),
			Data:    data,
		}
	} else {
		return data
	}
}
