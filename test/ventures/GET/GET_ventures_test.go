package GET

import (
	"testing"

	"github.com/PaulioRandall/go-cookies/toastify"
	"github.com/PaulioRandall/go-qlueless-api/api/ventures"
	"github.com/PaulioRandall/go-qlueless-api/test"
	vtest "github.com/PaulioRandall/go-qlueless-api/test/ventures"
	"github.com/stretchr/testify/require"
)

func init() {
	test.SetWorkingDir("../../../bin")
}

// ****************************************************************************
// (GET) /ventures
// ****************************************************************************

func TestGET_Ventures_1(t *testing.T) {

	test.PrintTestDescription(t, `
		Given some Ventures already exist on the server
		When all Ventures are requested
		Ensure the response code is 200
		And header includes:
			Content-Type:                   'application/json; charset=utf-8'
			Access-Control-Allow-Origin:    '*'
			Access-Control-Allow-Headers:   '*'
			Access-Control-Allow-Methods:   'GET, POST, PUT, DELETE, OPTIONS'
		And the body is a JSON array containing all injected Ventures
	`)

	vtest.SetupEmptyTest()
	defer vtest.TearDown()

	injected := vtest.InjectAll([]ventures.NewVenture{
		ventures.NewVenture{
			Description: "White wizard",
			State:       "Not started",
		},
		ventures.NewVenture{
			Description: "Green lizard",
			State:       "In progress",
		},
		ventures.NewVenture{
			Description: "Pink gizzard",
			State:       "Finished",
		},
	})

	req := test.APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "GET",
	}
	res := req.Fire()
	defer res.Body.Close()

	require.Equal(t, 200, res.StatusCode)
	toastify.AssertHeaderEqual(t, "Content-Type", res.Header, "application/json; charset=utf-8")
	toastify.AssertHeaderEqual(t, "Access-Control-Allow-Origin", res.Header, "*")
	toastify.AssertHeaderEqual(t, "Access-Control-Allow-Headers", res.Header, "*")
	toastify.AssertHeaderEqual(t, "Access-Control-Allow-Methods", res.Header, "GET, POST, PUT, DELETE, OPTIONS")

	body := test.PrintBody(t, res)

	result := ventures.RequireSliceOfVentures(t, body)
	stored, err := ventures.QueryAll()
	require.Nil(t, err)

	ventures.AssertVenturesEqual(t, injected, result, true)
	ventures.AssertVenturesEqual(t, stored, result, true)
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
		And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
		And the body is a JSON object wrapping the with meta information
		And the wrapped meta information contains a message and self link
		And the wrapped data is a JSON array containing all living Ventures
		...`)

	vtest.SetupTest()
	defer vtest.TearDown()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?wrap",
		Method: "GET",
	}
	res := req.Fire()

	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", "GET, POST, PUT, DELETE, OPTIONS")

	_, out := ventures.AssertWrappedVentureSliceFromReader(t, res.Body)
	exp := vtest.DBQueryAll()
	ventures.AssertOrderlessSlicesEqual(t, exp, out)
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
		And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
		And the body is a JSON array containing only the living Venture requested
		...`)

	vtest.SetupTest()
	defer vtest.TearDown()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?ids=1",
		Method: "GET",
	}
	res := req.Fire()

	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", "GET, POST, PUT, DELETE, OPTIONS")

	out := ventures.AssertVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 1)

	exp := vtest.DBQueryMany("1")
	ventures.AssertOrderlessSlicesEqual(t, exp, out)
}

func TestGET_Ventures_4(t *testing.T) {

	t.Log(`Given some Ventures already exist on the server
		When a specific living Venture is requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
		And the body is a JSON array containing only the living Ventures requested
		...`)

	vtest.SetupTest()
	defer vtest.TearDown()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?ids=1,2,3",
		Method: "GET",
	}
	res := req.Fire()

	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", "GET, POST, PUT, DELETE, OPTIONS")

	out := ventures.AssertVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 3)

	exp := vtest.DBQueryMany("1,2,3")
	ventures.AssertOrderlessSlicesEqual(t, exp, out)
}

func TestGET_Ventures_5(t *testing.T) {

	t.Log(`Given some Ventures already exist on the server
		When non-existent Ventures are requested
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
		And the body is an empty JSON array of Ventures
		...`)

	vtest.SetupTest()
	defer vtest.TearDown()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?ids=888888,999999",
		Method: "GET",
	}
	res := req.Fire()

	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", "GET, POST, PUT, DELETE, OPTIONS")

	out := ventures.AssertVentureSliceFromReader(t, res.Body)
	require.Empty(t, out)
}

func TestGET_Ventures_6(t *testing.T) {

	t.Log(`Given some Ventures already exist on the server
	When existent and non-existent Ventures are requested
	Then ensure the response code is 200
	And the 'Content-Type' header contains 'application/json'
	And 'Access-Control-Allow-Origin' is '*'
	And 'Access-Control-Allow-Headers' is '*'
	And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
	And the body is a JSON array containing only the living Ventures requested
	...`)

	vtest.SetupTest()
	defer vtest.TearDown()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?ids=1,88888,2,99999",
		Method: "GET",
	}
	res := req.Fire()

	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", "GET, POST, PUT, DELETE, OPTIONS")

	out := ventures.AssertVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 2)

	exp := vtest.DBQueryMany("1,2")
	ventures.AssertOrderlessSlicesEqual(t, exp, out)
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
		And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
		And the body is a JSON object wrapping the with meta information
		And the wrapped meta information contains a message and self link
		And the body is a JSON array containing only the living Ventures requested
		...`)

	vtest.SetupTest()
	defer vtest.TearDown()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?wrap&ids=1,4,5",
		Method: "GET",
	}
	res := req.Fire()

	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", "GET, POST, PUT, DELETE, OPTIONS")

	_, out := ventures.AssertWrappedVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 3)

	exp := vtest.DBQueryMany("1,4,5")
	ventures.AssertOrderlessSlicesEqual(t, exp, out)
}
