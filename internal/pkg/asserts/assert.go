package asserts

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func RequireStatusCode(t *testing.T, expected int, res *http.Response) {
	if expected != res.StatusCode {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		t.Log("Response: " + string(b))
		require.Equal(t, expected, res.StatusCode)
	}
}

func AssertGenericIntCSV(t *testing.T, csv string) {
	p := "^((0,)|(-?[1-9][0-9]*,)+)*(-?[1-9][0-9]*)$" // Example: 1,2,999,-123,-6
	match, _ := regexp.MatchString(p, csv)
	assert.True(t, match)
}

func AssertExactKeys(t *testing.T, expect []string, actual map[string]interface{}) {
	for _, k := range expect {
		assert.Contains(t, actual, k)
	}
	assert.Len(t, len(expect), len(actual))
}

func assertHeaderExists(t *testing.T, name string, h []string) bool {
	if len(h) < 1 {
		assert.Fail(t, "Expected header: "+name)
		return false
	}
	return true
}

func AssertHeadersEquals(t *testing.T, h http.Header, equals map[string]string) {
	for k, exp := range equals {
		act := h[k]
		if assertHeaderExists(t, k, act) {
			assert.Equal(t, exp, act[0])
		}
	}
}

func AssertHeadersNotEquals(t *testing.T, h http.Header, equals map[string]string) {
	for k, notExp := range equals {
		act := h[k]
		if assertHeaderExists(t, k, act) {
			assert.NotEqual(t, notExp, act[0])
		}
	}
}

func AssertHeadersContains(t *testing.T, h http.Header, contains map[string][]string) {
	for k, s := range contains {
		act := h[k]
		if assertHeaderExists(t, k, act) {
			for _, exp := range s {
				assert.Contains(t, act[0], exp)
			}
		}
	}
}

func AssertHeadersNotContains(t *testing.T, h http.Header, contains map[string][]string) {
	for k, s := range contains {
		act := h[k]
		if assertHeaderExists(t, k, act) {
			for notExp := range s {
				assert.NotContains(t, act[0], notExp)
			}
		}
	}
}

func AssertHeadersMatches(t *testing.T, h http.Header, patterns map[string]string) {
	for k, reg := range patterns {
		act := h[k]
		if assertHeaderExists(t, k, act) {
			assert.Regexp(t, reg, act[0])
		}
	}
}

func AssertHeadersNotMatches(t *testing.T, h http.Header, patterns map[string]string) {
	for k, reg := range patterns {
		act := h[k]
		if assertHeaderExists(t, k, act) {
			assert.NotRegexp(t, reg, act[0])
		}
	}
}
