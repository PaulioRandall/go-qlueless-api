package batches

import (
	"net/http"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
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

	reply := shr.Reply{
		Message: "Found all batches",
		Data:    batches,
	}

	shr.WriteJsonReply(reply, w, r)
}
