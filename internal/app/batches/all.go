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

	batches := LoadBatches()
	if batches == nil {
		Http_500(&r)
		return
	}

	WriteJsonReply(&r, Str("Found all batches"), batches, nil)
}
