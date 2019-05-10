package uhttp

import (
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

// AssertHeaderExists asserts that a named header exists within a headers
// container
func AssertHeaderExists(t *testing.T, h http.Header, k string) {
	assert.NotEmpty(t, h.Get(k))
}

// AssertHeaderValue asserts that a named header exists with a specific value
// within a headers container
func AssertHeaderValue(t *testing.T, h http.Header, k string, exp string) {
	assert.Equal(t, exp, h.Get(k))
}

// SetupRequest sets up a request, recorder and response for tests
func SetupRequest(path string) (*http.Request, *http.ResponseWriter, *httptest.ResponseRecorder) {
	req, err := http.NewRequest("GET", "http://example.com"+path, nil)
	if err != nil {
		panic(err)
	}
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	return req, &res, rec
}
