package ventures

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PaulioRandall/go-cookies/cookies"
	"github.com/PaulioRandall/go-cookies/uhttp"
	"github.com/PaulioRandall/go-qlueless-api/shared/writers"
)

// cors contains the CORS headers for /Venture responses.
var cors uhttp.CorsHeaders = uhttp.CorsHeaders{
	Origin:  "*",
	Headers: "*",
	Methods: "GET, POST, PUT, DELETE, OPTIONS",
}

// Handler handles requests to do with collections of, or individual, Ventures.
func Handler(res http.ResponseWriter, req *http.Request) {
	uhttp.LogRequest(req)
	uhttp.UseCors(&res, &cors)

	switch {
	case req.Method == "GET":
		get(&res, req)
	case req.Method == "POST":
		post(&res, req)
	case req.Method == "PUT":
		put(&res, req)
	case req.Method == "OPTIONS":
		res.WriteHeader(http.StatusOK)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// get handles client requests for any amount of living Ventures.
func get(res *http.ResponseWriter, req *http.Request) {

	ids := req.FormValue("ids")
	ids = cookies.StripWhitespace(ids)
	var vens []Venture

	switch {
	case ids == "":
		var err error
		vens, err = QueryAll()
		if err != nil {
			writers.WriteServerError(res, req)
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
	writers.WriteSuccessReply(res, req, http.StatusOK, vens, m)
}

// post handles client requests for creating new Ventures.
func post(res *http.ResponseWriter, req *http.Request) {
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
	writers.WriteSuccessReply(res, req, http.StatusCreated, ven, m)
}

// put handles client requests for updating Ventures.
func put(res *http.ResponseWriter, req *http.Request) {
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
	writers.WriteSuccessReply(res, req, http.StatusOK, vens, m)
}
