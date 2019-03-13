// Package internal/app contains non-reusable internal application code
package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

var reply shr.Reply = shr.Reply{
	Message: "All service dictionaries and their entries",
}
var once_dict sync.Once

// DictionaryHandler handles requests for the service dictionaries
func DictionaryHandler(w http.ResponseWriter, r *http.Request) {
	shr.Log_request(r)

	once_dict.Do(loadJson)

	if reply.Data == nil {
		shr.Http_500(&w)
		return
	}

	shr.WriteJsonReply(reply, w, r)
}

// loadJson loads the dictionary response from a file
func loadJson() {

	go_path := os.Getenv("GOPATH")
	path := go_path +
		"/src/github.com/PaulioRandall/qlueless-assembly-line-api" +
		"/web/dictionaries.json"

	bytes, err := ioutil.ReadFile(path)
	if shr.Log_if_err(err) {
		reply.Data = nil
		return
	}

	err = json.Unmarshal(bytes, &reply.Data)
	if shr.Log_if_err(err) {
		reply.Data = nil
	}
}
