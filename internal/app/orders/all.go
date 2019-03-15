package orders

import (
	"net/http"

	shr "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// AllOrdersHandler handles requests for all orders currently within the
// service
func AllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	shr.LogRequest(r)

	orders := LoadOrders()
	if orders == nil {
		shr.Http_500(w)
		return
	}

	shr.WriteJsonReply("Found all orders", orders, w, r)
}
