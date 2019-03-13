package orders

import (
	"net/http"

	"github.com/gorilla/mux"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

// Single_order_handler handles requests for a specific orders currently
// within the service
func Single_order_handler(w http.ResponseWriter, r *http.Request) {
	shr.Log_request(r)

	orders := Load_orders()
	if orders == nil {
		shr.Http_500(&w)
		return
	}

	id := mux.Vars(r)["order_id"]
	var order *shr.WorkItem = nil

	for _, o := range orders {
		if o.Work_item_id == id {
			order = &o
		}
	}

	if order == nil {
		shr.Http_4xx(&w, 404, "Order not found")
		return
	}

	reply := shr.Reply{
		Message: "Found order",
		Data:    order,
	}

	shr.WriteJsonReply(reply, w, r)
}
