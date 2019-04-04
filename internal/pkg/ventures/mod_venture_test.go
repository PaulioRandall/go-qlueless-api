package ventures

import (
	"strings"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// ****************************************************************************
// DecodeModVenture()
// ****************************************************************************

func TestDecodeModVenture_1(t *testing.T) {
	a := strings.NewReader(`{
		"set": "description,state,order_ids,is_alive,extra",
		"values": {
			"description": "description",
			"venture_id": "1",
			"order_ids": "2,3,4,999",
			"state": "state",
			"is_alive": true,
			"extra": "extra"
		}
	}`)

	exp := ModVenture{
		Props: "description,state,order_ids,is_alive,extra",
		Values: Venture{
			Description: "description",
			ID:          "1",
			OrderIDs:    "2,3,4,999",
			State:       "state",
			IsAlive:     true,
			Extra:       "extra",
		},
	}

	b, err := DecodeModVenture(a)
	require.Nil(t, err)
	assert.Equal(t, exp, b)
}

func TestDecodeModVenture_2(t *testing.T) {
	a := strings.NewReader(`{}`)

	exp := ModVenture{}

	b, err := DecodeModVenture(a)
	require.Nil(t, err)
	assert.Empty(t, exp, b)
}

func TestDecodeModVenture_3(t *testing.T) {
	a := strings.NewReader(``)
	_, err := DecodeModVenture(a)
	require.NotNil(t, err)
}

// ****************************************************************************
// ModVenture.SplitIDs()
// ****************************************************************************

func TestModVenture_SplitIDs_1(t *testing.T) {
	a := ModVenture{
		IDs: "1,2,3",
	}

	s := a.SplitIDs()
	exp := []string{"1", "2", "3"}
	assert.Equal(t, exp, s)
}

func TestModVenture_SplitIDs_2(t *testing.T) {
	a := ModVenture{
		IDs: "1",
	}

	s := a.SplitIDs()
	exp := []string{"1"}
	assert.Equal(t, exp, s)
}

func TestModVenture_SplitIDs_3(t *testing.T) {
	a := ModVenture{}
	s := a.SplitIDs()
	assert.Empty(t, s)
}

// ****************************************************************************
// ModVenture.SplitProps()
// ****************************************************************************

func TestModVenture_SplitProps_1(t *testing.T) {
	a := ModVenture{
		Props: "description,state,extra",
	}

	s := a.SplitProps()
	exp := []string{"description", "state", "extra"}
	assert.Equal(t, exp, s)
}

func TestModVenture_SplitProps_2(t *testing.T) {
	a := ModVenture{
		Props: "description",
	}

	s := a.SplitProps()
	exp := []string{"description"}
	assert.Equal(t, exp, s)
}

func TestModVenture_SplitProps_3(t *testing.T) {
	a := ModVenture{}
	s := a.SplitProps()
	assert.Empty(t, s)
}

// ****************************************************************************
// ModVenture.Clean()
// ****************************************************************************

func TestModVenture_Clean_1(t *testing.T) {
	a := ModVenture{
		IDs:   "\n\t\v 1 \r\f ,    42,66\v  \f\t",
		Props: "\n\t\v   \r\f  state,extra,\v    is_alive \f\t",
		Values: Venture{
			Description: "\n\t\v description \r\f ",
			OrderIDs:    "\n\t\v 2 \r\f ,    3,4,\v 999 \f\t",
			State:       "\n\t\v state \r\f ",
			IsAlive:     true,
			Extra:       "\n\t\v extra \r\f",
		},
	}

	a.Clean()

	assert.Equal(t, "1,42,66", a.IDs)
	assert.Equal(t, "state,extra,is_alive", a.Props)
	assert.Equal(t, "description", a.Values.Description)
	assert.Equal(t, "2,3,4,999", a.Values.OrderIDs)
	assert.Equal(t, "state", a.Values.State)
	assert.True(t, a.Values.IsAlive)
	assert.Equal(t, "\n\t\v extra \r\f", a.Values.Extra)
}

func TestModVenture_Clean_2(t *testing.T) {
	a := ModVenture{
		IDs: "1,2,99",
	}
	a.Clean()
	assert.Equal(t, "1,2,99", a.IDs)
}

func TestModVenture_Clean_3(t *testing.T) {
	a := ModVenture{
		Props: "description,state,extra",
	}
	a.Clean()
	assert.Equal(t, "description,state,extra", a.Props)
}

func TestModVenture_Clean_4(t *testing.T) {
	a := ModVenture{}
	a.Clean()
	assert.Equal(t, ModVenture{}, a)
}

// ****************************************************************************
// ModVenture.Validate()
// ****************************************************************************

func TestModVenture_Validate_1(t *testing.T) {
	a := ModVenture{
		IDs:   "1,2,3",
		Props: "description,state,extra",
		Values: Venture{
			Description: "updated description",
			OrderIDs:    "66,101,202",
			State:       "updated state",
			IsAlive:     false,
			Extra:       "updated extra",
		},
	}

	errMsgs := a.Validate()
	assert.Empty(t, errMsgs)
}

func TestModVenture_Validate_2(t *testing.T) {
	a := ModVenture{
		IDs:   "1",
		Props: "description",
		Values: Venture{
			Description: "updated description",
		},
	}

	errMsgs := a.Validate()
	assert.Empty(t, errMsgs)
}

func TestModVenture_Validate_3(t *testing.T) {
	a := ModVenture{
		IDs:    "1",
		Props:  "",
		Values: Venture{},
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 1)
}

func TestModVenture_Validate_4(t *testing.T) {
	a := ModVenture{
		IDs:    "1",
		Props:  "description",
		Values: Venture{},
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 1)
}

func TestModVenture_Validate_5(t *testing.T) {
	a := ModVenture{
		IDs:    "1",
		Props:  "INVALID",
		Values: Venture{},
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 1)
}

func TestModVenture_Validate_6(t *testing.T) {
	a := ModVenture{
		IDs:    "1",
		Props:  "extra",
		Values: Venture{},
	}

	errMsgs := a.Validate()
	assert.Empty(t, errMsgs)
}

func TestModVenture_Validate_7(t *testing.T) {
	a := ModVenture{
		IDs:   "1",
		Props: "order_ids",
		Values: Venture{
			OrderIDs: "1,2,ILLEGAL,99,100",
		},
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 1)
}

func TestModVenture_Validate_8(t *testing.T) {
	a := ModVenture{
		Props: "description,,state,order_ids,INVALID,extra",
		Values: Venture{
			OrderIDs: "ILLEGAL,66,101,202",
			State:    "updated state",
			IsAlive:  false,
			Extra:    "updated extra",
		},
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 5)
}

func TestModVenture_Validate_9(t *testing.T) {
	a := ModVenture{
		IDs:   "",
		Props: "description",
		Values: Venture{
			Description: "description",
		},
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 1)
}
