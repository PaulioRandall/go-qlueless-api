package openapi

import (
	"encoding/json"
	"testing"

	test "github.com/PaulioRandall/go-qlueless-api/test"
	require "github.com/stretchr/testify/require"
)

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
		And 'Access-Control-Allow-Methods' is 'GET, OPTIONS'
		And the body is a valid JSON object
		...`)

	test.StartServer("../../bin")
	defer test.StopServer()

	req := test.APICall{
		URL:    "http://localhost:8080/openapi",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/vnd.oai.openapi+json", "GET, OPTIONS")

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
		And 'Access-Control-Allow-Methods' is 'GET, OPTIONS'
		And there is NO response body
		...`)

	test.StartServer("../../bin")
	defer test.StopServer()

	req := test.APICall{
		URL:    "http://localhost:8080/openapi",
		Method: "OPTIONS",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertCorsHeaders(t, res, "GET, OPTIONS")
	test.AssertEmptyBody(t, res.Body)
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
		And 'Access-Control-Allow-Methods' is 'GET, OPTIONS'
		And there is NO response body
		...`)

	test.StartServer("../../bin")
	defer test.StopServer()

	test.VerifyBadMethods(t, "http://localhost:8080/openapi", "GET, OPTIONS", []string{
		"POST",
		"PUT",
		"DELETE",
		"HEAD",
		"CONNECT",
		"TRACE",
		"PATCH",
		"CUSTOM",
	})
}
