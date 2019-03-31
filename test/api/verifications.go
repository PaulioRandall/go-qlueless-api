package api

import (
	"testing"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"

	"github.com/stretchr/testify/assert"
)

const (
	CORS_METHODS_PATTERN = "^((\\s*[A-Z]*\\s*,)+)*(\\s*[A-Z]*\\s*)$" // Example: 'GET,  POST   ,OPTIONS'
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

// verifyNotAllowedMethods asserts that the supplied methods are not allowed
// for provided endpoint
func verifyNotAllowedMethods(t *testing.T, url string, allowedMethods []string) {

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
		res := req.fire()
		defer res.Body.Close()
		defer PrintResponse(t, res.Body)

		ok := assert.Equal(t, 405, res.StatusCode, "Expected method not allowed: ("+m+")")
		if !ok {
			continue
		}
		assertDefaultHeaders(t, res, "application/json", allowedMethods)
		assertWrappedErrorBody(t, res)
	}
}
