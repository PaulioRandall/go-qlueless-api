package things

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// ThingsHandler handles requests to do with collections of Things
func ThingsHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	switch req.Method {
	case "GET":
		GetAllThings(&res, req)
	case "POST":
		StoreNewThing(&res, req)
	case "OPTIONS":
		WriteEmptyReply(&res)
	default:
		MethodNotAllowed(&res, req)
	}
}

// GetAllThings generates responses for requests for all Things
func GetAllThings(res *http.ResponseWriter, req *http.Request) {

	t := make([]Thing, 0)
	for _, v := range Things.GetAll() {
		if !v.IsDead {
			t = append(t, v)
		}
	}

	m := fmt.Sprintf("Found %d Things", len(t))
	data := PrepResponseData(req, t, m)
	WriteReply(res, req, data)
}

// StoreThingHandler handles requests to create new things
func StoreNewThing(res *http.ResponseWriter, req *http.Request) {

	var t Thing
	d := json.NewDecoder(req.Body)

	err := d.Decode(&t)
	if err != nil {
		r := Reply4XX{
			Res:     res,
			Req:     req,
			Message: "Unable to decode request body into a Thing",
		}
		Write4XXReply(400, &r)
		return
	}

	CleanThing(&t)
	e := ValidateThing(&t, true)
	if e != nil {
		r := Reply4XX{
			Res:     res,
			Req:     req,
			Message: strings.Join(e, " "),
		}
		Write4XXReply(400, &r)
		return
	}

	t = Things.Add(t)
	m := fmt.Sprintf("New Thing with ID %d created", t.ID)
	data := PrepResponseData(req, t, m)
	WriteReply(res, req, data)
}