//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"github.com/PaulioRandall/go-qlueless-api/api/server"
)

// Main is the primary entry point for the HTTP server.
func main() {
	server.StartUp(false)
}
