package changelog

import (
	"testing"

	server "github.com/PaulioRandall/go-qlueless-api/api/server"
	test "github.com/PaulioRandall/go-qlueless-api/test"
	require "github.com/stretchr/testify/require"
)

func init() {
	test.SetWorkingDir("../../bin")
}

// ****************************************************************************
// (GET) /changelog
// ****************************************************************************

// TODO: Assert the body is a valid markdown
func TestGET_Changelog(t *testing.T) {
	t.Log(`Given a loaded changelog
		When the changelog is requested
		Then ensure the response code is 405
		And the 'Content-Type' header contains 'text/markdown'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET and OPTIONS
		And the body contains some data
		...`)

	server.StartUp(true)
	defer server.Shutdown()

	req := test.APICall{
		URL:    "http://localhost:8080/changelog",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "text/markdown", "GET, OPTIONS")
	test.AssertNotEmptyBody(t, res.Body)
}

// ****************************************************************************
// (OPTIONS) /changelog
// ****************************************************************************

func TestOPTIONS_Changelog(t *testing.T) {
	t.Log(`Given a loaded changelog
		When only /changelog OPTIONS are requested
		Then ensure the response code is 200
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET and OPTIONS
		And there is NO response body
		...`)

	server.StartUp(true)
	defer server.Shutdown()

	req := test.APICall{
		URL:    "http://localhost:8080/changelog",
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
// (?) /changelog
// ****************************************************************************

func TestINVALID_Changelog(t *testing.T) {
	t.Log(`Given a loaded changelog
	  When /changelog is called using invalid methods
		Then ensure the response code is 405
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' is 'GET, OPTIONS'
		And there is NO response body
		...`)

	server.StartUp(true)
	defer server.Shutdown()

	test.VerifyBadMethods(t, "http://localhost:8080/changelog", "GET, OPTIONS", []string{
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
