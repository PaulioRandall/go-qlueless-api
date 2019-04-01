package uhttp

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func CheckIsPositive(t *testing.T, i int, m ...interface{}) {
	if i < 1 {
		assert.Fail(t, "Not positive", m)
	}
}

func CheckNotBlank(t *testing.T, s string, m ...interface{}) {
	v := strings.TrimSpace(s)
	assert.NotEmpty(t, v, m)
}

func CheckPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			assert.Fail(t, "Expected code to panic but it didn't")
		}
	}()
	f()
}

func CheckHeaderExists(t *testing.T, h http.Header, k string) {
	assert.NotEmpty(t, h.Get(k))
}

func CheckHeaderValue(t *testing.T, h http.Header, k string, exp string) {
	assert.Equal(t, exp, h.Get(k))
}

func SetupRequest(path string) (*http.Request, *http.ResponseWriter, *httptest.ResponseRecorder) {
	req, err := http.NewRequest("GET", "http://example.com"+path, nil)
	if err != nil {
		panic(err)
	}
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	return req, &res, rec
}
