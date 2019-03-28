//usr/bin/env go run "$0" "$@"; exit "$?"

package api

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
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

// Given some Ventures already exist on the server
// When all Ventures are requested
// Then ensure the response code is 200
// And the 'Content-Type' header contains 'application/json'
// And the 'Access-Control-Allow-Origin' is '*'
// And the 'Access-Control-Allow-Headers' is '*'
// And the 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, HEAD, and OPTIONS
// And the body is a JSON array of valid Ventures
func TestGET_Ventures(t *testing.T) {
	req := APICall{
		URL:    "http://localhost:8080/ventures",
		Method: GET,
	}
	res := req.fire()
	defer res.Body.Close()

	if true {
		return // REMOVE when ready to start development
	}

	require.Equal(t, 200, res.StatusCode)
	AssertHeadersEquals(t, res.Header, map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "*",
	})
	AssertHeadersContains(t, res.Header, map[string][]string{
		"Content-Type":                 []string{"application/json"},
		"Access-Control-Allow-Methods": []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
	})
	AssertHeadersMatches(t, res.Header, map[string]string{
		"Access-Control-Allow-Methods": "^((\\s*[a-zA-Z]*\\s*,)+)*(\\s*[a-zA-Z]*\\s*)$",
	})

	var ven []Venture
	err := json.NewDecoder(res.Body).Decode(&ven)
	require.Nil(t, err)
	AssertGenericVentures(t, ven)
}
