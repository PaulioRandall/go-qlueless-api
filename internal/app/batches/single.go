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

	id := mux.Vars(req)["batch_id"]
	b, ok := batches[id]

	if !ok {
		r := Reply4XX{
			Res:     &res,
			Req:     req,
			Message: fmt.Sprintf("Batch %v not found", id),
		}
		Write4XXReply(404, &r)
		return
	}

	data := PrepResponseData(req, b, "Found batch")
	WriteReply(&res, req, data)
}
