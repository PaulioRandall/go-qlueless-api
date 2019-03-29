package ventures

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

var ventures = v.NewVentureStore()

// VenturesHandler handles requests to do with collections of, or individual,
// Ventures.
func VenturesHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)
	AppendCORSHeaders(&res, "GET, POST, PUT, DELETE, HEAD, OPTIONS")

	id := req.FormValue("id")
	switch {
	case req.Method == "GET" && id == "":
		get_AllVentures(&res, req)
	case req.Method == "GET":
		get_OneVenture(id, &res, req)
	case req.Method == "POST":
		post_NewVenture(&res, req)
	//post_NewThing(&res, req)
	//case req.Method == "PUT":
	//put_OneThing(&res, req)
	case req.Method == "HEAD":
		fallthrough
	case req.Method == "OPTIONS":
		AppendJSONHeader(&res, "")
		res.WriteHeader(http.StatusOK)
	default:
		MethodNotAllowed(&res, req)
	}
}

// get_AllVentures handles client requests for all living Ventures.
func get_AllVentures(res *http.ResponseWriter, req *http.Request) {
	vens := ventures.GetAllAlive()
	m := fmt.Sprintf("Found %d Ventures", len(vens))
	data := PrepResponseData(req, vens, m)

	AppendJSONHeader(res, "")
	(*res).WriteHeader(http.StatusOK)
	json.NewEncoder(*res).Encode(data)
}

// get_OneVenture handles client requests for a specific Venture.
func get_OneVenture(id string, res *http.ResponseWriter, req *http.Request) {
	ven, ok := findVenture(id, res, req)
	if !ok {
		return
	}

	m := fmt.Sprintf("Found Venture '%s'", id)
	data := PrepResponseData(req, ven, m)

	AppendJSONHeader(res, "")
	(*res).WriteHeader(http.StatusOK)
	json.NewEncoder(*res).Encode(data)
}

// post_NewVenture handles client requests for creating new Ventures.
func post_NewVenture(res *http.ResponseWriter, req *http.Request) {
	ven, ok := decodeVenture(res, req)
	if !ok {
		return
	}

	ven.Clean()
	ven.IsAlive = true
	ven, ok = validateNewVenture(ven, res, req)
	if !ok {
		return
	}

	ven = ventures.Add(ven)
	m := fmt.Sprintf("New Venture with ID '%s' created", ven.ID)
	log.Println(m)
	data := PrepResponseData(req, ven, m)

	AppendJSONHeader(res, "")
	(*res).WriteHeader(http.StatusCreated)
	json.NewEncoder(*res).Encode(data)
}

// findVenture finds the Venture with the specified ID.
func findVenture(id string, res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	ven, ok := ventures.Get(id)
	if !ok || !ven.IsAlive {
		r := WrappedReply{
			Message: fmt.Sprintf("Thing '%s' not found", id),
		}
		Write4XXReply(res, req, 404, r)
		return v.Venture{}, false
	}
	return ven, true
}

// decodeVenture decodes a Venture from a Request.Body.
func decodeVenture(res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	var ven v.Venture
	d := json.NewDecoder(req.Body)
	err := d.Decode(&ven)
	if err != nil {
		r := WrappedReply{
			Message: "Unable to decode request body into a Venture",
		}
		Write4XXReply(res, req, 400, r)
		return v.Venture{}, false
	}
	return ven, true
}

// validateNewVenture validates a new Venture that has yet to be assigned an ID.
func validateNewVenture(ven v.Venture, res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	errMsgs := ven.Validate(true)
	if len(errMsgs) != 0 {
		r := WrappedReply{
			Message: strings.Join(errMsgs, " "),
		}
		Write4XXReply(res, req, 400, r)
		return v.Venture{}, false
	}
	return ven, true
}

// InjectDummyVentures injects dummy Ventures so the API testing can performed.
// This function is expected to be removed once a database and formal test data
// has been crafted.
func InjectDummyVentures() {
	ventures.Add(v.Venture{
		Description: "White wizard",
		State:       "Not started",
		Extra:       "colour: white; power: 9000",
		IsAlive:     true,
	})
	ventures.Add(v.Venture{
		Description: "Green lizard",
		State:       "In progress",
		OrderIDs:    "4,5,6,7,8",
		IsAlive:     true,
	})
	ventures.Add(v.Venture{
		Description: "Pink gizzard",
		State:       "Finished",
		OrderIDs:    "1,2,3",
		IsAlive:     false,
	})
}
