package main

import (
	"net/http"
	"sync"

	gor "github.com/gorilla/mux"

	chg "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/changelog"
	oai "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/openapi"
	thg "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/things"
	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// QServer represents the... err... server
type QServer struct {
	preloadOnce sync.Once
	routeOnce   sync.Once
	router      *gor.Router
}

// preload performs any loading of configurations or preloading of static values
func (s *QServer) preload() {
	s.preloadOnce.Do(func() {
		chg.LoadChangelog()
		oai.LoadSpec()
		CreateDummyThings()
	})
}

// routes attaches the service routes to the servers router
func (s *QServer) routes() {
	s.routeOnce.Do(func() {
		s.router.HandleFunc("/changelog", chg.ChangelogHandler)
		s.router.HandleFunc("/openapi", oai.OpenAPIHandler)
		s.router.HandleFunc("/things", thg.ThingsHandler)
		s.router.HandleFunc("/things/{id}", thg.ThingHandler)

		s.router.NotFoundHandler = http.HandlerFunc(HomeHandler)
		http.Handle("/", s.router)
	})
}

// HomeHandler handles requests to the root path and requests to nothing (404s)
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	//if req.URL.Path != "/" {
	r := ReplyMeta{
		Message: "Resource not found",
	}

	Write4XXReply(&res, req, 404, r)
	//}
}
