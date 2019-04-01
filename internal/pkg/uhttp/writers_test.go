package uhttp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	a "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	w "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/wrapped"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// ****************************************************************************
// WriteServerError()
// ****************************************************************************

func TestWriteServerError___1(t *testing.T) {
	req, res, rec := SetupRequest("/")

	WriteServerError(res, req)
	assert.Equal(t, 500, rec.Code)
}

func TestWriteServerError___2(t *testing.T) {
	req, res, rec := SetupRequest("/")

	WriteServerError(res, req)
	a.AssertHeadersEquals(t, (*rec).Header(), map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	})
}

func TestWriteServerError___3(t *testing.T) {
	req, res, rec := SetupRequest("/")

	WriteServerError(res, req)

	require.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	require.Nil(t, err)
	require.Len(t, m, 2)
	assert.Contains(t, m, "message")
	assert.Contains(t, m, "self")
}

// ****************************************************************************
// WriteBadRequest()
// ****************************************************************************

func TestWriteBadRequest___1(t *testing.T) {
	req, res, rec := SetupRequest("/")
	WriteBadRequest(res, req, "message")
	assert.Equal(t, 400, rec.Code)
}

func TestWriteBadRequest___2(t *testing.T) {
	req, res, rec := SetupRequest("/")

	WriteBadRequest(res, req, "message")
	exp := w.WrappedReply{
		Message: "message",
		Self:    "/",
	}

	require.NotNil(t, rec.Body)
	var a w.WrappedReply
	err := json.NewDecoder(rec.Body).Decode(&a)

	require.Nil(t, err)
	assert.Equal(t, exp, a)
}

func TestWriteBadRequest___3(t *testing.T) {
	req, res, rec := SetupRequest("/")
	WriteBadRequest(res, req, "")
	assert.Equal(t, 500, rec.Code)
}

// ****************************************************************************
// Write4XXReply()
// ****************************************************************************

func TestWrite4XXReply___1(t *testing.T) {
	req, res, rec := SetupRequest("/")
	r := w.WrappedReply{
		Message: "message",
	}

	Write4XXReply(res, req, 300, r)
	assert.Equal(t, 500, rec.Code)
}

func TestWrite4XXReply___2(t *testing.T) {
	req, res, rec := SetupRequest("/")
	r := w.WrappedReply{}

	Write4XXReply(res, req, 400, r)
	assert.Equal(t, 500, rec.Code)
}

func TestWrite4XXReply___3(t *testing.T) {
	req, res, rec := SetupRequest("/search?q=dan+north")
	r := w.WrappedReply{
		Message: "abc",
		Self:    (*req).URL.String(),
		Hints:   "xyz",
	}

	Write4XXReply(res, req, 400, r)
	assert.Equal(t, 400, rec.Code)
}

func TestWrite4XXReply___4(t *testing.T) {
	req, res, rec := SetupRequest("/search?q=dan+north")
	r := w.WrappedReply{
		Message: "abc",
		Self:    (*req).URL.String(),
		Hints:   "xyz",
	}

	Write4XXReply(res, req, 400, r)
	a.AssertHeadersEquals(t, (*rec).Header(), map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	})
}

func TestWrite4XXReply___5(t *testing.T) {
	req, res, rec := SetupRequest("/search?q=dan+north")
	r := w.WrappedReply{
		Message: "abc",
		Self:    (*req).URL.Path + "?" + (*req).URL.RawQuery,
		Hints:   "xyz",
	}

	Write4XXReply(res, req, 400, r)

	require.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)

	require.Nil(t, err)
	require.Len(t, m, 3)
	assert.Equal(t, "abc", m["message"])
	assert.Equal(t, "/search?q=dan+north", m["self"])
	assert.Equal(t, "xyz", m["hints"])
}

func TestWrite4XXReply___6(t *testing.T) {
	req, res, rec := SetupRequest("/search?q=dan+north")
	r := w.WrappedReply{
		Message: "abc",
		Hints:   "xyz",
	}

	Write4XXReply(res, req, 400, r)

	require.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	require.Nil(t, err)
	assert.Equal(t, "/search?q=dan+north", m["self"])
}

// ****************************************************************************
// WriteReply()
// ****************************************************************************

func TestWriteReply___1(t *testing.T) {
	_, res, rec := SetupRequest("/")
	b := []byte("Ghost in the moon")

	WriteReply(res, &b, "text/plain")
	assert.Equal(t, 200, rec.Code)
}

func TestWriteReply___2(t *testing.T) {
	_, res, rec := SetupRequest("/")
	b := []byte("Ghost in the moon")

	WriteReply(res, &b, "text/plain")
	CheckHeaderValue(t, rec.Header(), "Content-Type", "text/plain")
}

func TestWriteReply___3(t *testing.T) {
	_, res, rec := SetupRequest("/")
	b := []byte("Ghost in the moon")

	WriteReply(res, &b, "text/plain")

	require.NotNil(t, rec.Body)
	s := string(rec.Body.String())
	assert.Equal(t, "Ghost in the moon", s)
}

// ****************************************************************************
// WriteEmptyReply()
// ****************************************************************************

func TestWriteEmptyReply___1(t *testing.T) {
	_, res, rec := SetupRequest("/")
	WriteEmptyReply(res, "text/plain")
	assert.Equal(t, 200, rec.Code)
}

func TestWriteEmptyReply___2(t *testing.T) {
	_, res, rec := SetupRequest("/")
	WriteEmptyReply(res, "text/plain")
	CheckHeaderValue(t, rec.Header(), "Content-Type", "text/plain")
}

func TestWriteEmptyReply___3(t *testing.T) {
	_, res, rec := SetupRequest("/")
	WriteEmptyReply(res, "text/plain")
	assert.Empty(t, rec.Body)
}

// ****************************************************************************
// WriteJSONReply()
// ****************************************************************************

func TestWriteJsonReply___1(t *testing.T) {
	req, res, rec := SetupRequest("/")
	m := make(map[string]interface{})
	m["killswitch"] = "engage"

	WriteJSONReply(res, req, m, "")
	assert.Equal(t, 200, rec.Code)
}

func TestWriteJsonReply___2(t *testing.T) {
	req, res, rec := SetupRequest("/")
	m := make(map[string]interface{})
	m["killswitch"] = "engage"

	WriteJSONReply(res, req, m, "")
	a.AssertHeadersEquals(t, (*rec).Header(), map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	})
}

func TestWriteJsonReply___3(t *testing.T) {
	req, res, rec := SetupRequest("/")
	data := make(map[string]interface{})
	data["killswitch"] = "engage"

	WriteJSONReply(res, req, data, "")

	require.NotNil(t, rec.Body)
	var m map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&m)
	require.Nil(t, err)
	assert.Equal(t, data, m)
}

// ****************************************************************************
// WriteEmptyJSONReply()
// ****************************************************************************

func TestWriteEmptyJSONReply___1(t *testing.T) {
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	WriteEmptyJSONReply(&res, "")
	assert.Equal(t, 200, rec.Code)
}

func TestWriteEmptyJSONReply___2(t *testing.T) {
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	WriteEmptyJSONReply(&res, "")
	a.AssertHeadersEquals(t, (*rec).Header(), map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	})
}

func TestWriteEmptyJSONReply___3(t *testing.T) {
	rec := httptest.NewRecorder()
	var res http.ResponseWriter = rec
	WriteEmptyJSONReply(&res, "")
	assert.Empty(t, rec.Body)
}
