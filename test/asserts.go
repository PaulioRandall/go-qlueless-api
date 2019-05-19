package test

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	toastify "github.com/PaulioRandall/go-cookies/toastify"
	wrapped "github.com/PaulioRandall/go-qlueless-api/shared/wrapped"

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

// PrintResponse prints the 'body' of a response to the test logs.
func PrintResponse(t *testing.T, body io.Reader) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	if len(b) > 0 {
		t.Log("\n" + string(b))
	}
}

// AssertCorsHeaders asserts that the response 'res' contains the expect CORS
// headers and values; including the endpoint dependent 'methods'.
func AssertCorsHeaders(t *testing.T, res *http.Response, methods string) {
	toastify.AssertHeadersEqual(t, res.Header, map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "*",
		"Access-Control-Allow-Methods": methods,
	})
}

// AssertDefaultHeaders asserts that the response 'res' default headers exist.
// This includes the value of dynamic headers 'contentType' and CORS 'methods'
func AssertDefaultHeaders(t *testing.T, res *http.Response, contentType, methods string) {
	AssertCorsHeaders(t, res, methods)
	toastify.AssertHeaderEqual(t, "Content-Type", res.Header, contentType+"; charset=utf-8")
}

// AssertEmptyBody asserts that a response 'body' is empty.
func AssertEmptyBody(t *testing.T, body io.Reader) {
	b, err := ioutil.ReadAll(body)
	require.Nil(t, err)
	assert.Empty(t, b)
}

// AssertNotEmptyBody asserts that a response 'body' is NOT empty. The body is
// returned.
func AssertNotEmptyBody(t *testing.T, body io.Reader) []byte {
	b, err := ioutil.ReadAll(body)
	require.Nil(t, err)
	assert.NotEmpty(t, b)
	return b
}

// AssertErrorBody assert that a response 'body' is a generic response error.
// Returns the parsed response error.
func AssertErrorBody(t *testing.T, body io.Reader) wrapped.WrappedReply {
	var reply wrapped.WrappedReply
	err := json.NewDecoder(body).Decode(&reply)
	require.Nil(t, err)
	wrapped.AssertGenericError(t, reply)
	return reply
}

// VerifyBadMethods asserts that for a specific 'url', the 'corsMethods' are as
// expected and an error response is returned when each value of 'badMethods'
// is used in an API call.
func VerifyBadMethods(t *testing.T, url string, corsMethods string, badMethods []string) {

	for _, m := range badMethods {

		req := APICall{
			URL:    url,
			Method: m,
		}
		res := req.Fire()

		defer res.Body.Close()
		defer PrintResponse(t, res.Body)

		ok := assert.Equal(t, 405, res.StatusCode, "Expected 405 response code for method '"+m+"'")
		if !ok {
			continue
		}

		AssertCorsHeaders(t, res, corsMethods)
		AssertEmptyBody(t, res.Body)
	}
}
