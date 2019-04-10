package ventures

import (
	"bytes"
	"encoding/json"
	"testing"

	a "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
	test "github.com/PaulioRandall/go-qlueless-assembly-api/test"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// ****************************************************************************
// (PUT) /ventures
// ****************************************************************************

func TestPUT_Ventures_1(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When an existing Venture is modified and PUT to the server
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, and OPTIONS
		And the body is a JSON array containing all updated Ventures
		And those Ventures will have new 'last_updated' datetimes
		...`)

	beginVenTest()
	defer endVenTest()

	exp, ok := livingVens["1"]
	require.True(t, ok)

	input := v.ModVenture{
		IDs:   exp.ID,
		Props: "description, order_ids, extra",
		Values: v.Venture{
			Description: "Black blizzard",
			OrderIDs:    "1,2,3",
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
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	out := v.AssertVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 1)
	v.AssertGenericVenture(t, out[0])

	input.Values.State = exp.State
	input.Values.ID = out[0].ID
	input.Values.LastModified = out[0].LastModified
	fromDB := venDBQueryOne(out[0].ID)

	assert.Equal(t, input.Values, out[0])
	assert.Equal(t, fromDB, out[0])
}

func TestPUT_Ventures_2(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When an non-existent Venture is PUT to the server
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, and OPTIONS
		And the body is an empty JSON array
		...`)

	beginVenTest()
	defer endVenTest()

	input := v.ModVenture{
		IDs:   "999999",
		Props: "description, order_ids, extra",
		Values: v.Venture{
			Description: "Black blizzard",
			OrderIDs:    "1,2,3",
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
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	out := v.AssertVentureSliceFromReader(t, res.Body)
	require.Empty(t, out)
}

func TestPUT_Ventures_3(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When a venture modification without IDs is PUT to the server
		Then ensure the response code is 400
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, and OPTIONS
		And the body is a JSON object representing an error response
		...`)

	beginVenTest()
	defer endVenTest()

	input := v.ModVenture{
		Props: "description, order_ids, extra",
		Values: v.Venture{
			Description: "Black blizzard",
			OrderIDs:    "1,2,3",
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
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 400, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	test.AssertWrappedErrorBody(t, res.Body)
}

func TestPUT_Ventures_4(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When ventures updates are PUT to the server with invalid modifications
		Then ensure the response code is 400
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, and OPTIONS
		And the body is a JSON object representing an error response
		...`)

	beginVenTest()
	defer endVenTest()

	input := v.ModVenture{
		IDs:    "1",
		Props:  "description, order_ids, extra",
		Values: v.Venture{},
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
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 400, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)
	test.AssertWrappedErrorBody(t, res.Body)
}

func TestPUT_Ventures_5(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When multiple existing Ventures are modified as dead and PUT to the server
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, and OPTIONS
		And the body is a JSON array containing all updated Ventures
		And those Ventures will have new 'last_updated' datetimes
		...`)

	beginVenTest()
	defer endVenTest()

	input := v.ModVenture{
		IDs:   "4,5",
		Props: "is_dead",
		Values: v.Venture{
			IsDead: true,
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
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	out := v.AssertVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 2)

	fromDB := venDBQueryMany("4,5")
	assert.Empty(t, fromDB)
}

// ****************************************************************************
// (PUT) /ventures?wrap
// ****************************************************************************

func TestPUT_Ventures_6(t *testing.T) {
	t.Log(`Given some Ventures already exist on the server
		When an existing Venture is modified and PUT to the server
		Then ensure the response code is 200
		And the 'wrap' query parameter has been specified
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, and OPTIONS
		And the body is a JSON object representing a WrappedReply
		And the wrapped data is a JSON array containing all updated Ventures
		And those Ventures will have new 'last_updated' datetimes
		...`)

	beginVenTest()
	defer endVenTest()

	exp, ok := livingVens["1"]
	require.True(t, ok)

	input := v.ModVenture{
		IDs:   exp.ID,
		Props: "description, state, extra",
		Values: v.Venture{
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
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", ventureHttpMethods)

	_, out := v.AssertWrappedVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 1)
	v.AssertGenericVenture(t, out[0])

	input.Values.OrderIDs = exp.OrderIDs
	input.Values.ID = out[0].ID
	input.Values.LastModified = out[0].LastModified
	fromDB := venDBQueryOne(out[0].ID)

	assert.Equal(t, input.Values, out[0])
	assert.Equal(t, fromDB, out[0])
}
