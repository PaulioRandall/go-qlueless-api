package home

import (
	"net/http"

	p "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	w "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/wrapped"
)

// HomeHandler handles requests to the root path and requests to nothing (404s)
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	p.LogRequest(req)
	notFound(&res, req)
}

// notFound handles requests nothing (404s)
func notFound(res *http.ResponseWriter, req *http.Request) {
	r := w.WrappedReply{
		Message: "Resource not found",
	}

	p.Write4XXReply(res, req, 404, r)
}
