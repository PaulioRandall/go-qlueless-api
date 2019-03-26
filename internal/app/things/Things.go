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
		demuxGET(&res, req)
	case "POST":
		post_NewThing(&res, req)
	case "OPTIONS":
		WriteEmptyJSONReply(&res, "")
	default:
		MethodNotAllowed(&res, req)
	}
}

// demuxGET routes the handling of GET requests dependent on the query
// parameters passed
func demuxGET(res *http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		get_AllThings(res, req)
	} else {
		get_OneThing(id, res, req)
	}
}

// get_AllThings generates responses for requests for all Things
func get_AllThings(res *http.ResponseWriter, req *http.Request) {
	t := Things.GetAllAlive()
	m := fmt.Sprintf("Found %d Things", len(t))
	data := PrepResponseData(req, t, m)
	WriteJSONReply(res, req, data, "")
}

// get_OneThing generates responses for requests for a single Thing
func get_OneThing(id string, res *http.ResponseWriter, req *http.Request) {
	t, ok := findThing(id, res, req)
	if !ok {
		return
	}

	m := fmt.Sprintf("Found Thing with ID %s", id)
	data := PrepResponseData(req, t, m)
	WriteJSONReply(res, req, data, "")
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
	m := fmt.Sprintf("New Thing with ID %s created", t.ID)
	data := PrepResponseData(req, t, m)
	WriteJSONReply(res, req, data, "")
}
