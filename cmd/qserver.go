package main

import (
	"net/http"
	"sync"

	chg "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/changelog"
	hme "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/home"
	oai "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/openapi"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/ventures"
)

// QServer represents the... err... server
type QServer struct {
	preloadOnce sync.Once
	routeOnce   sync.Once
}

// preload performs any loading of configurations or preloading of static values
func (s *QServer) preload() {
	s.preloadOnce.Do(func() {
		chg.LoadChangelog()
		oai.LoadSpec()
		v.InjectDummyVentures()
	})
}

// routes attaches the service routes to the servers router
func (s *QServer) routes() {
	s.routeOnce.Do(func() {
		http.HandleFunc("/", hme.HomeHandler)
		http.HandleFunc("/changelog", chg.ChangelogHandler)
		http.HandleFunc("/openapi", oai.OpenAPIHandler)
		http.HandleFunc("/ventures", v.VenturesHandler)
	})
}
