//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"log"
	"net/http"

	c "github.com/PaulioRandall/go-qlueless-api/cmd/changelog"
	h "github.com/PaulioRandall/go-qlueless-api/cmd/home"
	o "github.com/PaulioRandall/go-qlueless-api/cmd/openapi"
	v "github.com/PaulioRandall/go-qlueless-api/cmd/ventures"
	q "github.com/PaulioRandall/go-qlueless-api/internal/qserver"
)

// Main is the entry point for the web server
func main() {
	log.Println("[Go Qlueless Assembly API]: Starting application")

	_preload()
	q.Sev.Init()
	defer q.Sev.Close()
	_routes()

	log.Println("[Go Qlueless Assembly API]: Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// _preload performs any loading of configurations or preloading of static
// values
func _preload() {
	c.LoadChangelog()
	o.LoadSpec()
}

// _routes attaches the service routes to the servers router
func _routes() {
	http.HandleFunc("/", h.HomeHandler)
	http.HandleFunc("/changelog", c.ChangelogHandler)
	http.HandleFunc("/openapi", o.OpenAPIHandler)
	http.HandleFunc("/ventures", v.Handler)
}