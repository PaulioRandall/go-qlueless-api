package batches

import (
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// AllBatchesHandler handles requests for all batches currently within the
// service
func AllBatchesHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	b := make([]Thing, 0)
	for _, v := range batches {
		b = append(b, v)
	}

	data := PrepResponseData(req, b, "Found all batches")
	WriteReply(&res, req, data)
}
