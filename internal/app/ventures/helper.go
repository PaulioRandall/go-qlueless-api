package ventures

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	h "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/uhttp"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

// InjectDummyVentures injects dummy Ventures so the API testing can performed.
// This function is expected to be removed once a database and formal test data
// has been crafted.
func InjectDummyVentures() {
	ventures.Add(v.NewVenture{
		Description: "White wizard",
		State:       "Not started",
		Extra:       "colour: white; power: 9000",
	})
	ventures.Add(v.NewVenture{
		Description: "Green lizard",
		State:       "In progress",
		OrderIDs:    "4,5,6,7,8",
	})
	ventures.Add(v.NewVenture{
		Description: "Pink gizzard",
		State:       "Finished",
		OrderIDs:    "1,2,3",
	})
	ventures.Add(v.NewVenture{
		Description: "Eddie Izzard",
		State:       "In Progress",
		OrderIDs:    "4,5,6",
	})
	ventures.Add(v.NewVenture{
		Description: "The Count of Tuscany",
		State:       "In Progress",
		OrderIDs:    "4,5,6",
	})
	ventures.Update(v.VentureUpdate{
		Props: "is_alive",
		Values: v.Venture{
			IsAlive: false,
		},
	})
}

// writeSuccessReply writes a success response.
func writeSuccessReply(res *http.ResponseWriter, req *http.Request, code int, data interface{}, msg string) {
	h.AppendJSONHeader(res, "")
	(*res).WriteHeader(code)
	reply := h.PrepResponseData(req, data, msg)
	json.NewEncoder(*res).Encode(reply)
}

// findVenture finds the Venture with the specified ID.
func findVenture(id string, res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	ven, ok := ventures.Get(id)
	if !ok {
		h.WriteBadRequest(res, req, fmt.Sprintf("Thing '%s' not found", id))
		return v.Venture{}, false
	}
	return ven, true
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
func validateNewVenture(ven v.NewVenture, res *http.ResponseWriter, req *http.Request) bool {
	errMsgs := ven.Validate()
	if len(errMsgs) != 0 {
		h.WriteBadRequest(res, req, strings.Join(errMsgs, " "))
		return false
	}
	return true
}

// decodeVentureUpdate decodes an update to a Venture from a Request.Body.
func decodeVentureUpdate(res *http.ResponseWriter, req *http.Request) (v.VentureUpdate, bool) {
	vu, err := v.DecodeVentureUpdate(req.Body)
	if err != nil {
		h.WriteBadRequest(res, req,
			"Unable to decode request body into a Venture update")
		return v.VentureUpdate{}, false
	}
	return vu, true
}

// validateVentureUpdate validates a Venture update.
func validateVentureUpdate(vu v.VentureUpdate, res *http.ResponseWriter, req *http.Request) bool {
	errMsgs := vu.Validate()
	if len(errMsgs) != 0 {
		h.WriteBadRequest(res, req, strings.Join(errMsgs, " "))
		return false
	}
	return true
}

// updateVenture updates a Venture in the data store.
func updateVenture(vu v.VentureUpdate, res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	ven, ok := ventures.Update(vu)
	if !ok {
		h.WriteBadRequest(res, req,
			fmt.Sprintf("Venture with ID '%s' could not be found", vu.Values.ID))
		return v.Venture{}, false
	}
	return ven, true
}

// deleteVenture deletes a Venture from the data store.
func deleteVenture(id string, res *http.ResponseWriter, req *http.Request) (v.Venture, bool) {
	ven, ok := ventures.Delete(id)
	if !ok {
		h.WriteBadRequest(res, req, fmt.Sprintf("Thing '%s' not found", id))
		return v.Venture{}, false
	}
	return ven, true
}
