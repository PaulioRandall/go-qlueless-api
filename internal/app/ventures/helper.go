package ventures

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	p "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

// InjectDummyVentures injects dummy Ventures so the API testing can performed.
// This function is expected to be removed once a database and formal test data
// has been crafted.
func InjectDummyVentures() {
	ventures.Add(v.Venture{
		Description: "White wizard",
		State:       "Not started",
		Extra:       "colour: white; power: 9000",
		IsAlive:     true,
	})
	ventures.Add(v.Venture{
		Description: "Green lizard",
		State:       "In progress",
		OrderIDs:    "4,5,6,7,8",
		IsAlive:     true,
	})
	ventures.Add(v.Venture{
		Description: "Pink gizzard",
		State:       "Finished",
		OrderIDs:    "1,2,3",
		IsAlive:     false,
	})
	ventures.Add(v.Venture{
		Description: "Eddie Izzard",
		State:       "In Progress",
		OrderIDs:    "4,5,6",
		IsAlive:     true,
	})
	ventures.Add(v.Venture{
		Description: "The Count of Tuscany",
		State:       "In Progress",
		OrderIDs:    "4,5,6",
		IsAlive:     true,
	})
}

// writeSuccessReply writes a success response.
func writeSuccessReply(res *http.ResponseWriter, req *http.Request, code int, data interface{}, msg string) {
	p.AppendJSONHeader(res, "")
	(*res).WriteHeader(code)
	reply := p.PrepResponseData(req, data, msg)
	json.NewEncoder(*res).Encode(reply)
}

// findVenture finds the Venture with the specified ID.
func findVenture(id string, res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	ven, ok := ventures.Get(id)
	if !ok || !ven.IsAlive {
		p.WriteBadRequest(res, req, fmt.Sprintf("Thing '%s' not found", id))
		return v.Venture{}, false
	}
	return ven, true
}

// decodeVenture decodes a Venture from a Request.Body.
func decodeVenture(res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	ven, err := v.DecodeVenture(req.Body)
	if err != nil {
		p.WriteBadRequest(res, req, "Unable to decode request body into a Venture")
		return v.Venture{}, false
	}
	return ven, true
}

// validateNewVenture validates a new Venture that has yet to be assigned an ID.
func validateNewVenture(ven v.Venture, res *http.ResponseWriter, req *http.Request) bool {
	errMsgs := ven.Validate(true)
	if len(errMsgs) != 0 {
		p.WriteBadRequest(res, req, strings.Join(errMsgs, " "))
		return false
	}
	return true
}

// decodeVentureUpdate decodes an update to a Venture from a Request.Body.
func decodeVentureUpdate(res *http.ResponseWriter, req *http.Request) (v.VentureUpdate, bool) {
	vu, err := v.DecodeVentureUpdate(req.Body)
	if err != nil {
		p.WriteBadRequest(res, req,
			"Unable to decode request body into a Venture update")
		return v.VentureUpdate{}, false
	}
	return vu, true
}

// validateVentureUpdate validates a Venture update.
func validateVentureUpdate(vu v.VentureUpdate, res *http.ResponseWriter, req *http.Request) bool {
	errMsgs := vu.Validate()
	if len(errMsgs) != 0 {
		p.WriteBadRequest(res, req, strings.Join(errMsgs, " "))
		return false
	}
	return true
}

// updateVenture updates a Venture in the data store.
func updateVenture(vu v.VentureUpdate, res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	ven, ok := ventures.Update(vu)
	if !ok {
		p.WriteBadRequest(res, req,
			fmt.Sprintf("Venture with ID '%s' could not be found", vu.Values.ID))
		return v.Venture{}, false
	}
	return ven, true
}

// deleteVenture deletes a Venture from the data store.
func deleteVenture(id string, res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	ven, ok := ventures.Delete(id)
	if !ok {
		p.WriteBadRequest(res, req, fmt.Sprintf("Thing '%s' not found", id))
		return v.Venture{}, false
	}
	return ven, true
}
