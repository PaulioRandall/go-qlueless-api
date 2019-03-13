package batches

import (
	"net/http"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

// All_batches_handler handles requests for all batches currently within the
// service
func All_batches_handler(w http.ResponseWriter, r *http.Request) {
	shr.Log_request(r)

	batches := Load_batches()
	if batches == nil {
		shr.Http_500(&w)
		return
	}

	reply := shr.Reply{
		Message: "Found all batches",
		Data:    batches,
	}

	shr.WriteJsonReply(reply, w, r)
}
