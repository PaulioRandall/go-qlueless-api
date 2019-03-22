package pkg

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CheckIsPositive(t *testing.T, i int, m ...interface{}) {
	assert.Greater(t, 0, i, m)
}

func CheckNotBlank(t *testing.T, s string, m ...interface{}) {
	v := strings.TrimSpace(s)
	assert.NotEmpty(t, v, m)
}

func CheckThing(t *testing.T, w Thing) {
	CheckNotBlank(t, w.Description, "Thing.Description")
	CheckIsPositive(t, w.ID, "Thing.ID")
	CheckChildrenIds(t, w)
	CheckNotBlank(t, w.State, "Thing.State")
}

func CheckChildrenIds(t *testing.T, o Thing) {
	for _, id := range o.ChildrenIDs {
		CheckIsPositive(t, id, "Thing.ChildrenIDs")
	}
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

func CheckJSONResponseHeaders(t *testing.T, h http.Header) {
	CheckHeaderValue(t, h, "Content-Type", "application/json; charset=utf-8")
	CheckHeaderValue(t, h, "Access-Control-Allow-Origin", "*")
	CheckHeaderExists(t, h, "Access-Control-Allow-Methods")
	CheckHeaderExists(t, h, "Access-Control-Allow-Headers")
}
