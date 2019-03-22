package things

import (
	"fmt"
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// ThingsHandler handles requests to do with collections of Things
func ThingsHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	switch req.Method {
	case "GET":
		get_AllThings(&res, req)
	case "POST":
		post_NewThing(&res, req)
	case "OPTIONS":
		WriteEmptyReply(&res)
	default:
		MethodNotAllowed(&res, req)
	}
}

// get_AllThings generates responses for requests for all Things
func get_AllThings(res *http.ResponseWriter, req *http.Request) {
	t := Things.GetAllAlive()
	m := fmt.Sprintf("Found %d Things", len(t))
	data := PrepResponseData(req, t, m)
	WriteReply(res, req, data)
}

// post_NewThing handles requests to create new things
func post_NewThing(res *http.ResponseWriter, req *http.Request) {

	t, ok := decodeThing(res, req)
	if !ok {
		return
	}

	t, ok = checkThing(t, res, req)
	if !ok {
		return
	}

	t = Things.Add(t)
	m := fmt.Sprintf("New Thing with ID %d created", t.ID)
	data := PrepResponseData(req, t, m)
	WriteReply(res, req, data)
}
