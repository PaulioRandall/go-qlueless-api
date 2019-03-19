package things

import (
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// AllThingsHandler handles requests for all things currently within the
// service
func AllThingsHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	o := make([]Thing, 0)
	for _, v := range things {
		if !v.IsDead {
			o = append(o, v)
		}
	}

	data := PrepResponseData(req, o, "Found all things")
	WriteReply(&res, req, data)
}
