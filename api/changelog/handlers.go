package changelog

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	cookies "github.com/PaulioRandall/go-cookies/cookies"
	uhttp "github.com/PaulioRandall/go-cookies/uhttp"
	writers "github.com/PaulioRandall/go-qlueless-api/shared/writers"
)

const mime_md = "text/markdown; charset=utf-8"

var changelog *[]byte = nil
var once sync.Once

var cors uhttp.CorsHeaders = uhttp.CorsHeaders{
	Origin:  "*",
	Headers: "*",
	Methods: "GET, OPTIONS",
}

// ChangelogHandler handles requests for the APIs changelog
func ChangelogHandler(res http.ResponseWriter, req *http.Request) {
	uhttp.LogRequest(req)
	uhttp.UseCors(&res, &cors)

	switch req.Method {
	case "GET":
		get(&res, req)
	case "OPTIONS":
		res.WriteHeader(http.StatusOK)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// get generates responses for obtaining the CHANGELOG
func get(res *http.ResponseWriter, req *http.Request) {
	once.Do(load)

	if changelog == nil {
		log.Println("[BUG] CHANGELOG not loaded")
		writers.WriteServerError(res, req)
		return
	}

	(*res).Header().Set("Content-Type", mime_md)
	(*res).WriteHeader(http.StatusOK)
	(*res).Write(*changelog)
}

// load loads the changelog from a file
func load() {
	bytes, err := ioutil.ReadFile("./CHANGELOG.md")
	if cookies.LogIfErr(err) {
		changelog = nil
		return
	}

	changelog = &bytes
	return
}
