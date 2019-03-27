//usr/bin/env go run "$0" "$@"; exit "$?"

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

type Thing struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	ChildIDs    string `json:"child_ids,omitempty"`
	ParentIDs   string `json:"parent_ids,omitempty"`
	State       string `json:"state"`
	IsDead      bool   `json:"-"`
	Additional  string `json:"additional,omitempty"`
}

type ApiCall struct {
	URL    string
	Method HTTPMethod
}

func adminPrint(m string) {
	log.Println("[API-TEST] " + m)
}

func StartServer() *exec.Cmd {
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
	time.Sleep(1 * time.Second)
	return cmd
}

func StopServer(cmd *exec.Cmd) {
	adminPrint("Killing server: " + cmd.Path)
	err := cmd.Process.Kill()
	if err != nil {
		log.Fatal("StopServer(): ", err)
	}
}

func TestMain(m *testing.M) {
	cmd := StartServer()
	exitCode := 1

	defer func() {
		if r := recover(); r != nil {
			adminPrint("Recovered from panic, exiting gracefully")
			exitCode = 1
		}
		StopServer(cmd)
		adminPrint(fmt.Sprintf("Exit code: %d", exitCode))
		os.Exit(exitCode)
	}()

	exitCode = m.Run()
}

func NewRequest(call ApiCall) *http.Request {
	req, err := http.NewRequest(call.Method.String(), call.URL, nil)
	if err != nil {
		log.Panic("NewRequest(): ", err)
	}
	return req
}

func InvokeRequest(req *http.Request) *http.Response {
	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	res, err := client.Do(req)
	if err != nil {
		log.Panic("InvokeRequest(): ", err)
	}
	return res
}

func TestAbc(t *testing.T) {
	req := NewRequest(ApiCall{
		URL:    "http://localhost:8080/things?id=1",
		Method: GET,
	})

	res := InvokeRequest(req)
	defer res.Body.Close()

	var thing Thing
	err := json.NewDecoder(res.Body).Decode(&thing)
	require.Nil(t, err)

	assert.Equal(t, thing.ID, "1")
	assert.NotEmpty(t, thing.Description)
	assert.NotEmpty(t, thing.State)
}
