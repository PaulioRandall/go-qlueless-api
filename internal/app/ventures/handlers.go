package ventures

import (
	"fmt"
	"log"
	"net/http"

	q "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/qserver"
	h "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/uhttp"
	u "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/utils"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

// VenturesHandler handles requests to do with collections of, or individual,
// Ventures.
func VenturesHandler(res http.ResponseWriter, req *http.Request) {
	h.LogRequest(req)
	h.AppendCORSHeaders(&res, "GET, POST, PUT, OPTIONS")

	switch {
	case req.Method == "GET":
		_GET_Ventures(&res, req)
	case req.Method == "POST":
		_POST_NewVenture(&res, req)
	case req.Method == "PUT":
		_PUT_ModifiedVentures(&res, req)
	case req.Method == "OPTIONS":
		res.WriteHeader(http.StatusOK)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// _GET_Ventures handles client requests for any amount of living Ventures.
func _GET_Ventures(res *http.ResponseWriter, req *http.Request) {

	ids := req.FormValue("ids")
	ids = u.StripWhitespace(ids)
	var vens []v.Venture

	switch {
	case ids == "":
		var err error
		vens, err = v.QueryAll(q.Sev.DB)
		if err != nil {
			h.WriteServerError(res, req)
			return
		}
	default:
		var ok bool
		vens, ok = findVentures(ids, res, req)
		if !ok {
			return
		}
	}

	m := fmt.Sprintf("Found %d Ventures", len(vens))
	h.WriteSuccessReply(res, req, http.StatusOK, vens, m)
}

// _POST_NewVenture handles client requests for creating new Ventures.
func _POST_NewVenture(res *http.ResponseWriter, req *http.Request) {
	new, ok := decodeNewVenture(res, req)
	if !ok {
		return
	}

	new.Clean()
	ok = validateNewVenture(&new, res, req)
	if !ok {
		return
	}

	ven, ok := insertNewVenture(&new, res, req)
	if !ok {
		return
	}

	m := fmt.Sprintf("New Venture with ID '%s' created", ven.ID)
	log.Println(m)
	h.WriteSuccessReply(res, req, http.StatusCreated, ven, m)
}

// _PUT_ModifiedVentures handles client requests for updating Ventures.
func _PUT_ModifiedVentures(res *http.ResponseWriter, req *http.Request) {
	mv, ok := decodeModVentures(res, req)
	if !ok {
		return
	}

	mv.Clean()
	ok = validateModVentures(mv, res, req)
	if !ok {
		return
	}

	vens, ok := pushModifiedVentures(mv, res, req)
	if !ok {
		return
	}

	ids := ventureIDsToCSV(vens)
	m := fmt.Sprintf("Updated Ventures with the following IDs '%s'", ids)
	log.Println(m)
	h.WriteSuccessReply(res, req, http.StatusOK, vens, m)
}
