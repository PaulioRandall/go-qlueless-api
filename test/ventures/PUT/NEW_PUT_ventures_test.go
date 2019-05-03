package PUT

import (
	"bytes"
	"encoding/json"
	"testing"

	a "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	v "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/ventures"
	test "github.com/PaulioRandall/go-qlueless-assembly-api/test"
	vtest "github.com/PaulioRandall/go-qlueless-assembly-api/test/ventures"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

const FEATURE_OFF = true

// ****************************************************************************
// (PUT) /ventures
// ****************************************************************************

func TestPUT_Ventures_1(t *testing.T) {
	if FEATURE_OFF {
		return
	}

	t.Log(`Given a Venture already exists on the server
		When the Venture is modified and PUT to the server
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is a JSON object containing a 'message' and 'self' URL
		And the updated Venture has an appropriate 'last_modified' value
		...`)

	var before *v.Venture
	var input *v.Venture
	var results []v.Venture
	var result v.Venture
	var after v.Venture

	vtest.BeginEmptyTest("../../../bin")
	defer vtest.EndTest()

	vtest.DBInject(v.NewVenture{
		Description: "Black blizzard",
		State:       "STARTED",
		Orders:      "1,2,3",
		Extra:       "colour: black",
	})

	before = vtest.DBQueryFirst()
	require.NotNil(t, before)

	id := before.ID

	input = &v.Venture{
		ID:          id,
		Description: "White wizzard",
		State:       "FINISHED",
		Orders:      "4,5,6",
		Extra:       "colour: white",
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(input)

	req := test.APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "PUT",
		Body:   buf,
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)

	results = v.AssertVentureSliceFromReader(t, res.Body)
	require.Len(t, results, 1)

	result = results[0]
	v.AssertGenericVenture(t, result)

	input.LastModified = result.LastModified
	assert.Equal(t, input, result)

	after = vtest.DBQueryOne(id)
	v.AssertEqualsModified(t, &after, &result)
}

func TestPUT_Ventures_2(t *testing.T) {
	if FEATURE_OFF {
		return
	}

	t.Log(`Given some Ventures already exist on the server
		When an non-existent Venture is PUT to the server
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is an empty JSON array
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := v.ModVenture{
		IDs:   "999999",
		Props: "description, orders, extra",
		Values: v.Venture{
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
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)

	out := v.AssertVentureSliceFromReader(t, res.Body)
	require.Empty(t, out)
}

func TestPUT_Ventures_3(t *testing.T) {
	if FEATURE_OFF {
		return
	}

	t.Log(`Given some Ventures already exist on the server
		When a venture modification without IDs is PUT to the server
		Then ensure the response code is 400
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is a JSON object representing an error response
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := v.ModVenture{
		Props: "description, orders, extra",
		Values: v.Venture{
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
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 400, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)
	test.AssertWrappedErrorBody(t, res.Body)
}

func TestPUT_Ventures_4(t *testing.T) {
	if FEATURE_OFF {
		return
	}

	t.Log(`Given some Ventures already exist on the server
		When ventures updates are PUT to the server with invalid modifications
		Then ensure the response code is 400
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is a JSON object representing an error response
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := v.ModVenture{
		IDs:    "1",
		Props:  "description, orders, extra",
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
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)
	test.AssertWrappedErrorBody(t, res.Body)
}

func TestPUT_Ventures_5(t *testing.T) {
	if FEATURE_OFF {
		return
	}

	t.Log(`Given some Ventures already exist on the server
		When multiple existing Ventures are modified as dead and PUT to the server
		Then ensure the response code is 200
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is a JSON array containing all updated Ventures
		And those Ventures will have new 'last_updated' datetimes
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := v.ModVenture{
		IDs:   "4,5",
		Props: "dead",
		Values: v.Venture{
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
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)

	out := v.AssertVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 2)

	fromDB := vtest.DBQueryMany("4,5")
	assert.Empty(t, fromDB)
}

// ****************************************************************************
// (PUT) /ventures?wrap
// ****************************************************************************

func TestPUT_Ventures_6(t *testing.T) {
	if FEATURE_OFF {
		return
	}

	t.Log(`Given some Ventures already exist on the server
		When an existing Venture is modified and PUT to the server
		Then ensure the response code is 200
		And the 'wrap' query parameter has been specified
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is a JSON object representing a WrappedReply
		And the wrapped data is a JSON array containing all updated Ventures
		And those Ventures will have new 'last_updated' datetimes
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := v.ModVenture{
		IDs:   "1",
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
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)

	_, out := v.AssertWrappedVentureSliceFromReader(t, res.Body)
	require.Len(t, out, 1)
	v.AssertGenericVenture(t, out[0])

	input.Values.Orders = out[0].Orders
	input.Values.ID = out[0].ID
	input.Values.LastModified = out[0].LastModified
	assert.Equal(t, input.Values, out[0])

	fromDB := vtest.DBQueryOne(out[0].ID)
	v.AssertVentureModEquals(t, fromDB, out[0])
}
