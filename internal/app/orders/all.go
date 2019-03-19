package orders

import (
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// AllOrdersHandler handles requests for all orders currently within the
// service
func AllOrdersHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	o := make([]Thing, 0)
	for _, v := range orders {
		if !v.IsDead {
			o = append(o, v)
		}
	}

	data := PrepResponseData(req, o, "Found all orders")
	WriteReply(&res, req, data)
}
