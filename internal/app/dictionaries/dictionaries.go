// Package internal/app contains non-reusable internal application code
package dictionaries

import (
	"net/http"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

// All_dictionaries_handler handles requests for the service dictionaries
func All_dictionaries_handler(w http.ResponseWriter, r *http.Request) {
	shr.Log_request(r)

	reply := Load_dictionaries_reply()
	if reply == nil {
		shr.Http_500(&w)
		return
	}

	shr.WriteJsonReply(*reply, w, r)
}
