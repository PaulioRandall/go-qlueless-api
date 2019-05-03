package test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

// APICall represents a single call to an API
type APICall struct {
	URL    string
	Method string
	Body   io.Reader
}

// _newRequest is a file private function for creating new requests
func (c *APICall) _newRequest() *http.Request {
	req, err := http.NewRequest(c.Method, c.URL, c.Body)
	if err != nil {
		log.Panic("newRequest(): ", err)
	}
	return req
}

// _invokeRequest is a file private for carrying out requests
func (c *APICall) _invokeRequest(req *http.Request) *http.Response {
	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	res, err := client.Do(req)
	if err != nil {
		log.Panic("InvokeRequest(): ", err)
	}
	return res
}

// fire allows a built APICall to be actioned
func (c *APICall) Fire() *http.Response {
	req := c._newRequest()
	res := c._invokeRequest(req)
	return res
}

// CallWithJSON calls an API endpoint with the specified details.
func CallWithJSON(method string, url string, data interface{}) *http.Response {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&data)

	req := APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "PUT",
		Body:   buf,
	}

	return req.Fire()
}
