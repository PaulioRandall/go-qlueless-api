package batches

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

// Single_batch_handler handles requests for a specific batches currently
// within the service
func Single_batch_handler(w http.ResponseWriter, r *http.Request) {
	shr.Log_request(r)

	batches := Load_batches()
	if batches == nil {
		shr.Http_500(&w)
		return
	}

	id := mux.Vars(r)["batch_id"]
	var batch *shr.WorkItem = shr.FindWorkItem(batches, id)

	if batch == nil {
		shr.Http_4xx(&w, 404, fmt.Sprintf("Batch %v not found", id))
		return
	}

	reply := shr.Reply{
		Message: fmt.Sprintf("Found batch %v", id),
		Data:    batch,
	}

	shr.WriteJsonReply(reply, w, r)
}
