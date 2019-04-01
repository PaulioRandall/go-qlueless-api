package main

import (
	"net/http"
	"sync"

	c "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/changelog"
	h "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/home"
	o "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/openapi"
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
		c.LoadChangelog()
		o.LoadSpec()
		v.InjectDummyVentures()
	})
}

// routes attaches the service routes to the servers router
func (s *QServer) routes() {
	s.routeOnce.Do(func() {
		http.HandleFunc("/", h.HomeHandler)
		http.HandleFunc("/changelog", c.ChangelogHandler)
		http.HandleFunc("/openapi", o.OpenAPIHandler)
		http.HandleFunc("/ventures", v.VenturesHandler)
	})
}
