package batches

import (
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// AllBatchesHandler handles requests for all batches currently within the
// service
func AllBatchesHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)
	r := Reply{
		Req: req,
		Res: &res,
	}

	b := make([]WorkItem, 0)
	for _, v := range batches {
		b = append(b, v)
	}

	WriteJsonReply(&r, Str("Found all batches"), b, nil)
}
