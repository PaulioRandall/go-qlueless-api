package batches

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	shr "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// SingleBatchHandler handles requests for a specific batches currently
// within the service
func SingleBatchHandler(w http.ResponseWriter, r *http.Request) {
	shr.LogRequest(r)

	batches := LoadBatches()
	if batches == nil {
		shr.Http_500(w)
		return
	}

	id := mux.Vars(r)["batch_id"]
	var b *shr.WorkItem = shr.FindWorkItem(batches, id)

	if b == nil {
		shr.Http_4xx(w, 404, fmt.Sprintf("Batch %v not found", id), "")
		return
	}

	shr.WriteJsonReply(fmt.Sprintf("Found batch %v", id), b, "", w, r)
}
