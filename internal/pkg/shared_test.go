package pkg

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dummyThings() *[]Thing {
	return &[]Thing{
		Thing{
			Description: "# Outline the saga\nCreate a rough outline of the new saga.",
			ID:          "1",
			ChildrenIDs: []string{
				"2",
				"3",
				"4",
			},
			State: "in_progress",
		},
		Thing{
			Description: "# Name the saga\nThink of a name for the saga.",
			ID:          "2",
			State:       "potential",
		},
		Thing{
			Description: "# Outline the first chapter",
			ID:          "3",
			State:       "delivered",
			Additional:  "archive_note:Done but not a compelling start",
		},
		Thing{
			Description: "# Outline the second chapter",
			ID:          "4",
			State:       "in_progress",
		},
	}
}

func createRequest(path string) (*http.Request, *http.ResponseWriter, *httptest.ResponseRecorder) {
	req, err := http.NewRequest("GET", "http://example.com"+path, nil)
	if err != nil {
		panic(err)
	}
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	return req, &res, rec
}

// When given an empty string, true is returned
func TestIsBlank___1(t *testing.T) {
	act := IsBlank("")
	assert.True(t, act)
}

// When given a string with whitespace, true is returned
func TestIsBlank___2(t *testing.T) {
	act := IsBlank("\r\n \t\f")
	assert.True(t, act)
}

// When given a string with no whitespaces, false is returned
func TestIsBlank___3(t *testing.T) {
	act := IsBlank("Captain Vimes")
	assert.False(t, act)
}

// When given a request, returns the absolute relative URL of the request
func TestRelURL(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/character?q=Nobby", nil)
	assert.Nil(t, err)

	act := RelURL(req)
	assert.Equal(t, "/character?q=Nobby", act)
}

// When a value is present, it is returned
func TestValueOrEmpty___1(t *testing.T) {
	m := make(map[string]interface{})
	m["key"] = "value"
	act := ValueOrEmpty(m, "key")
	assert.Equal(t, "value", act)
}

// When a value is not present, empty string is returned
func TestValueOrEmpty___2(t *testing.T) {
	m := make(map[string]interface{})
	m["key"] = "value"
	act := ValueOrEmpty(m, "responsibilities")
	assert.Empty(t, act)
}

// When a value is present, it is returned
func TestValueOrFalse___1(t *testing.T) {
	m := make(map[string]interface{})
	m["key"] = true
	act := ValueOrFalse(m, "key")
	assert.True(t, act)
}

// When a value is not present, false is returned
func TestValueOrFalse___2(t *testing.T) {
	m := make(map[string]interface{})
	m["key"] = "value"
	act := ValueOrFalse(m, "responsibilities")
	assert.False(t, act)
}

// When a value is present, it is returned
func TestValueOrEmptyArray___1(t *testing.T) {
	m := make(map[string]interface{})
	m["key"] = []string{"1", "2"}
	act := ValueOrEmptyArray(m, "key")
	assert.Equal(t, []string{"1", "2"}, act)
}

// When a value is not present, empty array is returned
func TestValueOrEmptyArray___2(t *testing.T) {
	m := make(map[string]interface{})
	m["key"] = []string{"1", "2"}
	act := ValueOrEmptyArray(m, "responsibilities")
	assert.Empty(t, act)
}

// When given nil, does nothing
func TestCheck___1(t *testing.T) {
	Check(nil)
}

// When given an error, a panics ensues
func TestCheck___2(t *testing.T) {
	CheckPanic(t, func() {
		err := errors.New("Computer says no!")
		Check(err)
	})
}

// When given nil, returns false and doesn't print anything
func TestLogIfErr___1(t *testing.T) {
	act := LogIfErr(nil)
	assert.False(t, act)
	// Output:
	//
}

// When given an error, returns true and prints the error message
func TestLogIfErr___2(t *testing.T) {
	err := errors.New("Computer says no!")
	act := LogIfErr(err)
	assert.True(t, act)
	// Output:
	// Computer says no!
}

// When invoked, sets 500 status code
func TestWrite500Reply___1(t *testing.T) {
	req, res, rec := createRequest("/")

	Write500Reply(res, req)
	assert.Equal(t, 500, rec.Code)
}

// When invoked, writes JSON headers
func TestWrite500Reply___2(t *testing.T) {
	req, res, rec := createRequest("/")

	Write500Reply(res, req)
	CheckJSONResponseHeaders(t, (*rec).Header())
}

// When invoked, writes JSON headers
func TestWrite500Reply___3(t *testing.T) {
	req, res, rec := createRequest("/")

	Write500Reply(res, req)

	assert.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	assert.Nil(t, err)
	assert.Contains(t, m, "message")
	assert.Contains(t, m, "self")
	assert.Len(t, m, 2)
}

// When not 4XX status code, sets 500 status code
func TestWrite4XXReply___1(t *testing.T) {
	req, res, rec := createRequest("/")
	r := Reply4XX{
		Req:     req,
		Res:     res,
		Message: "message",
	}

	Write4XXReply(300, &r)
	assert.Equal(t, 500, rec.Code)
}

// When Reply4XX.Message not set, sets 500 status code
func TestWrite4XXReply___2(t *testing.T) {
	req, res, rec := createRequest("/")
	r := Reply4XX{
		Req: req,
		Res: res,
	}

	Write4XXReply(400, &r)
	assert.Equal(t, 500, rec.Code)
}

