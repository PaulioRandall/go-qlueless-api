package api

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type HTTPMethod string

func (m *HTTPMethod) String() string {
	return (string)(*m)
}

const (
	GET     HTTPMethod = "GET"
	POST    HTTPMethod = "POST"
	PUT     HTTPMethod = "PUT"
	DELETE  HTTPMethod = "DELETE"
	HEAD    HTTPMethod = "HEAD"
	OPTIONS HTTPMethod = "OPTIONS"
)

type APICall struct {
	URL    string
	Method HTTPMethod
}

func adminPrint(m string) {
	log.Println("[API-TEST] " + m)
}

func attemptRecover(exitCode *int) {
	if r := recover(); r != nil {
		adminPrint("Recovered from panic, exiting a little more gracefully")
		*exitCode = 1
	}
}

func startServer() *exec.Cmd {
	cmd := &exec.Cmd{
		Path:   "./go-qlueless-assembly-api",
		Dir:    "../../bin",
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	adminPrint("Starting server: " + cmd.Path)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	adminPrint("Pause to let server start")
	time.Sleep(500 * time.Millisecond)
	return cmd
}

func stopServer(cmd *exec.Cmd) {
	adminPrint("Killing server: " + cmd.Path)
	err := cmd.Process.Kill()
	if err != nil {
		log.Fatal("StopServer(): ", err)
	}
}

func newRequest(method string, url string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Panic("newRequest(): ", err)
	}
	return req
}

func invokeRequest(req *http.Request) *http.Response {
	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	res, err := client.Do(req)
	if err != nil {
		log.Panic("InvokeRequest(): ", err)
	}
	return res
}

func (c *APICall) fire() *http.Response {
	req := newRequest(c.Method.String(), c.URL)
	res := invokeRequest(req)
	return res
}
