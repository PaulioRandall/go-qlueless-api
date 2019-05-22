package test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
	"unicode"

	"github.com/PaulioRandall/go-cookies/cookies"
)

// APICall represents a single call to an API
type APICall struct {
	URL    string
	Method string
	Body   io.Reader
}

// newRequest is a file private function for creating new requests
func (c *APICall) newRequest() *http.Request {
	req, err := http.NewRequest(c.Method, c.URL, c.Body)
	if err != nil {
		log.Panic("newRequest(): ", err)
	}
	return req
}

// invokeRequest is a file private for carrying out requests
func (c *APICall) invokeRequest(req *http.Request) *http.Response {
	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	res, err := client.Do(req)
	if err != nil {
		log.Panic("invokeRequest(): ", err)
	}
	return res
}

// fire allows a built APICall to be actioned
func (c *APICall) Fire() *http.Response {
	req := c.newRequest()
	res := c.invokeRequest(req)
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

// SetWorkingDir sets the working directory so the server has access to
// resources.
func SetWorkingDir(binDir string) {
	abs, err := filepath.Abs(binDir)
	if err != nil {
		log.Panic(err)
	}

	err = os.Chdir(abs)
	if err != nil {
		log.Panic(err)
	}
}

// PrintTestDescription prints the test description to the test logs.
func PrintTestDescription(t *testing.T, desc string) {

	firstChar := func(s string) int {
		firstChar := -1
		for i, v := range s {
			if !unicode.IsSpace(v) {
				firstChar = i
				break
			}
		}
		return firstChar
	}

	desc = cookies.ForEachToken(desc, "\n", func(i int, s string) string {
		spaces := firstChar(s)
		if spaces > 0 {
			spaces -= 2
		}

		s = cookies.TrimPrefixSpace(s)
		if spaces > 0 {
			spaces *= 2
			s = strings.Repeat(" ", spaces) + s
		}

		return s
	})

	t.Log("\n" + desc + "\n")
}
