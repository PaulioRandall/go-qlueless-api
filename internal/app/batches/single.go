package batches

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// SingleBatchHandler handles requests for a specific batches currently
// within the service
func SingleBatchHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)
	r := Reply{
		Req: req,
		Res: &res,
	}

	batches := LoadBatches()
	if batches == nil {
		Http_500(&r)
		return
	}

	id := mux.Vars(req)["batch_id"]
	b, ok := batches[id]

	if !ok {
		r.Message = Str(fmt.Sprintf("Batch %v not found", id))
		Http_4xx(&r, 404)
		return
	}

	m := Str(fmt.Sprintf("Found batch %v", id))
	WriteJsonReply(&r, m, b, nil)
}
