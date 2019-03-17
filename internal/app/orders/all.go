package orders

import (
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// AllOrdersHandler handles requests for all orders currently within the
// service
func AllOrdersHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	o := make([]WorkItem, 0)
	for _, v := range orders {
		o = append(o, v)
	}

	data := prepOrdersData(req, o)
	WriteReply(&res, req, data)
}

// prepOrdersData prepares the data by wrapping it up if the client has
// requested
func prepOrdersData(req *http.Request, data interface{}) interface{} {
	if WrapUpReply(req) {
		return ReplyWrapped{
			Message: "Found all orders",
			Self:    req.URL.String(),
			Data:    data,
		}
	} else {
		return data
	}
}
