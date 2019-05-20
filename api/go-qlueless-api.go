//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"net/http"

	"github.com/PaulioRandall/go-qlueless-api/api/changelog"
	"github.com/PaulioRandall/go-qlueless-api/api/home"
	"github.com/PaulioRandall/go-qlueless-api/api/openapi"
	"github.com/PaulioRandall/go-qlueless-api/api/std"
	"github.com/PaulioRandall/go-qlueless-api/api/ventures"
)

// Main is the primary entry point for the HTTP server.
func main() {
	Start()
}

// Start initialises and starts the HTTP server.
func Start() {
	http.HandleFunc("/", home.HomeHandler)
	http.HandleFunc("/changelog", changelog.ChangelogHandler)
	http.HandleFunc("/openapi", openapi.OpenAPIHandler)
	http.HandleFunc("/ventures", ventures.Handler)

	defer std.Sev.Close()
	std.Sev.Start()
}
