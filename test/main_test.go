package test

import (
	"fmt"
	"os"
	"testing"

	a "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// _exit is file private function that accepts a referenced integer instead of
// an actual one
func _exit(exitCode *int) {
	os.Exit(*exitCode)
}

// TestMain is the entry point for the 'go test'
func TestMain(m *testing.M) {
	var exitCode int = 1
	//startServer()

	defer _exit(&exitCode)
	defer attemptRecover(&exitCode)
	//defer stopServer()

	exitCode = m.Run()
	adminPrint(fmt.Sprintf("Test exit code: %d", exitCode))
}

// verifyNotAllowedMethods asserts that the supplied methods are not allowed
// for provided endpoint
func verifyNotAllowedMethods(t *testing.T, url string, allowedMethods []string) {

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
		defer a.PrintResponse(t, res.Body)

		ok := assert.Equal(t, 405, res.StatusCode, "Expected method not allowed: ("+m+")")
		if !ok {
			continue
		}
		assertNoContentHeaders(t, res, allowedMethods)
		assertEmptyBody(t, res.Body)
	}
}

// verifyDefaultHeaders asserts that the default headers were provided
func verifyDefaultHeaders(t *testing.T, c APICall, expCode int, allowedMethods []string) {
	res := c.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, expCode, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", allowedMethods)
}
