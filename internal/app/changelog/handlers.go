package changelog

import (
	"io/ioutil"
	"log"
	"net/http"

	h "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/uhttp"
	u "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/utils"
)

const mime_md = "text/markdown; charset=utf-8"

var changelog *[]byte = nil

// ChangelogHandler handles requests for the APIs changelog
func ChangelogHandler(res http.ResponseWriter, req *http.Request) {
	h.LogRequest(req)
	h.AppendCORSHeaders(&res, "GET, HEAD, OPTIONS")

	switch req.Method {
	case "GET":
		_GET_Changelog(&res, req)
	case "HEAD":
		fallthrough
	case "OPTIONS":
		res.Header().Set("Content-Type", mime_md)
		res.WriteHeader(http.StatusOK)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// _GET_Changelog generates responses for obtaining the CHANGELOG
func _GET_Changelog(res *http.ResponseWriter, req *http.Request) {
	if changelog == nil {
		log.Println("[BUG] CHANGELOG not loaded")
		h.WriteServerError(res, req)
		return
	}

	(*res).Header().Set("Content-Type", mime_md)
	(*res).WriteHeader(http.StatusOK)
	(*res).Write(*changelog)
}

// LoadChangelog loads the changelog from a file
func LoadChangelog() {

	path := "./CHANGELOG.md"
	bytes, err := ioutil.ReadFile(path)
	if u.LogIfErr(err) {
		changelog = nil
		return
	}

	changelog = &bytes
}
