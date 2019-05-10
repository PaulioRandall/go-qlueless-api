package GET

import (
	"testing"

	v "github.com/PaulioRandall/go-qlueless-assembly-api/cmd/ventures"
	a "github.com/PaulioRandall/go-qlueless-assembly-api/internal/asserts"
	test "github.com/PaulioRandall/go-qlueless-assembly-api/test"
	vtest "github.com/PaulioRandall/go-qlueless-assembly-api/test/ventures"
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
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is a JSON array containing all living Ventures
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)

	out := v.AssertVentureSliceFromReader(t, res.Body)
	exp := vtest.DBQueryAll()
	v.AssertOrderlessSlicesEqual(t, exp, out)
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
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is a JSON object wrapping the with meta information
		And the wrapped meta information contains a message and self link
		And the wrapped data is a JSON array containing all living Ventures
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?wrap",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)

	_, out := v.AssertWrappedVentureSliceFromReader(t, res.Body)
	exp := vtest.DBQueryAll()
	v.AssertOrderlessSlicesEqual(t, exp, out)
}

// ****************************************************************************
// (GET) /ventures?ids={ids}
// ****************************************************************************

func TestGET_Ventures_3(t *testing.T) {

	t.Log(`Given some Ventures already exist on the server
		When a specific living Venture is requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is a JSON array containing only the living Venture requested
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?ids=1",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)

	out := v.AssertVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 1)

	exp := vtest.DBQueryMany("1")
	v.AssertOrderlessSlicesEqual(t, exp, out)
}

func TestGET_Ventures_4(t *testing.T) {

	t.Log(`Given some Ventures already exist on the server
		When a specific living Venture is requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is a JSON array containing only the living Ventures requested
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?ids=1,2,3",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)

	out := v.AssertVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 3)

	exp := vtest.DBQueryMany("1,2,3")
	v.AssertOrderlessSlicesEqual(t, exp, out)
}

func TestGET_Ventures_5(t *testing.T) {

	t.Log(`Given some Ventures already exist on the server
		When non-existent Ventures are requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is an empty JSON array of Ventures
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?ids=888888,999999",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)

	out := v.AssertVentureSliceFromReader(t, res.Body)
	require.Empty(t, out)
}

func TestGET_Ventures_6(t *testing.T) {

	t.Log(`Given some Ventures already exist on the server
	When existent and non-existent Ventures are requested
	Then ensure the response code is 200
	And the 'Content-Type' header contains 'application/json'
	And 'Access-Control-Allow-Origin' is '*'
	And 'Access-Control-Allow-Headers' is '*'
	And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
	And the body is a JSON array containing only the living Ventures requested
	...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?ids=1,88888,2,99999",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)

	out := v.AssertVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 2)

	exp := vtest.DBQueryMany("1,2")
	v.AssertOrderlessSlicesEqual(t, exp, out)
}

// ****************************************************************************
// (GET) /ventures?wrap&ids={ids}
// ****************************************************************************

func TestGET_Ventures_7(t *testing.T) {

	t.Log(`Given some Ventures already exist on the server
		When specific living Ventures are requested
		And the 'wrap' query parameter has been specified
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is a JSON object wrapping the with meta information
		And the wrapped meta information contains a message and self link
		And the body is a JSON array containing only the living Ventures requested
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?wrap&ids=1,4,5",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)

	_, out := v.AssertWrappedVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 3)

	exp := vtest.DBQueryMany("1,4,5")
	v.AssertOrderlessSlicesEqual(t, exp, out)
}
