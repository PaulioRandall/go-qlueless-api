package ventures

import (
	"fmt"
	"log"
	"net/http"

	q "github.com/PaulioRandall/go-qlueless-assembly-api/internal/qserver"
	h "github.com/PaulioRandall/go-qlueless-assembly-api/internal/uhttp"
	u "github.com/PaulioRandall/go-qlueless-assembly-api/internal/utils"
)

// Handler handles requests to do with collections of, or individual, Ventures.
func Handler(res http.ResponseWriter, req *http.Request) {
	h.LogRequest(req)
	h.AppendCORSHeaders(&res, "GET, POST, PUT, DELETE, OPTIONS")

	switch {
	case req.Method == "GET":
		_GET(&res, req)
	case req.Method == "POST":
		_POST(&res, req)
	case req.Method == "PUT":
		_PUT(&res, req)
	case req.Method == "OPTIONS":
		res.WriteHeader(http.StatusOK)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// _GET handles client requests for any amount of living Ventures.
func _GET(res *http.ResponseWriter, req *http.Request) {

	ids := req.FormValue("ids")
	ids = u.StripWhitespace(ids)
	var vens []Venture

	switch {
	case ids == "":
		var err error
		vens, err = QueryAll(q.Sev.DB)
		if err != nil {
			h.WriteServerError(res, req)
			return
		}
	default:
		var ok bool
		vens, ok = find(ids, res, req)
		if !ok {
			return
		}
	}

	m := fmt.Sprintf("Found %d Ventures", len(vens))
	h.WriteSuccessReply(res, req, http.StatusOK, vens, m)
}

// _POST handles client requests for creating new Ventures.
func _POST(res *http.ResponseWriter, req *http.Request) {
	new, ok := decodeNew(res, req)
	if !ok {
		return
	}

	new.Clean()
	ok = validateNew(&new, res, req)
	if !ok {
		return
	}

	ven, ok := insertNew(&new, res, req)
	if !ok {
		return
	}

	m := fmt.Sprintf("New Venture with ID '%s' created", ven.ID)
	log.Println(m)
	h.WriteSuccessReply(res, req, http.StatusCreated, ven, m)
}

// _PUT handles client requests for updating Ventures.
func _PUT(res *http.ResponseWriter, req *http.Request) {
	mv, ok := decodeMod(res, req)
	if !ok {
		return
	}

	mv.Clean()
	ok = validateMod(mv, res, req)
	if !ok {
		return
	}

	vens, ok := pushMod(mv, res, req)
	if !ok {
		return
	}

	ids := idsToCSV(vens)
	m := fmt.Sprintf("Updated Ventures with the following IDs '%s'", ids)
	log.Println(m)
	h.WriteSuccessReply(res, req, http.StatusOK, vens, m)
}
