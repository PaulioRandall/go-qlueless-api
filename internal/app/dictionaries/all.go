// Package internal/app contains non-reusable internal application code
package dictionaries

import (
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// All_dictionaries_handler handles requests for the service dictionaries
func AllDictsHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	if dicts == nil {
		Write500Reply(&res, req)
		return
	}

	data := PrepResponseData(req, dicts, "All service dictionaries")
	WriteReply(&res, req, data)
}
