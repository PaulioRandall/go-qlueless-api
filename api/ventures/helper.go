package ventures

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PaulioRandall/go-cookies/cookies"
	"github.com/PaulioRandall/go-qlueless-api/shared/writers"
)

// find finds the Ventures with the specified IDs.
func find(ids string, res *http.ResponseWriter, req *http.Request) ([]Venture, bool) {
	idSlice := strings.Split(ids, ",")
	s := make([]interface{}, len(idSlice))

	for i, id := range idSlice {
		s[i] = id
	}

	vens, err := QueryMany(s)

	if err != nil {
		writers.WriteServerError(res, req)
		return nil, false
	}

	return vens, true
}

// decodeNew decodes a NewVenture from a Request.Body.
func decodeNew(res *http.ResponseWriter, req *http.Request) (NewVenture, bool) {
	ven, err := DecodeNewVenture(req.Body)
	if err != nil {
		writers.WriteBadRequest(res, req, "Unable to decode request body into a Venture")
		return NewVenture{}, false
	}
	return ven, true
}

// validateNew validates a NewVenture that has yet to be assigned an ID.
func validateNew(ven *NewVenture, res *http.ResponseWriter, req *http.Request) bool {
	errMsgs := ven.Validate()
	if len(errMsgs) != 0 {
		writers.WriteBadRequest(res, req, strings.Join(errMsgs, " "))
		return false
	}
	return true
}

// insertNew inserts a new Venture into the database.
func insertNew(new *NewVenture, res *http.ResponseWriter, req *http.Request) (*Venture, bool) {
	ven, ok := new.Insert()
	if !ok {
		writers.WriteServerError(res, req)
	}
	return ven, ok
}

// decodeMod decodes modifications to Ventures from a Request.Body.
func decodeMod(res *http.ResponseWriter, req *http.Request) (*ModVenture, bool) {
	mv, err := DecodeModVenture(req.Body)
	if err != nil {
		writers.WriteBadRequest(res, req,
			"Unable to decode request body into a Venture update")
		return nil, false
	}
	return &mv, true
}

// validateMod validates a Venture update.
func validateMod(mv *ModVenture, res *http.ResponseWriter, req *http.Request) bool {
	errMsgs := mv.Validate()
	if len(errMsgs) != 0 {
		writers.WriteBadRequest(res, req, strings.Join(errMsgs, " "))
		return false
	}
	return true
}

// idCsvToSlice validates then parses a CSV string of IDs into a slice.
func idCsvToSlice(idCsv string, res *http.ResponseWriter, req *http.Request) ([]string, bool) {
	idCsv = cookies.StripWhitespace(idCsv)

	if idCsv == "" {
		writers.WriteBadRequest(res, req, "Query parameter 'ids' is missing or empty")
		return nil, false
	}

	if !cookies.IsUintCSV(idCsv) {
		writers.WriteBadRequest(res, req, fmt.Sprintf("Could not parse query parameter"+
			" 'ids=%s' into a list of Venture IDs", idCsv))
		return nil, false
	}

	ids := strings.Split(idCsv, ",")
	return ids, true
}

// idsToCSV returns a CSV string of all Venture IDs within the given slice.
func idsToCSV(vens []Venture) string {
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

// pushMod performs the specified modification operation and pushes the result
// to the database.
func pushMod(mv *ModVenture, res *http.ResponseWriter, req *http.Request) ([]Venture, bool) {
	vens, ok := mv.Update()
	if !ok {
		writers.WriteServerError(res, req)
		return nil, false
	}
	return vens, true
}
