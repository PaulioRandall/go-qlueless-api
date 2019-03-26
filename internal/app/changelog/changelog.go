package changelog

import (
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

const mime_md = "text/markdown; charset=utf-8"

var changelog *[]byte = nil

// ChangelogHandler handles requests for the APIs changelog
func ChangelogHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	switch req.Method {
	case "GET":
		get_Changelog(&res, req)
	case "HEAD":
		WriteEmptyReply(&res, mime_md)
	case "OPTIONS":
		WriteEmptyReply(&res, mime_md)
	default:
		MethodNotAllowed(&res, req)
	}
}

// get_Changelog generates responses for obtaining the CHANGELOG
func get_Changelog(res *http.ResponseWriter, req *http.Request) {
	if changelog == nil {
		log.Println("[BUG] CHANGELOG not loaded")
		Write500Reply(res, req)
		return
	}

	WriteReply(res, changelog, mime_md)
}

// LoadChangelog loads the changelog from a file
func LoadChangelog() {

	path := "./CHANGELOG.md"
	bytes, err := ioutil.ReadFile(path)
	if LogIfErr(err) {
		changelog = nil
		return
	}

	changelog = &bytes
}
