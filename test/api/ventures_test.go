//usr/bin/env go run "$0" "$@"; exit "$?"

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Venture struct {
	Description string `json:"description"`
	VentureID   string `json:"id,omitempty"`
	OrderIDs    string `json:"order_ids,omitempty"`
	State       string `json:"state"`
	IsAlive     bool   `json:"is_alive"`
	Extra       string `json:"extra,omitempty"`
}

func exit(exitCode *int) {
	os.Exit(*exitCode)
}

func TestMain(m *testing.M) {
	var exitCode int = 1
	cmd := startServer()

	defer exit(&exitCode)
	defer attemptRecover(&exitCode)
	defer stopServer(cmd)

	exitCode = m.Run()
	adminPrint(fmt.Sprintf("Test exit code: %d", exitCode))
}

func assertGenericVentures(t *testing.T, vens []Venture) {
	for _, v := range vens {
		assertGenericVenture(t, v)
	}
}

func assertGenericVenture(t *testing.T, ven Venture) {
	assert.Equal(t, ven.VentureID, "1")
	assert.NotEmpty(t, ven.Description)
	assert.NotEmpty(t, ven.State)
	if ven.OrderIDs != "" {
		assertGenericIntCSV(t, ven.OrderIDs)
	}
}

func assertGenericIntCSV(t *testing.T, csv string) {
	p := "^((0,)|(-?[1-9][0-9]*,)+)*(-?[1-9][0-9]*)$" // Example: 1,2,999,-123,-6
	match, _ := regexp.MatchString(p, csv)
	assert.True(t, match)
}

func assertExactKeys(t *testing.T, expect []string, actual map[string]interface{}) {
	for _, k := range expect {
		assert.Contains(t, actual, k)
	}
	assert.Len(t, len(expect), len(actual))
}

func assertHeadersEquals(t *testing.T, h http.Header, equals map[string]string) {
	for k, exp := range equals {
		act := h[k]
		assert.Equal(t, exp, act)
	}
}

func assertHeadersNotEquals(t *testing.T, h http.Header, equals map[string]string) {
	for k, notExp := range equals {
		act := h[k]
		assert.NotEqual(t, notExp, act)
	}
}

func assertHeadersContains(t *testing.T, h http.Header, contains map[string][]string) {
	for k, s := range contains {
		act := h[k]
		for exp := range s {
			assert.Contains(t, act, exp)
		}
	}
}

func assertHeadersNotContains(t *testing.T, h http.Header, contains map[string][]string) {
	for k, s := range contains {
		act := h[k]
		for notExp := range s {
			assert.NotContains(t, act, notExp)
		}
	}
}

func assertHeadersMatches(t *testing.T, h http.Header, patterns map[string]string) {
	for k, reg := range patterns {
		act := h[k]
		assert.Regexp(t, reg, act)
	}
}

func assertHeadersNotMatches(t *testing.T, h http.Header, patterns map[string]string) {
	for k, reg := range patterns {
		act := h[k]
		assert.NotRegexp(t, reg, act)
	}
}

// Given some Ventures already exist on the server
// When all Ventures are requested
// Then ensure the response code is 200
// And the 'Content-Type' header contains 'application/json'
// And the 'Access-Control-Allow-Origin' is '*'
// And the 'Access-Control-Allow-Headers' is '*'
// And the 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, HEAD, and OPTIONS
// And the body is a JSON array of valid Ventures
func TestGET_Ventures(t *testing.T) {
	req := APICall{
		URL:    "http://localhost:8080/ventures",
		Method: GET,
	}
	res := req.fire()
	defer res.Body.Close()

	if true {
		return // REMOVE when ready to start development
	}

	require.Equal(t, 200, res.StatusCode)
	assertHeadersEquals(t, res.Header, map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "*",
	})
	assertHeadersContains(t, res.Header, map[string][]string{
		"Content-Type":                 []string{"application/json"},
		"Access-Control-Allow-Methods": []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
	})
	assertHeadersMatches(t, res.Header, map[string]string{
		"Access-Control-Allow-Methods": "^((\\s*[a-zA-Z]*\\s*,)+)*(\\s*[a-zA-Z]*\\s*)$",
	})

	var ven []Venture
	err := json.NewDecoder(res.Body).Decode(&ven)
	require.Nil(t, err)
	assertGenericVentures(t, ven)
}
