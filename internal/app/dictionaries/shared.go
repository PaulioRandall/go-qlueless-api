package dictionaries

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	shr "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var dictionaries map[string]interface{}
var dictLoader sync.Once

// LoadDicts loads the service dictionaries
func LoadDicts() map[string]interface{} {
	dictLoader.Do(loadJson)
	return dictionaries
}

// loadJson loads the dictionary response from a file
func loadJson() {

	go_path := os.Getenv("GOPATH")
	path := go_path +
		"/src/github.com/PaulioRandall/go-qlueless-assembly-api" +
		"/web/dictionaries.json"

	bytes, err := ioutil.ReadFile(path)
	if shr.LogIfErr(err) {
		dictionaries = nil
		return
	}

	err = json.Unmarshal(bytes, &dictionaries)
	if shr.LogIfErr(err) {
		dictionaries = nil
	}
}
