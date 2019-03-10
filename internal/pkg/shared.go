package pkg

import "net/http"

type Reply struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func AppendJSONHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
