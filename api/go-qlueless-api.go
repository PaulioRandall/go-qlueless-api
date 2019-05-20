//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"log"
	"net/http"

	changelog "github.com/PaulioRandall/go-qlueless-api/api/changelog"
	home "github.com/PaulioRandall/go-qlueless-api/api/home"
	openapi "github.com/PaulioRandall/go-qlueless-api/api/openapi"
	std "github.com/PaulioRandall/go-qlueless-api/api/std"
	ventures "github.com/PaulioRandall/go-qlueless-api/api/ventures"
)

// Main is the entry point for the web server
func main() {
	log.Println("[Go Qlueless Assembly API]: Starting application")

	std.Sev.Init()
	defer std.Sev.Close()
	_routes()

	log.Println("[Go Qlueless Assembly API]: Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// _routes attaches the service routes to the servers router
func _routes() {
	http.HandleFunc("/", home.HomeHandler)
	http.HandleFunc("/changelog", changelog.ChangelogHandler)
	http.HandleFunc("/openapi", openapi.OpenAPIHandler)
	http.HandleFunc("/ventures", ventures.Handler)
}
