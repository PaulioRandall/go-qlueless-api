package pkg

import "net/http"

func AppendJSONHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
