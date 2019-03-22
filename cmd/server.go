package main

import (
	"net/http"

	gor "github.com/gorilla/mux"

	oai "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/openapi"
	thg "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/things"
	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// Server represents the... err... server
type Server struct {
	router *gor.Router
}

// preload performs any loading of configurations or preloading of static values
func (s *Server) preload() {
	oai.LoadSpec()
	CreateDummyThings()
}

// routes attaches the service routes to the servers router
func (s *Server) routes() {
	s.router.HandleFunc("/openapi", oai.OpenAPIHandler)
	s.router.HandleFunc("/things", thg.ThingsHandler)
	s.router.HandleFunc("/things/{id}", thg.ThingHandler)

	s.router.NotFoundHandler = http.HandlerFunc(HomeHandler)
	http.Handle("/", s.router)
}

// HomeHandler handles requests to the root path and requests to nothing (404s)
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	//if req.URL.Path != "/" {
	r := Reply4XX{
		Res:     &res,
		Req:     req,
		Message: "Resource not found",
	}

	Write4XXReply(404, &r)
	//}
}
