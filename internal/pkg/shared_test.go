package pkg

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// When given a valid int string, true is returned
func TestIsInt___1(t *testing.T) {
	assert.True(t, IsInt("123"))
}

// When given an invalid int string, false is returned
func TestIsInt___2(t *testing.T) {
	assert.False(t, IsInt("abc"))
}

// When given an invalid int string, false is returned
func TestIsInt___3(t *testing.T) {
	assert.False(t, IsInt("123abc"))
}

// When given a slice and index, the item is removed but other items remain
func TestDeleteStr___1(t *testing.T) {
	s := []string{"0", "1", "2", "3", "4"}
	act := DeleteStr(s, 2)
	assert.Contains(t, act, "0")
	assert.Contains(t, act, "1")
	assert.Contains(t, act, "3")
	assert.Contains(t, act, "4")
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
	req, res, rec := SetupRequest("/")

	Write500Reply(res, req)
	assert.Equal(t, 500, rec.Code)
}

// When invoked, writes JSON headers
func TestWrite500Reply___2(t *testing.T) {
	req, res, rec := SetupRequest("/")

	Write500Reply(res, req)
	CheckJSONResponseHeaders(t, (*rec).Header())
}

// When invoked, writes JSON headers
func TestWrite500Reply___3(t *testing.T) {
	req, res, rec := SetupRequest("/")

	Write500Reply(res, req)

	assert.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	assert.Nil(t, err)
	assert.Contains(t, m, "message")
	assert.Contains(t, m, "self")
	assert.Len(t, m, 2)
}

// When ReplyMeta.Message is set, returns true
func TestCheckReplyMetaMessage___1(t *testing.T) {
	req, res, _ := SetupRequest("/")
	r := ReplyMeta{
		Message: "message",
	}

	act := CheckReplyMetaMessage(res, req, r)
	assert.True(t, act)
}

// When ReplyMeta.Message not set, sets 500 status code and returns false
func TestCheckReplyMetaMessage___2(t *testing.T) {
	req, res, rec := SetupRequest("/")
	r := ReplyMeta{}

	act := CheckReplyMetaMessage(res, req, r)
	assert.False(t, act)
	assert.Equal(t, 500, rec.Code)
}

// When given a status between max and min, true is returned
func TestCheckStatusBetween___1(t *testing.T) {
	req, res, _ := SetupRequest("/")
	act := CheckStatusBetween(res, req, 404, 400, 499)
	assert.True(t, act)
}

// When given a status equal to max or min, true is returned
func TestCheckStatusCode___2(t *testing.T) {
	req, res, _ := SetupRequest("/")
	act := CheckStatusBetween(res, req, 400, 400, 499)
	assert.True(t, act)
	act = CheckStatusBetween(res, req, 499, 400, 499)
	assert.True(t, act)
}

// When given a status less than min or greater than max, false is returned and
// 500 status set in response
func TestCheckStatusCode___3(t *testing.T) {
	req, res, rec := SetupRequest("/")
	act := CheckStatusBetween(res, req, 300, 400, 499)
	assert.False(t, act)
	assert.Equal(t, 500, rec.Code)

	req, res, rec = SetupRequest("/")
	act = CheckStatusBetween(res, req, 500, 400, 499)
	assert.False(t, act)
	assert.Equal(t, 500, rec.Code)
}

// When not 4XX status code, sets 500 status code
func TestWrite4XXReply___1(t *testing.T) {
	req, res, rec := SetupRequest("/")
	r := ReplyMeta{
		Message: "message",
	}

	Write4XXReply(res, req, 300, r)
	assert.Equal(t, 500, rec.Code)
}

// When ReplyMeta.Message not set, sets 500 status code
func TestWrite4XXReply___2(t *testing.T) {
	req, res, rec := SetupRequest("/")
	r := ReplyMeta{}

	Write4XXReply(res, req, 400, r)
	assert.Equal(t, 500, rec.Code)
}

// When complete ReplyMeta passed, sets 200 status code
func TestWrite4XXReply___3(t *testing.T) {
	req, res, rec := SetupRequest("/search?q=dan+north")
	r := ReplyMeta{
		Message: "abc",
		Self:    (*req).URL.String(),
		Hints:   "xyz",
	}

	Write4XXReply(res, req, 400, r)
	assert.Equal(t, 400, rec.Code)
}

// When complete ReplyMeta passed, JSON headers are set
func TestWrite4XXReply___4(t *testing.T) {
	req, res, rec := SetupRequest("/search?q=dan+north")
	r := ReplyMeta{
		Message: "abc",
		Self:    (*req).URL.String(),
		Hints:   "xyz",
	}

	Write4XXReply(res, req, 400, r)
	CheckJSONResponseHeaders(t, (*rec).Header())
}

// When complete Reply4XX passed, body is set with expected JSON
func TestWrite4XXReply___5(t *testing.T) {
	req, res, rec := SetupRequest("/search?q=dan+north")
	r := ReplyMeta{
		Message: "abc",
		Self:    (*req).URL.Path + "?" + (*req).URL.RawQuery,
		Hints:   "xyz",
	}

	Write4XXReply(res, req, 400, r)

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
	req, res, rec := SetupRequest("/search?q=dan+north")
	r := ReplyMeta{
		Message: "abc",
		Hints:   "xyz",
	}

	Write4XXReply(res, req, 400, r)

	assert.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	assert.Nil(t, err)
	assert.Equal(t, "/search?q=dan+north", m["self"])
}

