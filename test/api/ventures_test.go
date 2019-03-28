//usr/bin/env go run "$0" "$@"; exit "$?"

package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

// Given some Ventures already exist on the server
// When all Ventures are requested
// Then ensure the response code is 200
// And the 'Content-Type' header contains 'application/json'
// And 'Access-Control-Allow-Origin' is '*'
// And 'Access-Control-Allow-Headers' is '*'
// And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, HEAD, and OPTIONS
// And the body is a JSON array of valid Ventures
//
// TODO: Craft some test data and pre-inject it into a SQLite database
func TestGET_Ventures(t *testing.T) {
	req := APICall{
		URL:    "http://localhost:8080/ventures",
		Method: GET,
	}
	res := req.fire()
	defer res.Body.Close()

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
		"Access-Control-Allow-Methods": CORS_METHODS_PATTERN,
	})

	var ven []Venture
	err := json.NewDecoder(res.Body).Decode(&ven)
	require.Nil(t, err)
	AssertGenericVentures(t, ven)
}
