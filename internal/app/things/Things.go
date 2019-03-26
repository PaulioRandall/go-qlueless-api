package things

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// ThingsHandler handles requests to do with collections of Things
func ThingsHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	id := req.FormValue("id")
	switch {
	case req.Method == "GET" && id == "":
		get_AllThings(&res, req)
	case req.Method == "GET":
		get_OneThing(id, &res, req)
	case req.Method == "POST":
		post_NewThing(&res, req)
	case req.Method == "PUT":
		put_OneThing(&res, req)
	case req.Method == "OPTIONS":
		WriteEmptyJSONReply(&res, "")
	default:
		MethodNotAllowed(&res, req)
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

	m := fmt.Sprintf("Found Thing with ID '%s'", id)
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
	m := fmt.Sprintf("New Thing with ID '%s' created", t.ID)
	log.Println(m)
	data := PrepResponseData(req, t, m)
	WriteJSONReply(res, req, data, "")
}

// put_OneThing handles requests to create new things
func put_OneThing(res *http.ResponseWriter, req *http.Request) {
	t, ok := decodeThing(res, req)
	if !ok {
		return
	}

	_, ok = findThing(t.ID, res, req)
	if !ok {
		return
	}

	t, ok = checkThing(t, res, req)
	if !ok {
		return
	}

	err := Things.Update(t)
	if LogIfErr(err) {
		Write500Reply(res, req)
		return
	}

	m := fmt.Sprintf("Thing with ID '%s' updated", t.ID)
	log.Println(m)
	data := PrepResponseData(req, t, m)
	WriteJSONReply(res, req, data, "")
}
