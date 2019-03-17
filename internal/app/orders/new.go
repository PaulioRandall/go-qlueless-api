package orders

import (
	"encoding/json"
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// NewOrderHandler handles requests to create new orders
func NewOrderHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)
	r := Reply{
		Req: req,
		Res: &res,
	}

	if req.Method == "OPTIONS" {
		AppendJSONHeaders(&r)
		res.WriteHeader(200)
		return
	}

	d := json.NewDecoder(req.Body)
	var m map[string]interface{}

	err := d.Decode(&m)
	if err != nil {
		r.Message = Str("Unable to decode create order request body")
		Http_4xx(&r, 400)
		return
	}

	LoadOrders()
	o := MapToOrder(m)
	o.WorkItemID, err = AddOrder(o)
	if err != nil {
		Http_500(&r)
		return
	}

	WriteJsonReply(&r, Str("New order created"), o, nil)
}