// When the 'wrap' query param is present in a request, returns true
func TestWrapReply___1(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/?wrap=", nil)
	assert.Nil(t, err)
	act := WrapReply(req)
	assert.True(t, act)
}

// When the 'wrap' query param is not present in a request, returns false
func TestIsMetaReply___2(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/?q=abc", nil)
	assert.Nil(t, err)
	act := WrapReply(req)
	assert.False(t, act)
}

// When 'wrap' not present and data is nil, nil is returned
func TestPrepResponseData___1(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/", nil)
	assert.Nil(t, err)

	act := PrepResponseData(req, nil, "ignored")
	assert.Nil(t, act)
}

// When 'wrap' not present and data is provided, data is returned unchanged
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
	req, err := http.NewRequest("GET", "http://example.com/?wrap", nil)
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

// When invoked, the CORS response headers are set
func TestAppendCORSHeaders___1(t *testing.T) {
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	AppendCORSHeaders(&res)
	CheckCORSResponseHeaders(t, rec.Header())
}

// When invoked, the JSON response headers are set
func TestAppendJSONHeaders___1(t *testing.T) {
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	AppendJSONHeaders(&res, "")
	CheckJSONResponseHeaders(t, rec.Header())
}

// When given valid inputs, 200 status code is set
func TestWriteReply___1(t *testing.T) {
	_, res, rec := SetupRequest("/")
	b := []byte("Ghost in the moon")

	WriteReply(res, &b, "text/plain")
	assert.Equal(t, 200, rec.Code)
}

// When given valid inputs, the response headers are set
func TestWriteReply___2(t *testing.T) {
	_, res, rec := SetupRequest("/")
	b := []byte("Ghost in the moon")

	WriteReply(res, &b, "text/plain")
	CheckCORSResponseHeaders(t, rec.Header())
	CheckHeaderValue(t, rec.Header(), "Content-Type", "text/plain")
}

// When given valid inputs, the data is written to the response body
func TestWriteReply___3(t *testing.T) {
	_, res, rec := SetupRequest("/")
	b := []byte("Ghost in the moon")

	WriteReply(res, &b, "text/plain")

	assert.NotNil(t, rec.Body)
	s := string(rec.Body.String())
	assert.Equal(t, "Ghost in the moon", s)
}

// When given valid inputs, 200 status code is set
func TestWriteEmptyReply___1(t *testing.T) {
	_, res, rec := SetupRequest("/")
	WriteEmptyReply(res, "text/plain")
	assert.Equal(t, 200, rec.Code)
}

// When given valid inputs, the response headers are set
func TestWriteEmptyReply___2(t *testing.T) {
	_, res, rec := SetupRequest("/")
	WriteEmptyReply(res, "text/plain")
	CheckCORSResponseHeaders(t, rec.Header())
	CheckHeaderValue(t, rec.Header(), "Content-Type", "text/plain")
}

// When given valid inputs, no data is written to the response body
func TestWriteEmptyReply___3(t *testing.T) {
	_, res, rec := SetupRequest("/")
	WriteEmptyReply(res, "text/plain")
	assert.Empty(t, rec.Body)
}

// When given valid inputs, 200 status code is set
func TestWriteJsonReply___1(t *testing.T) {
	req, res, rec := SetupRequest("/")
	m := make(map[string]interface{})
	m["killswitch"] = "engage"

	WriteJSONReply(res, req, m, "")
	assert.Equal(t, 200, rec.Code)
}

// When given valid inputs, the JSON response headers are set
func TestWriteJsonReply___2(t *testing.T) {
	req, res, rec := SetupRequest("/")
	m := make(map[string]interface{})
	m["killswitch"] = "engage"

	WriteJSONReply(res, req, m, "")
	CheckJSONResponseHeaders(t, rec.Header())
}

// When given valid inputs, the data is serialised into JSON the response body
// is set
func TestWriteJsonReply___3(t *testing.T) {
	req, res, rec := SetupRequest("/")
	data := make(map[string]interface{})
	data["killswitch"] = "engage"

	WriteJSONReply(res, req, data, "")

	assert.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	assert.Nil(t, err)
	assert.Equal(t, data, m)
}

// When given valid inputs, 200 status code is set
func TestWriteEmptyJSONReply___1(t *testing.T) {
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	WriteEmptyJSONReply(&res, "")
	assert.Equal(t, 200, rec.Code)
}

// When given valid inputs, the JSON response headers are set
func TestWriteEmptyJSONReply___2(t *testing.T) {
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	WriteEmptyJSONReply(&res, "")
	CheckJSONResponseHeaders(t, rec.Header())
}

// When given valid inputs, no response body is set
func TestWriteEmptyJSONReply___3(t *testing.T) {
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	WriteEmptyJSONReply(&res, "")
	assert.Empty(t, rec.Body)
}
