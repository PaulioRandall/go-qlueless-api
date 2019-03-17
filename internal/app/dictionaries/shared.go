package dictionaries

import (
	"encoding/json"
	"io/ioutil"
	"os"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var dicts map[string]interface{} = nil

// LoadDicts loads the dictionaries data
func LoadDicts() {

	go_path := os.Getenv("GOPATH")
	path := go_path +
		"/src/github.com/PaulioRandall/go-qlueless-assembly-api" +
		"/web/dictionaries.json"

	bytes, err := ioutil.ReadFile(path)
	if LogIfErr(err) {
		dicts = nil
		return
	}

	err = json.Unmarshal(bytes, &dicts)
	if LogIfErr(err) {
		dicts = nil
	}
}
