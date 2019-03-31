package ventures

import (
	"fmt"
	"log"
	"net/http"

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
		_GET_AllVentures(&res, req)
	case req.Method == "GET":
		_GET_Venture(id, &res, req)
	case req.Method == "POST":
		_POST_NewVenture(&res, req)
	case req.Method == "PUT":
		_PUT_UpdatedVenture(&res, req)
	case req.Method == "DELETE":
		_DELETE_Venture(id, &res, req)
	case req.Method == "HEAD":
		fallthrough
	case req.Method == "OPTIONS":
		AppendJSONHeader(&res, "")
		res.WriteHeader(http.StatusOK)
	default:
		MethodNotAllowed(&res, req)
	}
}

// _GET_AllVentures handles client requests for all living Ventures.
func _GET_AllVentures(res *http.ResponseWriter, req *http.Request) {
	vens := ventures.GetAllAlive()
	m := fmt.Sprintf("Found %d Ventures", len(vens))
	writeSuccessReply(res, req, http.StatusOK, vens, m)
}

// _GET_Venture handles client requests for a specific Venture.
func _GET_Venture(id string, res *http.ResponseWriter, req *http.Request) {
	ven, ok := findVenture(id, res, req)
	if !ok {
		return
	}

	m := fmt.Sprintf("Found Venture '%s'", id)
	writeSuccessReply(res, req, http.StatusOK, ven, m)
}

// _POST_NewVenture handles client requests for creating new Ventures.
func _POST_NewVenture(res *http.ResponseWriter, req *http.Request) {
	ven, ok := decodeVenture(res, req)
	if !ok {
		return
	}

	ven.Clean()
	ven.IsAlive = true
	ok = validateNewVenture(ven, res, req)
	if !ok {
		return
	}

	ven = ventures.Add(ven)
	m := fmt.Sprintf("New Venture with ID '%s' created", ven.ID)
	log.Println(m)
	writeSuccessReply(res, req, http.StatusCreated, ven, m)
}

// _PUT_UpdatedVenture handles client requests for updating Ventures.
func _PUT_UpdatedVenture(res *http.ResponseWriter, req *http.Request) {
	vu, ok := decodeVentureUpdate(res, req)
	if !ok {
		return
	}

	vu.Clean()
	ok = validateVentureUpdate(vu, res, req)
	if !ok {
		return
	}

	ven, ok := updateVenture(vu, res, req)
	if !ok {
		return
	}

	m := fmt.Sprintf("Venture with ID '%s' updated", ven.ID)
	log.Println(m)
	writeSuccessReply(res, req, http.StatusOK, ven, m)
}

// _DELETE_Venture handles client requests for deleting a specific Venture.
func _DELETE_Venture(id string, res *http.ResponseWriter, req *http.Request) {
	ven, ok := deleteVenture(id, res, req)
	if !ok {
		return
	}

	m := fmt.Sprintf("Venture with ID '%s' deleted", ven.ID)
	log.Println(m)
	writeSuccessReply(res, req, http.StatusOK, ven, m)
}
