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
	_writeSuccessReply(res, req, http.StatusOK, vens, m)
}

// _GET_Venture handles client requests for a specific Venture.
func _GET_Venture(id string, res *http.ResponseWriter, req *http.Request) {
	ven, ok := _findVenture(id, res, req)
	if !ok {
		return
	}

	m := fmt.Sprintf("Found Venture '%s'", id)
	_writeSuccessReply(res, req, http.StatusOK, ven, m)
}

// _POST_NewVenture handles client requests for creating new Ventures.
func _POST_NewVenture(res *http.ResponseWriter, req *http.Request) {
	ven, ok := _decodeVenture(res, req)
	if !ok {
		return
	}

	ven.Clean()
	ven.IsAlive = true
	ok = _validateNewVenture(ven, res, req)
	if !ok {
		return
	}

	ven = ventures.Add(ven)
	m := fmt.Sprintf("New Venture with ID '%s' created", ven.ID)
	log.Println(m)
	_writeSuccessReply(res, req, http.StatusCreated, ven, m)
}

// _PUT_UpdatedVenture handles client requests for updating Ventures.
func _PUT_UpdatedVenture(res *http.ResponseWriter, req *http.Request) {
	vu, ok := _decodeVentureUpdate(res, req)
	if !ok {
		return
	}

	vu.Clean()
	ok = _validateVentureUpdate(vu, res, req)
	if !ok {
		return
	}

	ven, ok := _updateVenture(vu, res, req)
	if !ok {
		return
	}

	m := fmt.Sprintf("Venture with ID '%s' updated", ven.ID)
	log.Println(m)
	_writeSuccessReply(res, req, http.StatusOK, ven, m)
}

// _DELETE_Venture handles client requests for deleting a specific Venture.
func _DELETE_Venture(id string, res *http.ResponseWriter, req *http.Request) {
	ven, ok := _deleteVenture(id, res, req)
	if !ok {
		return
	}

	m := fmt.Sprintf("Venture with ID '%s' deleted", ven.ID)
	log.Println(m)
	_writeSuccessReply(res, req, http.StatusOK, ven, m)
}

// _writeSuccessReply writes a success response.
func _writeSuccessReply(res *http.ResponseWriter, req *http.Request, code int, data interface{}, msg string) {
	AppendJSONHeader(res, "")
	(*res).WriteHeader(code)
	reply := PrepResponseData(req, data, msg)
	json.NewEncoder(*res).Encode(reply)
}

// _findVenture finds the Venture with the specified ID.
func _findVenture(id string, res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	ven, ok := ventures.Get(id)
	if !ok || !ven.IsAlive {
		WriteBadRequest(res, req, fmt.Sprintf("Thing '%s' not found", id))
		return v.Venture{}, false
	}
	return ven, true
}

// _decodeVenture decodes a Venture from a Request.Body.
func _decodeVenture(res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	ven, err := v.DecodeVenture(req.Body)
	if err != nil {
		WriteBadRequest(res, req, "Unable to decode request body into a Venture")
		return v.Venture{}, false
	}
	return ven, true
}

// _validateNewVenture validates a new Venture that has yet to be assigned an ID.
func _validateNewVenture(ven v.Venture, res *http.ResponseWriter, req *http.Request) bool {
	errMsgs := ven.Validate(true)
	if len(errMsgs) != 0 {
		WriteBadRequest(res, req, strings.Join(errMsgs, " "))
		return false
	}
	return true
}

// _decodeVentureUpdate decodes an update to a Venture from a Request.Body.
func _decodeVentureUpdate(res *http.ResponseWriter, req *http.Request) (v.VentureUpdate, bool) {
	vu, err := v.DecodeVentureUpdate(req.Body)
	if err != nil {
		WriteBadRequest(res, req,
			"Unable to decode request body into a Venture update")
		return v.VentureUpdate{}, false
	}
	return vu, true
}

// _validateVentureUpdate validates a Venture update.
func _validateVentureUpdate(vu v.VentureUpdate, res *http.ResponseWriter, req *http.Request) bool {
	errMsgs := vu.Validate()
	if len(errMsgs) != 0 {
		WriteBadRequest(res, req, strings.Join(errMsgs, " "))
		return false
	}
	return true
}

// _updateVenture updates a Venture in the data store.
func _updateVenture(vu v.VentureUpdate, res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	ven, ok := ventures.Update(vu)
	if !ok {
		WriteBadRequest(res, req,
			fmt.Sprintf("Venture with ID '%s' could not be found", vu.Values.ID))
		return v.Venture{}, false
	}
	return ven, true
}

// _deleteVenture deletes a Venture from the data store.
func _deleteVenture(id string, res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	ven, ok := ventures.Delete(id)
	if !ok {
		WriteBadRequest(res, req, fmt.Sprintf("Thing '%s' not found", id))
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
