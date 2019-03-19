package things

import (
	"encoding/json"
	"log"
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

	var t Thing
	d := json.NewDecoder(req.Body)
	err := d.Decode(&t)
	if err != nil {
		log.Println(err)
		r := Reply4XX{
			Res:     &res,
			Req:     req,
			Message: "Unable to decode request body into a Thing",
		}
		Write4XXReply(400, &r)
		return
	}

	result, err := addThing(t)
	if err != nil {
		Write500Reply(&res, req)
		return
	}

	data := PrepResponseData(req, result, "New thing created")
	WriteReply(&res, req, data)
}
