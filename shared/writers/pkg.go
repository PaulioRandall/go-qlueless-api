package writers

import (
	"encoding/json"
	"log"
	"net/http"

	uhttp "github.com/PaulioRandall/go-cookies/uhttp"
	wrapped "github.com/PaulioRandall/go-qlueless-api/shared/wrapped"
)

// CheckNotEmpty validates that a required string is not empty writing a generic
// 500 response if it is.
func CheckNotEmpty(res *http.ResponseWriter, req *http.Request, name string, m string) bool {
	if m == "" {
		log.Println("[BUG] Error missing '" + name + "'")
		WriteServerError(res, req)
		return false
	}
	return true
}

// PrepResponseData prepares the response data read writing to the client. If
// the client has specified, the data is wrapped and meta information added else
// the input data is returned.
func PrepResponseData(req *http.Request, data interface{}, msg string) interface{} {
	if req.URL.Query()["wrap"] != nil {
		return wrapped.WrappedReply{
			Message: msg,
			Self:    uhttp.RelURL(req),
			Data:    data,
		}
	} else {
		return data
	}
}

// WriteServerError writes the response for a generic 500 error to the client.
func WriteServerError(res *http.ResponseWriter, req *http.Request) {
	r := wrapped.WrappedReply{
		Message: "Bummer! Something went wrong on the server.",
		Self:    (*req).URL.String(),
	}

	uhttp.UseUTF8Json(res, "")
	(*res).WriteHeader(500)
	json.NewEncoder(*res).Encode(r)
}

// WriteBadRequest writes the response for a 400 error to the client.
func WriteBadRequest(res *http.ResponseWriter, req *http.Request, m string) {
	if !CheckNotEmpty(res, req, "response message", m) {
		return
	}

	r := wrapped.WrappedReply{
		Message: m,
		Self:    uhttp.RelURL(req),
	}

	uhttp.UseUTF8Json(res, "")
	(*res).WriteHeader(http.StatusBadRequest)
	json.NewEncoder(*res).Encode(r)
}

// WriteWrappedReply writes the response for a WrappedReply to the client.
func WriteWrappedReply(res *http.ResponseWriter, req *http.Request, status int, r wrapped.WrappedReply) {
	if !CheckNotEmpty(res, req, "response message", r.Message) {
		return
	}

	if r.Self == "" {
		r.Self = uhttp.RelURL(req)
	}

	uhttp.UseUTF8Json(res, "")
	(*res).WriteHeader(status)
	json.NewEncoder(*res).Encode(r)
}

// WriteSuccessReply writes a success response.
func WriteSuccessReply(res *http.ResponseWriter, req *http.Request, code int, data interface{}, msg string) {
	uhttp.UseUTF8Json(res, "")
	(*res).WriteHeader(code)
	reply := PrepResponseData(req, data, msg)
	json.NewEncoder(*res).Encode(reply)
}
