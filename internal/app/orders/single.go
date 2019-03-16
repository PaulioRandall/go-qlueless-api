package orders

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	shr "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// SingleOrderHandler handles requests for a specific orders currently
// within the service
func SingleOrderHandler(w http.ResponseWriter, r *http.Request) {
	shr.LogRequest(r)

	orders := LoadOrders()
	if orders == nil {
		shr.Http_500(w)
		return
	}

	id := mux.Vars(r)["order_id"]
	var o *shr.WorkItem = shr.FindWorkItem(orders, id)

	if o == nil {
		shr.Http_4xx(w, 404, fmt.Sprintf("Order %v not found", id), "")
		return
	}

	shr.WriteJsonReply(fmt.Sprintf("Found order %v", id), o, "", w, r)
}
