package pkg

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func CheckCORSResponseHeaders(t *testing.T, h http.Header) {
	CheckHeaderValue(t, h, "Access-Control-Allow-Origin", "*")
	CheckHeaderExists(t, h, "Access-Control-Allow-Methods")
	CheckHeaderExists(t, h, "Access-Control-Allow-Headers")
}

func CheckJSONResponseHeaders(t *testing.T, h http.Header) {
	CheckHeaderValue(t, h, "Content-Type", "application/json; charset=utf-8")
}

func NewDummyThing() Thing {
	return Thing{
		Description: "# Outline the saga\nCreate a rough outline of the new saga.",
		State:       "in_progress",
		IsDead:      false,
	}
}

func DummyThing() Thing {
	return Thing{
		Description: "# Outline the saga\nCreate a rough outline of the new saga.",
		State:       "in_progress",
		IsDead:      false,
		Self:        "/things/1",
		ID:          1,
	}
}

func DummyThings() *[]Thing {
	return &[]Thing{
		Thing{
			Description: "# Outline the saga\nCreate a rough outline of the new saga.",
			ID:          1,
			ChildrenIDs: []int{
				2,
				3,
				4,
			},
			State: "in_progress",
		},
		Thing{
			Description: "# Name the saga\nThink of a name for the saga.",
			ID:          2,
			State:       "potential",
		},
		Thing{
			Description: "# Outline the first chapter",
			ID:          3,
			State:       "delivered",
			Additional:  "archive_note:Done but not a compelling start",
		},
		Thing{
			Description: "# Outline the second chapter",
			ID:          4,
			State:       "in_progress",
		},
	}
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
