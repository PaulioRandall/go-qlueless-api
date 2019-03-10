package pkg

import "net/http"

type Reply struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type WorkItem struct {
	Title               string `json:"title"`
	Description         string `json:"description"`
	Work_item_id        int    `json:"work_item_id"`
	Parent_work_item_id int    `json:"parent_work_item_id,omitempty"`
	Tag_id              string `json:"tag_id"`
	Status_id           string `json:"status_id"`
	Additional          string `json:"additional,omitempty"`
}

func AppendJSONHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
