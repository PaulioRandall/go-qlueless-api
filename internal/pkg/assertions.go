package pkg

import (
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CheckIsNumber(t *testing.T, s string, m ...interface{}) {
	n, err := strconv.Atoi(s)
	assert.Nil(t, err, "Expected string to be a number")
	s2 := strconv.Itoa(n)
	assert.Equal(t, s, s2, "Expected stringified number to equal the original string")
}

func CheckNotBlank(t *testing.T, s string, m ...interface{}) {
	v := strings.TrimSpace(s)
	assert.NotEmpty(t, v, m)
}

func CheckThing(t *testing.T, w Thing) {
	CheckNotBlank(t, w.Description, "Thing.Description")
	CheckNotBlank(t, w.ID, "Thing.Work_item_id")
	CheckIsNumber(t, w.ID, "Thing.Work_item_id")
	CheckNotBlank(t, w.TagID, "Thing.Tag_id")
	CheckNotBlank(t, w.StatusID, "Thing.Status_id")
}

func CheckOrder(t *testing.T, o Thing) {
	CheckThing(t, o)
}

func CheckBatch(t *testing.T, b Thing) {
	CheckThing(t, b)
	CheckNotBlank(t, b.ParentID, "Thing.Parent_id")
	CheckIsNumber(t, b.ParentID, "Thing.Parent_id")
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
