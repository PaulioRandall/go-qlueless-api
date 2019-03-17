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

	data := prepData(req, dicts)
	WriteReply(&res, req, data)
}

// prepData prepares the data by wrapping it up if the client has requested
func prepData(req *http.Request, data interface{}) interface{} {
	if WrapUpReply(req) {
		return ReplyWrapped{
			Message: "All service dictionaries",
			Self:    req.URL.String(),
			Data:    data,
		}
	} else {
		return data
	}
}
