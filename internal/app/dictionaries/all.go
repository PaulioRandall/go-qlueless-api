// Package internal/app contains non-reusable internal application code
package dictionaries

import (
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// All_dictionaries_handler handles requests for the service dictionaries
func AllDictsHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)
	r := Reply{
		Req: req,
		Res: &res,
	}

	dicts := LoadDicts()
	if dicts == nil {
		Http_500(&r)
		return
	}

	WriteJsonReply(&r, Str("All service dictionaries"), dicts, nil)
}
