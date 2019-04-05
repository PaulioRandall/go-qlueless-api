package test

import (
	"testing"

	a "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"

	require "github.com/stretchr/testify/require"
)

const changelogMediaType = "text/markdown"

var changelogHttpMethods = []string{"GET", "OPTIONS"}

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

	StartServer("")
	defer StopServer()

	req := APICall{
		URL:    "http://localhost:8080/changelog",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	AssertDefaultHeaders(t, res, changelogMediaType, changelogHttpMethods)
	AssertNotEmptyBody(t, res.Body)
}

// ****************************************************************************
// (OPTIONS) /changelog
// ****************************************************************************

func TestOPTIONS_Changelog(t *testing.T) {
	t.Log(`Given a loaded changelog
		When only /changelog OPTIONS are requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'text/markdown'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET and OPTIONS
		And there is NO response body
		...`)

	StartServer("")
	defer StopServer()

	req := APICall{
		URL:    "http://localhost:8080/changelog",
		Method: "OPTIONS",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	AssertNoContentHeaders(t, res, changelogHttpMethods)
	AssertEmptyBody(t, res.Body)
}

// ****************************************************************************
// (?) /changelog
// ****************************************************************************

func TestINVALID_Changelog(t *testing.T) {
	t.Log(`Given a loaded changelog
	  When /changelog is called using invalid methods
		Then ensure the response code is 405
		And the 'Content-Type' header contains 'application/vnd.oai.openapi+json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET and OPTIONS
		And there is NO response body
		...`)

	StartServer("")
	defer StopServer()

	VerifyNotAllowedMethods(t, "http://localhost:8080/changelog", changelogHttpMethods)
}
