package asserts

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func AssertHeadersEquals(t *testing.T, h http.Header, equals map[string]string) {
	for k, exp := range equals {
		act := h[k]
		assert.Equal(t, exp, act)
	}
}

func AssertHeadersNotEquals(t *testing.T, h http.Header, equals map[string]string) {
	for k, notExp := range equals {
		act := h[k]
		assert.NotEqual(t, notExp, act)
	}
}

func AssertHeadersContains(t *testing.T, h http.Header, contains map[string][]string) {
	for k, s := range contains {
		act := h[k]
		for exp := range s {
			assert.Contains(t, act, exp)
		}
	}
}

func AssertHeadersNotContains(t *testing.T, h http.Header, contains map[string][]string) {
	for k, s := range contains {
		act := h[k]
		for notExp := range s {
			assert.NotContains(t, act, notExp)
		}
	}
}

func AssertHeadersMatches(t *testing.T, h http.Header, patterns map[string]string) {
	for k, reg := range patterns {
		act := h[k]
		assert.Regexp(t, reg, act)
	}
}

func AssertHeadersNotMatches(t *testing.T, h http.Header, patterns map[string]string) {
	for k, reg := range patterns {
		act := h[k]
		assert.NotRegexp(t, reg, act)
	}
}
