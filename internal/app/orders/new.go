package orders

import (
	"encoding/json"
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// NewOrderHandler handles requests to create new orders
func NewOrderHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	if req.Method == "OPTIONS" {
		WriteEmptyReply(&res)
		return
	}

	d := json.NewDecoder(req.Body)
	var m map[string]interface{}

	err := d.Decode(&m)
	if err != nil {
		r := Reply4XX{
			Res:     &res,
			Req:     req,
			Message: "Unable to decode create order request body",
		}
		Write4XXReply(400, &r)
		return
	}

	o := mapToOrder(m)
	o.WorkItemID, err = addOrder(o)
	if err != nil {
		Write500Reply(&res, req)
		return
	}

	data := prepNewOrderData(req, o)
	WriteReply(&res, req, data)
}

// prepNewOrderData prepares the data by wrapping it up if the client has
// requested
func prepNewOrderData(req *http.Request, data interface{}) interface{} {
	if WrapUpReply(req) {
		return ReplyWrapped{
			Message: "New order created",
			Self:    req.URL.String(),
			Data:    data,
		}
	} else {
		return data
	}
}
