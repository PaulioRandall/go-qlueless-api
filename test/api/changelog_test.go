package api

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
)

const changelogMediaType = "text/markdown"

var changelogDefaultMethods = []string{"GET", "HEAD", "OPTIONS"}

// TODO: Assert the body is a valid markdown
func TestGET_Changelog(t *testing.T) {
	t.Log(`Given a loaded changelog
		When the changelog is requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'text/markdown'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, HEAD, and OPTIONS
		And the body contains some data
		...`)

	req := APICall{
		URL:    "http://localhost:8080/changelog",
		Method: "GET",
	}
	res := req.fire()
	defer res.Body.Close()
	defer PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, changelogMediaType, changelogDefaultMethods)
	assertNotEmptyBody(t, res)
}

func TestHEAD_Changelog(t *testing.T) {
	t.Log(`Given a loaded changelog
		When the changelog is requested
		AND the HTTP method is 'HEAD'
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'text/markdown'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, HEAD, and OPTIONS
		And there is NO response body
		...`)

	req := APICall{
		URL:    "http://localhost:8080/changelog",
		Method: "HEAD",
	}
	res := req.fire()
	defer res.Body.Close()
	defer PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, changelogMediaType, changelogDefaultMethods)
	assertEmptyBody(t, res)
}

func TestOPTIONS_Changelog(t *testing.T) {
	t.Log(`Given a loaded changelog
		When the changelog is requested
		AND the HTTP method is 'OPTIONS'
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'text/markdown'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, HEAD, and OPTIONS
		And there is NO response body
		...`)

	req := APICall{
		URL:    "http://localhost:8080/changelog",
		Method: "OPTIONS",
	}
	res := req.fire()
	defer res.Body.Close()
	defer PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, changelogMediaType, changelogDefaultMethods)
	assertEmptyBody(t, res)
}
