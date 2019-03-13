package dictionaries

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

var reply shr.Reply = shr.Reply{
	Message: "All service dictionaries and their entries",
}
var dict_loader sync.Once

// Load_dictionaries_reply loads dictionaries and creates a Reply
func Load_dictionaries_reply() *shr.Reply {
	dict_loader.Do(loadJson)
	if reply.Data == nil {
		return nil
	}
	return &reply
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
