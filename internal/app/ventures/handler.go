package ventures

import (
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	ven "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

var ventures []ven.Venture = []ven.Venture{}

// VenturesHandler handles requests to do with collections of, or individual,
// Ventures
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

func get_AllVentures(res *http.ResponseWriter, req *http.Request) {

}
