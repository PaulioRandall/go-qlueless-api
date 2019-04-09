package ventures

import (
	"testing"

	a "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	test "github.com/PaulioRandall/go-qlueless-assembly-api/test"
	require "github.com/stretchr/testify/require"
)

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
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, and OPTIONS
		And there is NO response body
		...`)

	beginVenTest()
	defer endVenTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "OPTIONS",
	}
	res := req.Fire()
	defer res.Body.Close()
	defer a.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertNoContentHeaders(t, res, ventureHttpMethods)
	test.AssertEmptyBody(t, res.Body)
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
		And 'Access-Control-Allow-Methods' only contains GET, POST, PUT, and OPTIONS
		And there is NO response body
		...`)

	beginVenTest()
	defer endVenTest()

	test.VerifyNotAllowedMethods(t, "http://localhost:8080/ventures", ventureHttpMethods)
}
