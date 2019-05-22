package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/PaulioRandall/go-qlueless-api/api/changelog"
	"github.com/PaulioRandall/go-qlueless-api/api/database"
	"github.com/PaulioRandall/go-qlueless-api/api/home"
	"github.com/PaulioRandall/go-qlueless-api/api/openapi"
	"github.com/PaulioRandall/go-qlueless-api/api/ventures"
)

var server *http.Server = nil
var onShutdownHandlerComplete chan bool = make(chan bool)

// init attaches the endpoints to the default server.
func init() {
	http.HandleFunc("/", home.HomeHandler)
	http.HandleFunc("/changelog", changelog.ChangelogHandler)
	http.HandleFunc("/openapi", openapi.OpenAPIHandler)
	http.HandleFunc("/ventures", ventures.Handler)
}

// StartUp initialises and starts the HTTP server, blocking to handle requests
// if 'async' is true else the function will return once listening has started.
func StartUp(async bool) {
	initServer()
	registerShutdownHandler()

	log.Println("[Go Qlueless API]: Starting server")
	database.Open()
	ln := listen()

	if async {
		go serve(ln)
	} else {
		serve(ln)
	}
}

// Shutdown attempts to shutdown the server gracefully.
func Shutdown() {
	err := server.Shutdown(context.Background())
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}

	ok := <-onShutdownHandlerComplete
	if !ok {
		panic("Something went wrong while attempting to shutdown")
	}
}

// initServer creates and initialises the HTTP server.
func initServer() {
	log.Println("[Go Qlueless API]: Initialising server")

	if server != nil {
		panic("Server already in use")
	}

	server = &http.Server{Addr: ":8080"}
}

// registerShutdownHandler registers a shutdown handler to the current server.
func registerShutdownHandler() {
	server.RegisterOnShutdown(func() {
		assumeTheWorst := false
		var ok *bool = &assumeTheWorst

		defer func() {
			onShutdownHandlerComplete <- *ok
		}()

		log.Println("[Go Qlueless API]: Stopping server")
		database.Close()
		server = nil
		*ok = true
	})
}

// listen wraps http.Server.Listen() with any returned errors causing a panic.
func listen() net.Listener {
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		panic(err)
	}
	return ln
}

// serve wraps http.Server.Serve() so that it can be run as a separate Go
// routine if needed.
func serve(ln net.Listener) {
	err := server.Serve(ln.(*net.TCPListener))
	if err != http.ErrServerClosed {
		panic(err)
	}
}
