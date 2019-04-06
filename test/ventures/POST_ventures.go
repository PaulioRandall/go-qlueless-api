package ventures

import (
	"bytes"
	"encoding/json"
	"testing"

	a "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
	test "github.com/PaulioRandall/go-qlueless-assembly-api/test"
	require "github.com/stretchr/testify/require"
)

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

	beginVenTest()
	defer endVenTest()

	input := v.Venture{
		Description: "A new Venture",
		State:       "Not started",
		OrderIDs:    "1,2,3",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := test.APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "POST",
		Body:   buf,
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 201, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	output := v.AssertVentureFromReader(t, res.Body)

	input.ID = output.ID
	input.IsDead = false
	v.AssertGenericVenture(t, output)
	v.AssertVentureModEquals(t, input, output)
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

	beginVenTest()
	defer endVenTest()

	input := v.Venture{
		Description: "",
		State:       "",
		OrderIDs:    "invalid",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := test.APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "POST",
		Body:   buf,
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 400, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	test.AssertWrappedErrorBody(t, res.Body)
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

	beginVenTest()
	defer endVenTest()

	input := v.Venture{
		Description: "A new Venture",
		State:       "Not started",
		OrderIDs:    "1,2,3",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?wrap",
		Method: "POST",
		Body:   buf,
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 201, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	_, output := v.AssertWrappedVentureFromReader(t, res.Body)

	input.ID = output.ID
	input.IsDead = false
	v.AssertGenericVenture(t, output)
	v.AssertVentureModEquals(t, input, output)
}
