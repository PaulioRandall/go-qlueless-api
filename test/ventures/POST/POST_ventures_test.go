package POST

import (
	"bytes"
	"encoding/json"
	"testing"

	ventures "github.com/PaulioRandall/go-qlueless-api/api/ventures"
	a "github.com/PaulioRandall/go-qlueless-api/shared/asserts"
	test "github.com/PaulioRandall/go-qlueless-api/test"
	vtest "github.com/PaulioRandall/go-qlueless-api/test/ventures"
	assert "github.com/stretchr/testify/assert"
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
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is a JSON object representing the living input Venture
		And that Venture will have a new, unused, ID
		And that Venture will have a new 'last_updated' datetime
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := ventures.Venture{
		Description: "A new Venture",
		State:       "Not started",
		Orders:      "1,2,3",
		Extra:       "Extra, extra",
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
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)

	output := ventures.AssertVentureFromReader(t, res.Body)
	ventures.AssertGenericVenture(t, output)

	input.ID = output.ID
	input.LastModified = output.LastModified
	fromDB := vtest.DBQueryOne(output.ID)

	assert.Equal(t, input, output)
	assert.Equal(t, fromDB, output)
}

func TestPOST_Venture_2(t *testing.T) {

	t.Log(`Given some Ventures already exist on the server
		When a new but invalid Venture is POSTed
		Then ensure the response code is 400
		And the 'Content-Type' header contains 'application/json'
		And 'Access-Control-Allow-Origin' is '*'
		And 'Access-Control-Allow-Headers' is '*'
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is a JSON object representing an error response
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := ventures.Venture{
		Description: "",
		State:       "",
		Orders:      "invalid",
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
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)
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
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, DELETE and OPTIONS
		And the body is a JSON object representing a WrappedReply
		And the wrapped data is a JSON object representing the living input Venture
		And that Venture will have a new, unused, ID
		And that Venture will have a new 'last_updated' datetime
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	input := ventures.Venture{
		Description: "A new Venture",
		State:       "Not started",
		Orders:      "1,2,3",
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
	test.AssertDefaultHeaders(t, res, "application/json", vtest.VenHttpMethods)

	_, output := ventures.AssertWrappedVentureFromReader(t, res.Body)
	ventures.AssertGenericVenture(t, output)

	input.ID = output.ID
	input.LastModified = output.LastModified
	fromDB := vtest.DBQueryOne(output.ID)

	assert.Equal(t, input, output)
	assert.Equal(t, fromDB, output)
}
