package OPTIONS

import (
	"testing"

	test "github.com/PaulioRandall/go-qlueless-api/test"
	vtest "github.com/PaulioRandall/go-qlueless-api/test/ventures"
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
		And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
		And there is NO response body
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	req := test.APICall{
		URL:    "http://localhost:8080/ventures",
		Method: "OPTIONS",
	}
	res := req.Fire()

	defer res.Body.Close()
	defer test.PrintResponse(t, res.Body)

	require.Equal(t, 200, res.StatusCode)
	test.AssertCorsHeaders(t, res, "GET, POST, PUT, DELETE, OPTIONS")
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
		And 'Access-Control-Allow-Methods' is 'GET, POST, PUT, DELETE, OPTIONS'
		And there is NO response body
		...`)

	vtest.BeginTest("../../../bin")
	defer vtest.EndTest()

	goodMethods := "GET, POST, PUT, DELETE, OPTIONS"
	test.VerifyBadMethods(t, "http://localhost:8080/ventures", goodMethods, []string{
		"HEAD",
		"CONNECT",
		"TRACE",
		"PATCH",
		"CUSTOM",
	})
}
