package test

import (
	"bytes"
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"

	a "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
)

var ventureHttpMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

// ****************************************************************************
// (GET) /ventures
// ****************************************************************************

// TODO: Craft some test data and pre-inject it into a SQLite database
func TestGET_Ventures_1(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When all Ventures are requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON array of valid Ventures
		...`)

	req := APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "GET",
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	v.AssertVentureSliceFromReader(t, res.Body)
}

// ****************************************************************************
// (GET) /ventures?wrap
// ****************************************************************************

func TestGET_Ventures_2(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When all Ventures are requested
		And the 'wrap' query parameter has been specified
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON array of valid Ventures wrapped with meta information
		...`)

	req := APICall{
		URL:    "http://localhost:8080/ventures?wrap",
		Method: "GET",
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	v.AssertWrappedVentureSliceFromReader(t, res.Body)
}

// ****************************************************************************
// (GET) /ventures?id={id}
// ****************************************************************************

// TODO: Craft some test data and pre-inject it into a SQLite database
func TestGET_Venture_1(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When a specific existing Venture is requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing a valid Venture
		...`)

	req := APICall{
		URL:    "http://localhost:8080/ventures?id=1",
		Method: "GET",
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	v.AssertVentureFromReader(t, res.Body)
}

func TestGET_Venture_2(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When a specific non-existent Venture is requested
		Then ensure the response code is 404
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing an error response
		...`)

	req := APICall{
		URL:    "http://localhost:8080/ventures?id=999999",
		Method: "GET",
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 400, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	assertWrappedErrorBody(t, res.Body)
}

// ****************************************************************************
// (GET) /ventures?wrap&id={id}
// ****************************************************************************

func TestGET_Venture_3(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When a specific existing Venture is requested
		And the 'wrap' query parameter has been specified
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing a valid Venture wrapped with meta information
		...`)

	req := APICall{
		URL:    "http://localhost:8080/ventures?wrap&id=1",
		Method: "GET",
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	v.AssertWrappedVentureFromReader(t, res.Body)
}

// ****************************************************************************
// (POST) /ventures
// ****************************************************************************

func TestPOST_Venture_1(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When a new valid Venture is POSTed
		Then ensure the response code is 201
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing the living input Venture with a new assigned ID
		...`)

	input := v.Venture{
		Description: "A new Venture",
		State:       "Not started",
		OrderIDs:    "1,2,3",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "POST",
		Body:   buf,
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 201, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	output := v.AssertVentureFromReader(t, res.Body)

	input.ID = output.ID
	input.IsAlive = true
	assert.Equal(t, input, output)
}

