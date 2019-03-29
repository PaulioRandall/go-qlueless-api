package api

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type APICall struct {
	URL    string
	Method string
	Body   io.Reader
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

func (c *APICall) newRequest() *http.Request {
	req, err := http.NewRequest(c.Method, c.URL, c.Body)
	if err != nil {
		log.Panic("newRequest(): ", err)
	}
	return req
}

func (c *APICall) invokeRequest(req *http.Request) *http.Response {
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
	req := c.newRequest()
	res := c.invokeRequest(req)
	return res
}
