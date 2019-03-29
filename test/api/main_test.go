package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	CORS_METHODS_PATTERN = "^((\\s*[A-Z]*\\s*,)+)*(\\s*[A-Z]*\\s*)$" // Example: 'GET,  POST   ,OPTIONS'
)

func exit(exitCode *int) {
	os.Exit(*exitCode)
}

func TestMain(m *testing.M) {
	var exitCode int = 1
	cmd := startServer()

	defer exit(&exitCode)
	defer attemptRecover(&exitCode)
	defer stopServer(cmd)

	exitCode = m.Run()
	adminPrint(fmt.Sprintf("Test exit code: %d", exitCode))
}

// assertDefaultHeaders asserts that the services default headers were applied
func assertDefaultHeaders(t *testing.T, res *http.Response, contentType string, allowedMethods []string) {
	AssertHeadersEquals(t, res.Header, map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "*",
		"Content-Type":                 contentType + "; charset=utf-8",
	})
	AssertHeadersContains(t, res.Header, map[string][]string{
		"Access-Control-Allow-Methods": allowedMethods,
	})
	AssertHeadersMatches(t, res.Header, map[string]string{
		"Access-Control-Allow-Methods": CORS_METHODS_PATTERN,
	})
}

// assertEmptyBody asserts that a response body is empty
func assertEmptyBody(t *testing.T, res *http.Response) {
	body, err := ioutil.ReadAll(res.Body)
	require.Nil(t, err)
	assert.Empty(t, body)
}

// assertNotEmptyBody asserts that a response body is NOT empty
func assertNotEmptyBody(t *testing.T, res *http.Response) {
	body, err := ioutil.ReadAll(res.Body)
	require.Nil(t, err)
	assert.NotEmpty(t, body)
}

// TODO: Write 404 tests for non-existent resources
