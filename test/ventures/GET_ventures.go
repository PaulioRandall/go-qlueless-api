package ventures

import (
	"testing"

	a "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
	test "github.com/PaulioRandall/go-qlueless-assembly-api/test"
	require "github.com/stretchr/testify/require"
)

// ****************************************************************************
// (GET) /ventures
// ****************************************************************************

func TestGET_Ventures_1(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When all Ventures are requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, and OPTIONS
		And the body is a JSON array containing all living Ventures
		...`)

	beginVenTest()
	defer endVenTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	out := v.AssertVentureSliceFromReader(t, res.Body)
	v.AssertVentureSliceModEquals(t, livingVens, out)
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
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, and OPTIONS
		And the body is a JSON object wrapping the with meta information
		And the wrapped meta information contains a message and self link
		And the wrapped data is a JSON array containing all living Ventures
		...`)

	beginVenTest()
	defer endVenTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?wrap",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	_, out := v.AssertWrappedVentureSliceFromReader(t, res.Body)
	v.AssertVentureSliceModEquals(t, livingVens, out)
}

// ****************************************************************************
// (GET) /ventures?id={id}
// ****************************************************************************

func TestGET_Venture_1(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When a specific existing Venture is requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, and OPTIONS
		And the body is a JSON object representing the Venture requested
		...`)

	beginVenTest()
	defer endVenTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?id=1",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	out := v.AssertVentureFromReader(t, res.Body)
	v.AssertVentureModEquals(t, livingVens[out.ID], out)
}

func TestGET_Venture_2(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When a specific non-existent Venture is requested
		Then ensure the response code is 404
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, and OPTIONS
		And the body is a JSON object representing an error response
		...`)

	beginVenTest()
	defer endVenTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?id=999999",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 400, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	test.AssertWrappedErrorBody(t, res.Body)
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
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, and OPTIONS
		And the wrapped meta information contains a message and self link
		And the wrapped data is a JSON object representing the requested Venture
		...`)

	beginVenTest()
	defer endVenTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?wrap&id=1",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	_, out := v.AssertWrappedVentureFromReader(t, res.Body)
	v.AssertVentureModEquals(t, livingVens[out.ID], out)
}
