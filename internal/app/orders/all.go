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

	o := make([]WorkItem, 0)
	for _, v := range orders {
		o = append(o, v)
	}

	WriteJsonReply(&r, Str("Found all orders"), o, nil)
}
