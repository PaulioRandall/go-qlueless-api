package api

import (
	"io"
	"log"
	"net/http"
	"time"
)

type APICall struct {
	URL    string
	Method string
	Body   io.Reader
}

func (c *APICall) _newRequest() *http.Request {
	req, err := http.NewRequest(c.Method, c.URL, c.Body)
	if err != nil {
		log.Panic("newRequest(): ", err)
	}
	return req
}

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

func (c *APICall) fire() *http.Response {
	req := c._newRequest()
	res := c._invokeRequest(req)
	return res
}
