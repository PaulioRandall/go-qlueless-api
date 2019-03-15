// Package internal/app contains non-reusable internal application code
package dictionaries

import (
	"net/http"

	shr "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// All_dictionaries_handler handles requests for the service dictionaries
func AllDictsHandler(w http.ResponseWriter, r *http.Request) {
	shr.LogRequest(r)

	reply := LoadDictsReply()
	if reply == nil {
		shr.Http_500(w)
		return
	}

	shr.WriteJsonReply(*reply, w, r)
}
