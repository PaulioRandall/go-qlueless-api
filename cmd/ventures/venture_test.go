package ventures

import (
	"strings"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// ****************************************************************************
// DecodeVenture()
// ****************************************************************************

func TestDecodeVenture_1(t *testing.T) {
	a := strings.NewReader(`{
		"id": "1",
		"last_modified": 1554458321281,
		"description": "description",
		"orders": "2,3,4,999",
		"state": "state",
		"dead": false,
		"extra": "extra"
	}`)

	exp := Venture{
		ID:           "1",
		LastModified: 1554458321281,
		Description:  "description",
		Orders:       "2,3,4,999",
		State:        "state",
		Dead:         false,
		Extra:        "extra",
	}

	b, err := DecodeVenture(a)
	require.Nil(t, err)
	assert.Equal(t, exp, b)
}

func TestDecodeVenture_2(t *testing.T) {
	a := strings.NewReader(`{}`)

	exp := Venture{}

	b, err := DecodeVenture(a)
	require.Nil(t, err)
	assert.Empty(t, exp, b)
}

func TestDecodeVenture_3(t *testing.T) {
	a := strings.NewReader(``)
	_, err := DecodeVenture(a)
	require.NotNil(t, err)
}

// ****************************************************************************
// DecodeVentureSlice()
// ****************************************************************************

func TestDecodeVentureSlice_1(t *testing.T) {
	a := strings.NewReader(`[
		{
			"description": "1",
			"id": "1",
			"orders": "2,3,4",
			"state": "1",
			"dead": false,
			"extra": "1"
		},
		{
			"description": "2",
			"id": "2",
			"orders": "5,6,7",
			"state": "2",
			"dead": true,
			"extra": "2"
		}
	]`)

	exp := []Venture{
		Venture{
			Description: "1",
			ID:          "1",
			Orders:      "2,3,4",
			State:       "1",
			Dead:        false,
			Extra:       "1",
		},
		Venture{
			Description: "2",
			ID:          "2",
			Orders:      "5,6,7",
			State:       "2",
			Dead:        true,
			Extra:       "2",
		},
	}

	b, err := DecodeVentureSlice(a)
	require.Nil(t, err)
	assert.Equal(t, exp, b)
}

func TestDecodeVentureSlice_2(t *testing.T) {
	a := strings.NewReader(`[]`)
	exp := []Venture{}

	b, err := DecodeVentureSlice(a)
	require.Nil(t, err)
	assert.Empty(t, exp, b)
}

func TestDecodeVentureSlice_3(t *testing.T) {
	a := strings.NewReader(``)
	_, err := DecodeVentureSlice(a)
	require.NotNil(t, err)
}

// ****************************************************************************
// Venture.Clean()
// ****************************************************************************
func TestVenture_Clean_1(t *testing.T) {
	a := Venture{
		Description: "\n\t\v description \r\f ",
		ID:          "\n\t\v 1 \r\f ",
		Orders:      "\n\t\v 2 \r\f , 3,4,\v 999 \f\t",
		State:       "\n\t\v state \r\f ",
		Dead:        false,
		Extra:       "\n\t\v extra \r\f",
	}

	a.Clean()

	assert.Equal(t, "description", a.Description)
	assert.Equal(t, "1", a.ID)
	assert.Equal(t, "2,3,4,999", a.Orders)
	assert.Equal(t, "state", a.State)
	assert.False(t, a.Dead)
	assert.Equal(t, "\n\t\v extra \r\f", a.Extra)
}

func TestVenture_Clean_2(t *testing.T) {
	a := Venture{
		Description: "description",
		ID:          "1",
		Orders:      "2,3,4,999",
		State:       "state",
	}

	a.Clean()

	assert.Equal(t, "description", a.Description)
	assert.Equal(t, "1", a.ID)
	assert.Equal(t, "2,3,4,999", a.Orders)
	assert.Equal(t, "state", a.State)
}

func TestVenture_Clean_3(t *testing.T) {
	a := Venture{}
	b := Venture{}

	a.Clean()

	assert.Equal(t, b, a)
}

// ****************************************************************************
// Venture.Validate()
// ****************************************************************************

func TestVenture_Validate_1(t *testing.T) {
	a := Venture{
		ID:           "1",
		LastModified: 1554458321281,
		Description:  "description",
		Orders:       "2,3,4,999",
		State:        "state",
		Dead:         false,
		Extra:        "\n\t\v extra \r\f",
	}

	errMsgs := a.Validate(false)
	assert.Empty(t, errMsgs)
}

func TestVenture_Validate_2(t *testing.T) {
	a := Venture{
		ID:           "1",
		LastModified: 1554458321281,
		Description:  "description",
		State:        "state",
	}

	errMsgs := a.Validate(false)
	assert.Empty(t, errMsgs)
}

func TestVenture_Validate_3(t *testing.T) {
	a := Venture{
		Description:  "",
		ID:           "",
		State:        "",
		LastModified: 0,
	}

	errMsgs := a.Validate(false)
	assert.Len(t, errMsgs, 4)
}

func TestVenture_Validate_4(t *testing.T) {
	a := Venture{
		Description:  "valid",
		ID:           "invalid",
		Orders:       "3,invalid,4",
		State:        "valid",
		LastModified: 1554458321281,
	}

	errMsgs := a.Validate(false)
	assert.Len(t, errMsgs, 2)
}

func TestVenture_Validate_5(t *testing.T) {
	a := Venture{
		Description: "description",
		State:       "state",
	}

	errMsgs := a.Validate(true)
	assert.Empty(t, errMsgs)
}

func TestVenture_Validate_6(t *testing.T) {
	a := Venture{
		Description: "",
		State:       "",
		Orders:      "3,invalid,4",
	}

	errMsgs := a.Validate(true)
	assert.Len(t, errMsgs, 3)
}

// ****************************************************************************
// Venture.SplitOrders()
// ****************************************************************************

func TestVenture_SplitOrders_1(t *testing.T) {
	a := Venture{
		Orders: "1,2,3",
	}

	s := a.SplitOrders()
	exp := []string{"1", "2", "3"}
	assert.Equal(t, exp, s)
}

func TestVenture_SplitOrders_2(t *testing.T) {
	a := Venture{
		Orders: "1",
	}

	s := a.SplitOrders()
	exp := []string{"1"}
	assert.Equal(t, exp, s)
}

func TestVenture_SplitOrders_3(t *testing.T) {
	a := Venture{}
	s := a.SplitOrders()
	assert.Empty(t, s)
}

// ****************************************************************************
// Venture.SetOrders()
// ****************************************************************************

func TestVenture_SetOrders_1(t *testing.T) {
	a := Venture{}
	ids := []string{"1", "2", "3"}
	a.SetOrders(ids)

	assert.Equal(t, "1,2,3", a.Orders)
}

func TestVenture_SetOrders_2(t *testing.T) {
	a := Venture{}
	ids := []string{"1"}
	a.SetOrders(ids)

	assert.Equal(t, "1", a.Orders)
}

func TestVenture_SetOrders_3(t *testing.T) {
	a := Venture{}
	ids := []string{}
	a.SetOrders(ids)

	assert.Equal(t, "", a.Orders)
}
