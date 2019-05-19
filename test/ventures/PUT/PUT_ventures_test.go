package PUT

import (
	"bytes"
	"encoding/json"
	"testing"

	ventures "github.com/PaulioRandall/go-qlueless-api/api/ventures"
	test "github.com/PaulioRandall/go-qlueless-api/test"
	vtest "github.com/PaulioRandall/go-qlueless-api/test/ventures"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// ****************************************************************************
// (PUT) /ventures
// ****************************************************************************

func TestPUT_Ventures_1_OLD(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When an existing Venture is modified and PUT to the server
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
		And the body is a JSON array containing all updated Ventures
		And those Ventures will have new 'last_updated' datetimes
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := ventures.ModVenture{
		IDs:   "1",
		Props: "description, orders, extra",
		Values: ventures.Venture{
			Description: "Black blizzard",
			Orders:      "1,2,3",
			Extra:       "colour: black; power: 9000",
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := test.APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "PUT",
		Body:   buf,
	}
	res := req.Fire()

	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", "GET, POST, PUT, DELETE, OPTIONS")

	out := ventures.AssertVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 1)
	ventures.AssertGenericVenture(t, out[0])

	input.Values.State = out[0].State
	input.Values.ID = out[0].ID
	input.Values.LastModified = out[0].LastModified
	fromDB := vtest.DBQueryOne(out[0].ID)

	assert.Equal(t, input.Values, out[0])
	ventures.AssertVentureModEquals(t, fromDB, out[0])
}

func TestPUT_Ventures_2_OLD(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When an non-existent Venture is PUT to the server
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
		And the body is an empty JSON array
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := ventures.ModVenture{
		IDs:   "999999",
		Props: "description, orders, extra",
		Values: ventures.Venture{
			Description: "Black blizzard",
			Orders:      "1,2,3",
			Extra:       "colour: black; power: 9000",
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := test.APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "PUT",
		Body:   buf,
	}
	res := req.Fire()

	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", "GET, POST, PUT, DELETE, OPTIONS")

	out := ventures.AssertVentureSliceFromReader(t, res.Body)
	require.Empty(t, out)
}

func TestPUT_Ventures_3_OLD(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When a venture modification without IDs is PUT to the server
		Then ensure the response code is 400
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
		And the body is a JSON object representing an error response
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := ventures.ModVenture{
		Props: "description, orders, extra",
		Values: ventures.Venture{
			Description: "Black blizzard",
			Orders:      "1,2,3",
			Extra:       "colour: black; power: 9000",
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := test.APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "PUT",
		Body:   buf,
	}
	res := req.Fire()

	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 400, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", "GET, POST, PUT, DELETE, OPTIONS")
	test.AssertErrorBody(t, res.Body)
}

func TestPUT_Ventures_4_OLD(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When ventures updates are PUT to the server with invalid modifications
		Then ensure the response code is 400
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
		And the body is a JSON object representing an error response
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := ventures.ModVenture{
		IDs:    "1",
		Props:  "description, orders, extra",
		Values: ventures.Venture{},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := test.APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "PUT",
		Body:   buf,
	}
	res := req.Fire()

	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 400, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", "GET, POST, PUT, DELETE, OPTIONS")
	test.AssertErrorBody(t, res.Body)
}

func TestPUT_Ventures_5_OLD(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When multiple existing Ventures are modified as dead and PUT to the server
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
		And the body is a JSON array containing all updated Ventures
		And those Ventures will have new 'last_updated' datetimes
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := ventures.ModVenture{
		IDs:   "4,5",
		Props: "dead",
		Values: ventures.Venture{
			Dead: true,
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := test.APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "PUT",
		Body:   buf,
	}
	res := req.Fire()

	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", "GET, POST, PUT, DELETE, OPTIONS")

	out := ventures.AssertVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 2)

	fromDB := vtest.DBQueryMany("4,5")
	assert.Empty(t, fromDB)
}

// ****************************************************************************
// (PUT) /ventures?wrap
// ****************************************************************************

func TestPUT_Ventures_6_OLD(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When an existing Venture is modified and PUT to the server
		Then ensure the response code is 200
		And the 'wrap' query parameter has been specified
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
		And the body is a JSON object representing a WrappedReply
		And the wrapped data is a JSON array containing all updated Ventures
		And those Ventures will have new 'last_updated' datetimes
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := ventures.ModVenture{
		IDs:   "1",
		Props: "description, state, extra",
		Values: ventures.Venture{
			Description: "Black blizzard",
			State:       "In progress",
			Extra:       "colour: black; power: 9000",
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&input)

	req := test.APICall{
		URL:    "http://localhost:8080/ventures?wrap",
		Method: "PUT",
		Body:   buf,
	}
	res := req.Fire()

	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", "GET, POST, PUT, DELETE, OPTIONS")

	_, out := ventures.AssertWrappedVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 1)
	ventures.AssertGenericVenture(t, out[0])

	input.Values.Orders = out[0].Orders
	input.Values.ID = out[0].ID
	input.Values.LastModified = out[0].LastModified
	assert.Equal(t, input.Values, out[0])

	fromDB := vtest.DBQueryOne(out[0].ID)
	ventures.AssertVentureModEquals(t, fromDB, out[0])
}
