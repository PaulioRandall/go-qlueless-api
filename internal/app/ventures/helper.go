package ventures

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	q "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/qserver"
	h "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/uhttp"
	u "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/utils"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

// writeSuccessReply writes a success response.
func writeSuccessReply(res *http.ResponseWriter, req *http.Request, code int, data interface{}, msg string) {
	h.AppendJSONHeader(res, "")
	(*res).WriteHeader(code)
	reply := h.PrepResponseData(req, data, msg)
	json.NewEncoder(*res).Encode(reply)
}

// findVentures finds the Ventures with the IDs specified.
func findVentures(ids string, res *http.ResponseWriter, req *http.Request) ([]v.Venture, bool) {
	idSlice := strings.Split(ids, ",")
	s := make([]interface{}, len(idSlice))

	for i, id := range idSlice {
		s[i] = id
	}

	vens, err := v.QueryMany(q.Sev.DB, s)

	if err != nil {
		h.WriteServerError(res, req)
		return nil, false
	}

	return vens, true
}

// decodeNewVenture decodes a NewVenture from a Request.Body.
func decodeNewVenture(res *http.ResponseWriter, req *http.Request) (v.NewVenture, bool) {
	ven, err := v.DecodeNewVenture(req.Body)
	if err != nil {
		h.WriteBadRequest(res, req, "Unable to decode request body into a Venture")
		return v.NewVenture{}, false
	}
	return ven, true
}

// validateNewVenture validates a NewVenture that has yet to be assigned an ID.
func validateNewVenture(ven *v.NewVenture, res *http.ResponseWriter, req *http.Request) bool {
	errMsgs := ven.Validate()
	if len(errMsgs) != 0 {
		h.WriteBadRequest(res, req, strings.Join(errMsgs, " "))
		return false
	}
	return true
}

// insertNewVenture inserts a new Venture into the database
func insertNewVenture(new *v.NewVenture, res *http.ResponseWriter, req *http.Request) (*v.Venture, bool) {
	ven, ok := new.Insert(q.Sev.DB)
	if !ok {
		h.WriteServerError(res, req)
	}
	return ven, ok
}

// decodeModVentures decodes modifications to Ventures from a Request.Body.
func decodeModVentures(res *http.ResponseWriter, req *http.Request) (*v.ModVenture, bool) {
	mv, err := v.DecodeModVenture(req.Body)
	if err != nil {
		h.WriteBadRequest(res, req,
			"Unable to decode request body into a Venture update")
		return nil, false
	}
	return &mv, true
}

// validateModVentures validates a Venture update.
func validateModVentures(mv *v.ModVenture, res *http.ResponseWriter, req *http.Request) bool {
	errMsgs := mv.Validate()
	if len(errMsgs) != 0 {
		h.WriteBadRequest(res, req, strings.Join(errMsgs, " "))
		return false
	}
	return true
}

// ventureIdCsvToSlice validates then parses a CSV string of IDs into a slice.
func ventureIdCsvToSlice(idCsv string, res *http.ResponseWriter, req *http.Request) ([]string, bool) {
	idCsv = u.StripWhitespace(idCsv)

	if idCsv == "" {
		h.WriteBadRequest(res, req, "Query parameter 'ids' is missing or empty")
		return nil, false
	}

	if !u.IsPositiveIntCSV(idCsv) {
		h.WriteBadRequest(res, req, fmt.Sprintf("Could not parse query parameter"+
			" 'ids=%s' into a list of Venture IDs", idCsv))
		return nil, false
	}

	ids := strings.Split(idCsv, ",")
	return ids, true
}

// ventureIDsToCSV returns a CSV string of all Venture IDs within the given
// slice.
func ventureIDsToCSV(vens []v.Venture) string {
	ids := ""
	for i, ven := range vens {
		switch i {
		case 0:
			ids = ven.ID
		default:
			ids += ", " + ven.ID
		}
	}
	return ids
}

// pushModifiedVentures performs the specified update operation.
func pushModifiedVentures(mv *v.ModVenture, res *http.ResponseWriter, req *http.Request) ([]v.Venture, bool) {
	vens, ok := mv.Update(q.Sev.DB)
	if !ok {
		h.WriteServerError(res, req)
		return nil, false
	}
	return vens, true
}
