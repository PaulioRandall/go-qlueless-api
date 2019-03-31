package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	p "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// assertDefaultHeaders asserts that the services default headers were applied
func assertDefaultHeaders(t *testing.T, res *http.Response, contentType string, allowedMethods []string) {
	AssertHeadersEquals(t, res.Header, map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "*",
		"Content-Type":                 contentType + "; charset=utf-8",
	})
	AssertHeadersContains(t, res.Header, map[string][]string{
		"Access-Control-Allow-Methods": allowedMethods,
	})
	AssertHeadersMatches(t, res.Header, map[string]string{
		"Access-Control-Allow-Methods": CORS_METHODS_PATTERN,
	})
}

// assertEmptyBody asserts that a response body is empty
func assertEmptyBody(t *testing.T, res *http.Response) {
	body, err := ioutil.ReadAll(res.Body)
	require.Nil(t, err)
	assert.Empty(t, body)
}

// assertNotEmptyBody asserts that a response body is NOT empty
func assertNotEmptyBody(t *testing.T, res *http.Response) {
	body, err := ioutil.ReadAll(res.Body)
	require.Nil(t, err)
	assert.NotEmpty(t, body)
}

// assertWrappedErrorBody assert that a response body is a generic error
func assertWrappedErrorBody(t *testing.T, res *http.Response) {
	var reply p.WrappedReply
	err := json.NewDecoder(res.Body).Decode(&reply)
	require.Nil(t, err)
	AssertGenericError(t, reply)
}
