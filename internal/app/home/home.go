package home

import (
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// HomeHandler handles requests to the root path and requests to nothing (404s)
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)
	notFound(&res, req)
}

// notFound handles requests nothing (404s)
func notFound(res *http.ResponseWriter, req *http.Request) {
	r := ReplyMeta{
		Message: "Resource not found",
	}

	Write4XXReply(res, req, 404, r)
}
