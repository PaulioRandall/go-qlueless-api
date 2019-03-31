package asserts

import (
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"

	p "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

func PrintResponse(t *testing.T, body io.Reader) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	if len(b) > 0 {
		t.Log("\n" + string(b))
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

func _assertHeaderExists(t *testing.T, name string, h []string) bool {
	if len(h) < 1 {
		assert.Fail(t, "Expected header: "+name)
		return false
	}
	return true
}

func AssertHeadersEquals(t *testing.T, h http.Header, equals map[string]string) {
	for k, exp := range equals {
		act := h[k]
		if _assertHeaderExists(t, k, act) {
			assert.Equal(t, exp, act[0])
		}
	}
}

func AssertHeadersNotEquals(t *testing.T, h http.Header, equals map[string]string) {
	for k, notExp := range equals {
		act := h[k]
		if _assertHeaderExists(t, k, act) {
			assert.NotEqual(t, notExp, act[0])
		}
	}
}

func AssertHeadersContains(t *testing.T, h http.Header, contains map[string][]string) {
	for k, s := range contains {
		act := h[k]
		if _assertHeaderExists(t, k, act) {
			for _, exp := range s {
				assert.Contains(t, act[0], exp)
			}
		}
	}
}

func AssertHeadersNotContains(t *testing.T, h http.Header, contains map[string][]string) {
	for k, s := range contains {
		act := h[k]
		if _assertHeaderExists(t, k, act) {
			for notExp := range s {
				assert.NotContains(t, act[0], notExp)
			}
		}
	}
}

func AssertHeadersMatches(t *testing.T, h http.Header, patterns map[string]string) {
	for k, reg := range patterns {
		act := h[k]
		if _assertHeaderExists(t, k, act) {
			assert.Regexp(t, reg, act[0])
		}
	}
}

func AssertHeadersNotMatches(t *testing.T, h http.Header, patterns map[string]string) {
	for k, reg := range patterns {
		act := h[k]
		if _assertHeaderExists(t, k, act) {
			assert.NotRegexp(t, reg, act[0])
		}
	}
}

func AssertGenericError(t *testing.T, r p.WrappedReply) {
	assert.NotEmpty(t, r.Message)
	assert.NotEmpty(t, r.Self)
	assert.Empty(t, r.Data)
}

func AssertWrappedReply(t *testing.T, r p.WrappedReply) {
	assert.NotEmpty(t, r.Message)
	assert.NotEmpty(t, r.Self)
	assert.NotEmpty(t, r.Data)
}
