package batches

import (
	"net/http"

	shr "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// AllBatchesHandler handles requests for all batches currently within the
// service
func AllBatchesHandler(w http.ResponseWriter, r *http.Request) {
	shr.LogRequest(r)

	batches := LoadBatches()
	if batches == nil {
		shr.Http_500(w)
		return
	}

	shr.WriteJsonReply("Found all batches", batches, w, r)
}
