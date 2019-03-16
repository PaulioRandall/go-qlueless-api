package orders

import (
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// AllOrdersHandler handles requests for all orders currently within the
// service
func AllOrdersHandler(res http.ResponseWriter, req *http.Request) {
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

	WriteJsonReply(&r, Str("Found all orders"), orders, nil)
}
