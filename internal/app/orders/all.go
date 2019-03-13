package orders

import (
	"net/http"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

// All_orders_handler handles requests for all orders currently within the
// service
func All_orders_handler(w http.ResponseWriter, r *http.Request) {
	shr.Log_request(r)

	orders := Load_orders()
	if orders == nil {
		shr.Http_500(&w)
		return
	}

	reply := shr.Reply{
		Message: "Found all orders",
		Data:    orders,
	}

	shr.WriteJsonReply(reply, w, r)
}