// When complete Reply4XX passed, sets 200 status code
func TestWrite4XXReply___3(t *testing.T) {
	req, res, rec := createRequest("/search?q=dan+north")
	r := Reply4XX{
		Req:     req,
		Res:     res,
		Message: "abc",
		Self:    (*req).URL.String(),
		Hints:   "xyz",
	}

	Write4XXReply(400, &r)
	assert.Equal(t, 400, rec.Code)
}

// When complete Reply4XX passed, JSON headers are set
func TestWrite4XXReply___4(t *testing.T) {
	req, res, rec := createRequest("/search?q=dan+north")
	r := Reply4XX{
		Req:     req,
		Res:     res,
		Message: "abc",
		Self:    (*req).URL.String(),
		Hints:   "xyz",
	}

	Write4XXReply(400, &r)
	CheckJSONResponseHeaders(t, (*rec).Header())
}

// When complete Reply4XX passed, body is set with expected JSON
func TestWrite4XXReply___5(t *testing.T) {
	req, res, rec := createRequest("/search?q=dan+north")
	r := Reply4XX{
		Req:     req,
		Res:     res,
		Message: "abc",
		Self:    (*req).URL.Path + "?" + (*req).URL.RawQuery,
		Hints:   "xyz",
	}

	Write4XXReply(400, &r)

	assert.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)

	assert.Nil(t, err)
	assert.Equal(t, "abc", m["message"])
	assert.Equal(t, "/search?q=dan+north", m["self"])
	assert.Equal(t, "xyz", m["hints"])
	assert.Len(t, m, 3)
}

// When Reply4XX.Self is not set, Reply4XX.Self is set for us
func TestWrite4XXReply___6(t *testing.T) {
	req, res, rec := createRequest("/search?q=dan+north")
	r := Reply4XX{
		Req:     req,
		Res:     res,
		Message: "abc",
		Hints:   "xyz",
	}

	Write4XXReply(400, &r)

	assert.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	assert.Nil(t, err)
	assert.Equal(t, "/search?q=dan+north", m["self"])
}

// When the 'meta' query param is present in a request, returns true
func TestIsMetaReply___1(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/?meta=", nil)
	assert.Nil(t, err)
	act := IsMetaReply(req)
	assert.True(t, act)
}

// When the 'meta' query param is not present in a request, returns false
func TestIsMetaReply___2(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/?q=abc", nil)
	assert.Nil(t, err)
	act := IsMetaReply(req)
	assert.False(t, act)
}

// When 'meta' not present and data is nil, nil is returned
func TestPrepResponseData___1(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/", nil)
	assert.Nil(t, err)

	act := PrepResponseData(req, nil, "ignored")
	assert.Nil(t, act)
}

// When 'meta' not present and data is provided, data is returned unchanged
func TestPrepResponseData___2(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/", nil)
	assert.Nil(t, err)
	data := make(map[string]interface{})
	data["album"] = "As Daylight Dies"

	act := PrepResponseData(req, data, "ignored")
	assert.NotNil(t, act)
	assert.Equal(t, data, act)
}

// When 'meta' is present and data is provided, wrapped reply is returned
func TestPrepResponseData___3(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/?meta", nil)
	assert.Nil(t, err)

	data := make(map[string]interface{})
	data["album"] = "As Daylight Dies"

	exp := ReplyMeta{
		Message: "Cheese",
		Self:    req.URL.String(),
		Data:    data,
	}

	act := PrepResponseData(req, data, "Cheese")
	assert.NotNil(t, act)
	assert.Equal(t, exp, act)
}

// When invoked, the JSON response headers are set
func TestAppendJSONHeaders___1(t *testing.T) {
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	AppendJSONHeaders(&res)
	CheckJSONResponseHeaders(t, rec.Header())
}

// When given valid inputs, 200 status code is set
func TestWriteReply___1(t *testing.T) {
	req, res, rec := createRequest("/")
	m := make(map[string]interface{})
	m["killswitch"] = "engage"

	WriteReply(res, req, m)
	assert.Equal(t, 200, rec.Code)
}

// When given valid inputs, the JSON response headers are set
func TestWriteReply___2(t *testing.T) {
	req, res, rec := createRequest("/")
	m := make(map[string]interface{})
	m["killswitch"] = "engage"

	WriteReply(res, req, m)
	CheckJSONResponseHeaders(t, rec.Header())
}

// When given valid inputs, the data is serialised into JSON the response body
// is set
func TestWriteReply___3(t *testing.T) {
	req, res, rec := createRequest("/")
	data := make(map[string]interface{})
	data["killswitch"] = "engage"

	WriteReply(res, req, data)

	assert.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	assert.Nil(t, err)
	assert.Equal(t, data, m)
}

// When given valid inputs, 200 status code is set
func TestWriteEmptyReply___1(t *testing.T) {
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	WriteEmptyReply(&res)
	assert.Equal(t, 200, rec.Code)
}

// When given valid inputs, the JSON response headers are set
func TestWriteEmptyReply___2(t *testing.T) {
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	WriteEmptyReply(&res)
	CheckJSONResponseHeaders(t, rec.Header())
}

// When given valid inputs, no response body is set
func TestWriteEmptyReply___3(t *testing.T) {
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	WriteEmptyReply(&res)
	assert.Empty(t, rec.Body)
}