func TestPOST_Venture_2(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When a new invalid Venture is POSTed
		Then ensure the response code is 400
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing an error response
		...`)

	input := v.Venture{
		Description: "",
		State:       "",
		OrderIDs:    "invalid",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "POST",
		Body:   buf,
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 400, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	assertWrappedErrorBody(t, res.Body)
}

// ****************************************************************************
// (POST) /ventures?wrap
// ****************************************************************************

func TestPOST_Venture_3(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When a new valid Venture is POSTed
		And the 'wrap' query parameter has been specified
		Then ensure the response code is 201
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing a WrappedReply
		And that the wrapped data is the living input Venture with a new assigned ID
		...`)

	input := v.Venture{
		Description: "A new Venture",
		State:       "Not started",
		OrderIDs:    "1,2,3",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := APICall{
		URL:    "http://localhost:8080/ventures?wrap",
		Method: "POST",
		Body:   buf,
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 201, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	_, output := v.AssertWrappedVentureFromReader(t, res.Body)

	input.ID = output.ID
	input.IsAlive = true
	assert.Equal(t, input, output)
}

// ****************************************************************************
// (PUT) /ventures
// ****************************************************************************

func TestPUT_Venture_1(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When an existing Venture is modified and PUT to the server
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing the updated input Venture
		...`)

	input := v.VentureUpdate{
		Props: "description, state, order_ids, extra",
		Values: v.Venture{
			ID:          "1",
			Description: "Black blizzard",
			State:       "In progress",
			OrderIDs:    "1,2,3",
			Extra:       "colour: black; power: 9000",
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "PUT",
		Body:   buf,
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	output := v.AssertVentureFromReader(t, res.Body)

	input.Values.IsAlive = true
	assert.Equal(t, input.Values, output)
}

func TestPUT_Venture_2(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When an non-existent Venture is PUT to the server
		Then ensure the response code is 400
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing an error response
		...`)

	input := v.VentureUpdate{
		Props: "description, state, order_ids, extra",
		Values: v.Venture{
			ID:          "999999",
			Description: "Black blizzard",
			State:       "In progress",
			OrderIDs:    "1,2,3",
			Extra:       "colour: black; power: 9000",
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "PUT",
		Body:   buf,
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 400, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	assertWrappedErrorBody(t, res.Body)
}

func TestPUT_Venture_3(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When an Venture without an ID is PUT to the server
		Then ensure the response code is 400
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing an error response
		...`)

	input := v.VentureUpdate{
		Props: "description, state, order_ids, extra",
		Values: v.Venture{
			Description: "Black blizzard",
			State:       "In progress",
			OrderIDs:    "1,2,3",
			Extra:       "colour: black; power: 9000",
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "PUT",
		Body:   buf,
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 400, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	assertWrappedErrorBody(t, res.Body)
}

func TestPUT_Venture_4(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When an existent Venture is updated with invalid content and PUT to the server
		Then ensure the response code is 400
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing an error response
		...`)

	input := v.VentureUpdate{
		Props: "description, state, order_ids, extra",
		Values: v.Venture{
			ID: "1",
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "PUT",
		Body:   buf,
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 400, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	assertWrappedErrorBody(t, res.Body)
}

// ****************************************************************************
// (PUT) /ventures?wrap
// ****************************************************************************

func TestPUT_Venture_5(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When an existing Venture is modified and PUT to the server
		Then ensure the response code is 200
		And the 'wrap' query parameter has been specified
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing a WrappedReply
		And the wrapped data is the updated input Venture
		...`)

	input := v.VentureUpdate{
		Props: "description, state, order_ids, extra",
		Values: v.Venture{
			ID:          "1",
			Description: "Black blizzard",
			State:       "In progress",
			OrderIDs:    "1,2,3",
			Extra:       "colour: black; power: 9000",
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := APICall{
		URL:    "http://localhost:8080/ventures?wrap",
		Method: "PUT",
		Body:   buf,
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	_, output := v.AssertWrappedVentureFromReader(t, res.Body)

	input.Values.IsAlive = true
	assert.Equal(t, input.Values, output)
}

// ****************************************************************************
// (DELETE) /ventures?id={id}
// ****************************************************************************

func TestDELETE_Venture_1(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When a DELETE Venture requested is made for an existent Venture
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing the deleted Venture
		...`)

	req := APICall{
		URL:    "http://localhost:8080/ventures?id=4",
		Method: "DELETE",
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	output := v.AssertVentureFromReader(t, res.Body)
	assert.Equal(t, "4", output.ID)
}

func TestDELETE_Venture_2(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When a DELETE Venture requested is made for a non-existent Venture
		Then ensure the response code is 400
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing an error response
		...`)

	req := APICall{
		URL:    "http://localhost:8080/ventures?id=99999",
		Method: "DELETE",
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 400, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	assertWrappedErrorBody(t, res.Body)
}

// ****************************************************************************
// (DELETE) /ventures?wrap&id={id}
// ****************************************************************************

func TestDELETE_Venture_3(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When a DELETE Venture requested is made for an existent Venture
		And the 'wrap' query parameter has been specified
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And the body is a JSON object representing a WrappedReply
		And the wrapped data is the deleted Venture
		...`)

	req := APICall{
		URL:    "http://localhost:8080/ventures?wrap&id=5",
		Method: "DELETE",
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	_, output := v.AssertWrappedVentureFromReader(t, res.Body)
	assert.Equal(t, "5", output.ID)
}

// ****************************************************************************
// (OPTIONS) /ventures
// ****************************************************************************

func TestOPTIONS_Ventures(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When /ventures OPTIONS are requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And there is NO response body
		...`)

	req := APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "OPTIONS",
	}
	res := req.fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	assertNoContentHeaders(t, res, ventureHttpMethods)
	assertEmptyBody(t, res.Body)
}

// ****************************************************************************
// (?) /ventures
// ****************************************************************************

func TestINVALID_Ventures(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
	 	When /ventures is called using invalid methods
		Then ensure the response code is 405
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE, and OPTIONS
		And there is NO response body
		...`)

	verifyNotAllowedMethods(t, "http://localhost:8080/ventures", ventureHttpMethods)
}
