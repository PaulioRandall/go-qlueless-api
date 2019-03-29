package api

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
)

const openapiMediaType = "application/vnd.oai.openapi+json"

var openapiDefaultMethods = []string{"GET", "HEAD", "OPTIONS"}

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
	defer PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, openapiMediaType, openapiDefaultMethods)

	var spec map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&spec)
	require.Nil(t, err)
}

func TestHEAD_OpenAPI(t *testing.T) {
	t.Log(`Given a loaded OpenAPI specification
		When the specification is requested
		AND the HTTP method is 'HEAD'
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/vnd.oai.openapi+json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, HEAD, and OPTIONS
		And there is NO response body
		...`)

	req := APICall{
		URL:    "http://localhost:8080/openapi",
		Method: "HEAD",
	}
	res := req.fire()
	defer res.Body.Close()
	defer PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, openapiMediaType, openapiDefaultMethods)

	body, err := ioutil.ReadAll(res.Body)
	require.Nil(t, err)
	assert.Empty(t, body)
}

func TestOPTIONS_OpenAPI(t *testing.T) {
	t.Log(`Given a loaded OpenAPI specification
		When the specification is requested
		AND the HTTP method is 'OPTIONS'
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/vnd.oai.openapi+json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, HEAD, and OPTIONS
		And there is NO response body
		...`)

	req := APICall{
		URL:    "http://localhost:8080/openapi",
		Method: "OPTIONS",
	}
	res := req.fire()
	defer res.Body.Close()
	defer PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, openapiMediaType, openapiDefaultMethods)

	body, err := ioutil.ReadAll(res.Body)
	require.Nil(t, err)
	assert.Empty(t, body)
}
