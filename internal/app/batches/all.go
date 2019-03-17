package batches

import (
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// AllBatchesHandler handles requests for all batches currently within the
// service
func AllBatchesHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	b := make([]WorkItem, 0)
	for _, v := range batches {
		b = append(b, v)
	}

	data := prepBatchesData(req, b)
	WriteReply(&res, req, data)
}

// prepData prepares the data by wrapping it up if the client has requested
func prepBatchesData(req *http.Request, data interface{}) interface{} {
	if WrapUpReply(req) {
		return ReplyWrapped{
			Message: "Found all batches",
			Self:    req.URL.String(),
			Data:    data,
		}
	} else {
		return data
	}
}
