package ventures

import (
	"fmt"
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

var ventures = v.NewVentureStore()

// VenturesHandler handles requests to do with collections of, or individual,
// Ventures.
func VenturesHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	id := req.FormValue("id")
	switch {
	case req.Method == "GET" && id == "":
		get_AllVentures(&res, req)
	//case req.Method == "GET":
	//get_OneThing(id, &res, req)
	//case req.Method == "POST":
	//post_NewThing(&res, req)
	//case req.Method == "PUT":
	//put_OneThing(&res, req)
	//case req.Method == "HEAD":
	//fallthrough
	case req.Method == "OPTIONS":
		WriteEmptyJSONReply(&res, "")
	default:
		MethodNotAllowed(&res, req)
	}
}

// get_AllVentures handles client requests for all living Ventures.
func get_AllVentures(res *http.ResponseWriter, req *http.Request) {
	vens := ventures.GetAllAlive()
	m := fmt.Sprintf("Found %d Ventures", len(vens))
	data := PrepResponseData(req, vens, m)
	WriteJSONReply(res, req, data, "")
}

// InjectDummyVentures injects dummy Ventures so the API testing can performed.
// This function is expected to be removed once a database and formal test data
// has been crafted
func InjectDummyVentures() {
	//ventures
}
