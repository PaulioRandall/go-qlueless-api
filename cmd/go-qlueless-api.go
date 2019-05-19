//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"log"
	"net/http"

	changelog "github.com/PaulioRandall/go-qlueless-api/cmd/changelog"
	home "github.com/PaulioRandall/go-qlueless-api/cmd/home"
	openapi "github.com/PaulioRandall/go-qlueless-api/cmd/openapi"
	ventures "github.com/PaulioRandall/go-qlueless-api/cmd/ventures"
	qserver "github.com/PaulioRandall/go-qlueless-api/shared/qserver"
)

// Main is the entry point for the web server
func main() {
	log.Println("[Go Qlueless Assembly API]: Starting application")

	_preload()
	qserver.Sev.Init()
	defer qserver.Sev.Close()
	_routes()

	log.Println("[Go Qlueless Assembly API]: Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// _preload performs any loading of configurations or preloading of static
// values
func _preload() {
	changelog.LoadChangelog()
	openapi.LoadSpec()
}

// _routes attaches the service routes to the servers router
func _routes() {
	http.HandleFunc("/", home.HomeHandler)
	http.HandleFunc("/changelog", changelog.ChangelogHandler)
	http.HandleFunc("/openapi", openapi.OpenAPIHandler)
	http.HandleFunc("/ventures", ventures.Handler)
}
