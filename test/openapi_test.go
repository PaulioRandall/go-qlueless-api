package test

import (
	"encoding/json"
	"testing"

	require "github.com/stretchr/testify/require"

	a "github.com/PaulioRandall/go-qlueless-assembly-api/internal/asserts"
)

const openapiMediaType = "application/vnd.oai.openapi+json"

var openapiHttpMethods = []string{"GET", "OPTIONS"}

// ****************************************************************************
// (GET) /openapi
// ****************************************************************************

// TODO: Assert the body is a valid OpenAPI specification
func TestGET_OpenAPI(t *testing.T) {
	t.Log(`Given a loaded OpenAPI specification
		When the specification is requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/vnd.oai.openapi+json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET and OPTIONS
		And the body is a valid JSON object
		...`)

	StartServer("")
	defer StopServer()

	req := APICall{
		URL:    "http://localhost:8080/openapi",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	AssertDefaultHeaders(t, res, openapiMediaType, openapiHttpMethods)

	var spec map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&spec)
	require.Nil(t, err)
}

// ****************************************************************************
// (OPTIONS) /openapi
// ****************************************************************************

func TestOPTIONS_OpenAPI(t *testing.T) {
	t.Log(`Given a loaded OpenAPI specification
		When only /openapi OPTIONS are requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/vnd.oai.openapi+json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET and OPTIONS
		And there is NO response body
		...`)

	StartServer("")
	defer StopServer()

	req := APICall{
		URL:    "http://localhost:8080/openapi",
		Method: "OPTIONS",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	AssertNoContentHeaders(t, res, openapiHttpMethods)
	AssertEmptyBody(t, res.Body)
}

// ****************************************************************************
// (?) /openapi
// ****************************************************************************

func TestINVALID_OpenAPI(t *testing.T) {
	t.Log(`Given a loaded OpenAPI specification
	  When /openapi is called using invalid methods
		Then ensure the response code is 405
		And the 'Content-Type' header contains 'application/vnd.oai.openapi+json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET and OPTIONS
		And there is NO response body
		...`)

	StartServer("")
	defer StopServer()

	VerifyNotAllowedMethods(t, "http://localhost:8080/openapi", openapiHttpMethods)
}
