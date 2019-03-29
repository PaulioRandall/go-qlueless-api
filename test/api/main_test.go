package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	p "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	CORS_METHODS_PATTERN = "^((\\s*[A-Z]*\\s*,)+)*(\\s*[A-Z]*\\s*)$" // Example: 'GET,  POST   ,OPTIONS'
)

var ALL_STD_HTTP_METHODS = []string{
	"GET",
	"POST",
	"PUT",
	"DELETE",
	"HEAD",
	"OPTIONS",
	"CONNECT",
	"TRACE",
	"PATCH",
	"CUSTOM",
}

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

// assertWrappedErrorBody assert that a response body is a generic error
func assertWrappedErrorBody(t *testing.T, res *http.Response) {
	var reply p.WrappedReply
	err := json.NewDecoder(res.Body).Decode(&reply)
	require.Nil(t, err)
	AssertGenericError(t, reply)
}

// assertNotAllowedMethods asserts that the supplied methods are not allowed
// for provided endpoint
func assertNotAllowedMethods(t *testing.T, url string, allowedMethods []string) {

	isAllowed := func(m string) bool {
		for _, allowed := range allowedMethods {
			if m == allowed {
				return true
			}
		}
		return false
	}

	for _, m := range ALL_STD_HTTP_METHODS {
		if isAllowed(m) {
			continue
		}

		req := APICall{
			URL:    url,
			Method: m,
		}
		res := req.fire()
		defer res.Body.Close()
		defer PrintResponse(t, res.Body)

		ok := assert.Equal(t, 405, res.StatusCode, "Expected method not allowed: ("+m+")")
		if !ok {
			continue
		}
		assertDefaultHeaders(t, res, "application/json", allowedMethods)
		assertWrappedErrorBody(t, res)
	}
}

// TODO: Write 404 tests for non-existent resources
