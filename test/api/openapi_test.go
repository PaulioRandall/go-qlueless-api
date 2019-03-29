package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
)

// TODO: Assert the body is a valid OpenAPI specification
func TestGET_OpenAPI(t *testing.T) {
	t.Log(`Given a loaded OpenAPI specification
		When the specification is requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/vnd.oai.openapi+json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, HEAD, and OPTIONS
		And the body is a valid JSON object
		...`)

	req := APICall{
		URL:    "http://localhost:8080/openapi",
		Method: "GET",
	}
	res := req.fire()
	defer res.Body.Close()

	RequireStatusCode(t, 200, res)
	AssertHeadersEquals(t, res.Header, map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "*",
	})
	AssertHeadersContains(t, res.Header, map[string][]string{
		"Content-Type":                 []string{"application/vnd.oai.openapi+json"},
		"Access-Control-Allow-Methods": []string{"GET", "HEAD", "OPTIONS"},
	})
	AssertHeadersMatches(t, res.Header, map[string]string{
		"Access-Control-Allow-Methods": CORS_METHODS_PATTERN,
	})

	var spec map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&spec)
	require.Nil(t, err)
}
