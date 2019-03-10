package pkg

import "net/http"

func AppendStdHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
