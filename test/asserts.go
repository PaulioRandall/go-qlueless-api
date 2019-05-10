package test

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	a "github.com/PaulioRandall/go-qlueless-api/internal/asserts"
	w "github.com/PaulioRandall/go-qlueless-api/internal/wrapped"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

const (
	CORS_METHODS_PATTERN = "^((\\s*[A-Z]*\\s*,)+)*(\\s*[A-Z]*\\s*)$" // Example: 'GET,  POST, OPTIONS'
)

var ALL_STD_HTTP_METHODS = []string{
	"GET",
	"POST",
	"PUT",
	"DELETE",
	"HEAD",
	"OPTIONS",
	"CONNECT",
	"TRACE",
	"PATCH",
	"CUSTOM",
}

// AssertNoContentHeaders asserts that the services default headers were applied
// except the 'Content-Type' which should have been omitted
func AssertNoContentHeaders(t *testing.T, res *http.Response, allowedMethods []string) {
	a.AssertHeadersEquals(t, res.Header, map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "*",
	})
	a.AssertHeadersContains(t, res.Header, map[string][]string{
		"Access-Control-Allow-Methods": allowedMethods,
	})
	a.AssertHeadersMatches(t, res.Header, map[string]string{
		"Access-Control-Allow-Methods": CORS_METHODS_PATTERN,
	})
}

// AssertDefaultHeaders asserts that the services default headers were applied
func AssertDefaultHeaders(t *testing.T, res *http.Response, contentType string, allowedMethods []string) {
	AssertNoContentHeaders(t, res, allowedMethods)
	a.AssertHeadersEquals(t, res.Header, map[string]string{
		"Content-Type": contentType + "; charset=utf-8",
	})
}

// AssertEmptyBody asserts that a response body is empty
func AssertEmptyBody(t *testing.T, r io.Reader) {
	body, err := ioutil.ReadAll(r)
	require.Nil(t, err)
	assert.Empty(t, body)
}

// AssertNotEmptyBody asserts that a response body is NOT empty
func AssertNotEmptyBody(t *testing.T, r io.Reader) []byte {
	body, err := ioutil.ReadAll(r)
	require.Nil(t, err)
	assert.NotEmpty(t, body)
	return body
}

// AssertWrappedErrorBody assert that a response body is a generic error
func AssertWrappedErrorBody(t *testing.T, r io.Reader) w.WrappedReply {
	var reply w.WrappedReply
	err := json.NewDecoder(r).Decode(&reply)
	require.Nil(t, err)
	w.AssertGenericError(t, reply)
	return reply
}

// VerifyNotAllowedMethods asserts that the supplied methods are not allowed
// for provided endpoint
func VerifyNotAllowedMethods(t *testing.T, url string, allowedMethods []string) {

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
		res := req.Fire()
		defer res.Body.Close()
		defer a.PrintResponse(t, res.Body)

		ok := assert.Equal(t, 405, res.StatusCode, "Expected method not allowed: ("+m+")")
		if !ok {
			continue
		}
		AssertNoContentHeaders(t, res, allowedMethods)
		AssertEmptyBody(t, res.Body)
	}
}

// VerifyDefaultHeaders asserts that the default headers were provided
func VerifyDefaultHeaders(t *testing.T, c APICall, expCode int, allowedMethods []string) {
	res := c.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, expCode, res.StatusCode)
	AssertDefaultHeaders(t, res, "application/json", allowedMethods)
}
