package dictionaries

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	shr "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var reply shr.Reply = shr.Reply{
	Message: "All service dictionaries and their entries",
}
var dictLoader sync.Once

// LoadDictsReply loads dictionaries and creates a Reply
func LoadDictsReply() *shr.Reply {
	dictLoader.Do(loadJson)
	if reply.Data == nil {
		return nil
	}
	return &reply
}

// loadJson loads the dictionary response from a file
func loadJson() {

	go_path := os.Getenv("GOPATH")
	path := go_path +
		"/src/github.com/PaulioRandall/go-qlueless-assembly-api" +
		"/web/dictionaries.json"

	bytes, err := ioutil.ReadFile(path)
	if shr.LogIfErr(err) {
		reply.Data = nil
		return
	}

	err = json.Unmarshal(bytes, &reply.Data)
	if shr.LogIfErr(err) {
		reply.Data = nil
	}
}
