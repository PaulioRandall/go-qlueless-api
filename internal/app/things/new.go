package things

import (
	"encoding/json"
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// NewThingHandler handles requests to create new things
func NewThingHandler(res http.ResponseWriter, req *http.Request) {
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
			Message: "Unable to decode create thing request body",
		}
		Write4XXReply(400, &r)
		return
	}

	o := mapToThing(m)
	o.ID, err = addThing(o)
	if err != nil {
		Write500Reply(&res, req)
		return
	}

	data := PrepResponseData(req, o, "New thing created")
	WriteReply(&res, req, data)
}
