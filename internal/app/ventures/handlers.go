package ventures

import (
	"fmt"
	"log"
	"net/http"

	h "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/uhttp"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

var ventures = v.NewVentureStore()

// VenturesHandler handles requests to do with collections of, or individual,
// Ventures.
func VenturesHandler(res http.ResponseWriter, req *http.Request) {
	h.LogRequest(req)
	h.AppendCORSHeaders(&res, "GET, POST, PUT, DELETE, OPTIONS")

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
		_DELETE_Venture_NEW(&res, req)
	case req.Method == "OPTIONS":
		res.WriteHeader(http.StatusOK)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// _GET_AllVentures handles client requests for all living Ventures.
func _GET_AllVentures(res *http.ResponseWriter, req *http.Request) {
	vens := ventures.GetAll()
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
	new, ok := decodeNewVenture(res, req)
	if !ok {
		return
	}

	new.Clean()
	ok = validateNewVenture(new, res, req)
	if !ok {
		return
	}

	ven := ventures.Add(new)
	m := fmt.Sprintf("New Venture with ID '%s' created", ven.ID)
	log.Println(m)
	writeSuccessReply(res, req, http.StatusCreated, ven, m)
}

// _PUT_UpdatedVenture handles client requests for updating Ventures.
func _PUT_UpdatedVenture(res *http.ResponseWriter, req *http.Request) {
	mv, ok := decodeModVentures(res, req)
	if !ok {
		return
	}

	mv.Clean()
	ok = validateModVentures(mv, res, req)
	if !ok {
		return
	}

	vens := ventures.Update(mv)
	ids := ventureIDsToCSV(vens)

	m := fmt.Sprintf("Updated Ventures with the following IDs '%s'", ids)
	log.Println(m)
	writeSuccessReply(res, req, http.StatusOK, vens, m)
}

// _DELETE_Venture handles client requests for deleting a specific Venture.
func _DELETE_Venture_NEW(res *http.ResponseWriter, req *http.Request) {

	ids := req.FormValue("ids")
	idSlice, ok := ventureIdCsvToSlice(ids, res, req)

	if !ok {
		return
	}

	vens := ventures.Delete_NEW(idSlice)
	ids = ventureIDsToCSV(vens)

	m := fmt.Sprintf("Deleted Ventures with the following IDs '%s'", ids)
	log.Println(m)
	writeSuccessReply(res, req, http.StatusOK, vens, m)
}
