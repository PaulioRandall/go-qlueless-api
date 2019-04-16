package uhttp

import (
	"encoding/json"
	"net/http"

	w "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/wrapped"
)

// WriteServerError writes the response for a generic 500 error to the client.
func WriteServerError(res *http.ResponseWriter, req *http.Request) {
	r := w.WrappedReply{
		Message: "Bummer! Something went wrong on the server.",
		Self:    (*req).URL.String(),
	}

	AppendJSONHeader(res, "")
	(*res).WriteHeader(500)
	json.NewEncoder(*res).Encode(r)
}

// WriteBadRequest writes the response for a 400 error to the client.
func WriteBadRequest(res *http.ResponseWriter, req *http.Request, m string) {
	if !CheckNotEmpty(res, req, "response message", m) {
		return
	}

	r := w.WrappedReply{
		Message: m,
		Self:    RelURL(req),
	}

	AppendJSONHeader(res, "")
	(*res).WriteHeader(http.StatusBadRequest)
	json.NewEncoder(*res).Encode(r)
}

// WriteWrappedReply writes the response for a WrappedReply to the client.
func WriteWrappedReply(res *http.ResponseWriter, req *http.Request, status int, r w.WrappedReply) {
	if !CheckNotEmpty(res, req, "response message", r.Message) {
		return
	}

	if r.Self == "" {
		r.Self = RelURL(req)
	}

	AppendJSONHeader(res, "")
	(*res).WriteHeader(status)
	json.NewEncoder(*res).Encode(r)
}

// WriteSuccessReply writes a success response.
//
// @UNTESTED
func WriteSuccessReply(res *http.ResponseWriter, req *http.Request, code int, data interface{}, msg string) {
	AppendJSONHeader(res, "")
	(*res).WriteHeader(code)
	reply := PrepResponseData(req, data, msg)
	json.NewEncoder(*res).Encode(reply)
}

// WriteJSONReply appends the required JSON headers and then writes the response
// data
func WriteJSONReply(res *http.ResponseWriter, req *http.Request, data interface{}, extensionType string) {
	AppendJSONHeader(res, extensionType)
	(*res).WriteHeader(http.StatusOK)
	json.NewEncoder(*res).Encode(data)
}

// WriteEmptyJSONReply appends the required JSON headers and sets status as OK
func WriteEmptyJSONReply(res *http.ResponseWriter, extensionType string) {
	AppendJSONHeader(res, extensionType)
	(*res).WriteHeader(http.StatusOK)
}
