package ventures

import (
	"strings"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// ****************************************************************************
// DecodeVentureUpdate()
// ****************************************************************************

func TestDecodeVentureUpdate_1(t *testing.T) {
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

	exp := VentureUpdate{
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

	b, err := DecodeVentureUpdate(a)
	require.Nil(t, err)
	assert.Equal(t, exp, b)
}

func TestDecodeVentureUpdate_2(t *testing.T) {
	a := strings.NewReader(`{}`)

	exp := VentureUpdate{}

	b, err := DecodeVentureUpdate(a)
	require.Nil(t, err)
	assert.Empty(t, exp, b)
}

func TestDecodeVentureUpdate_3(t *testing.T) {
	a := strings.NewReader(``)
	_, err := DecodeVentureUpdate(a)
	require.NotNil(t, err)
}

// ****************************************************************************
// VentureUpdate.SplitProps()
// ****************************************************************************

func TestVentureUpdate_SplitProps_1(t *testing.T) {
	a := VentureUpdate{
		Props: "description,state,extra",
	}

	s := a.SplitProps()
	exp := []string{"description", "state", "extra"}
	assert.Equal(t, exp, s)
}

func TestVentureUpdate_SplitProps_2(t *testing.T) {
	a := VentureUpdate{
		Props: "description",
	}

	s := a.SplitProps()
	exp := []string{"description"}
	assert.Equal(t, exp, s)
}

func TestVentureUpdate_SplitProps_3(t *testing.T) {
	a := VentureUpdate{}
	s := a.SplitProps()
	assert.Empty(t, s)
}

// ****************************************************************************
// VentureUpdate.Clean()
// ****************************************************************************

func TestVentureUpdate_Clean_1(t *testing.T) {
	a := VentureUpdate{
		Props: "\n\t\v   \r\f  state,extra,\v    is_alive \f\t",
		Values: Venture{
			Description: "\n\t\v description \r\f ",
			ID:          "\n\t\v 1 \r\f ",
			OrderIDs:    "\n\t\v 2 \r\f ,    3,4,\v 999 \f\t",
			State:       "\n\t\v state \r\f ",
			IsAlive:     true,
			Extra:       "\n\t\v extra \r\f",
		},
	}

	a.Clean()

	assert.Equal(t, "state,extra,is_alive", a.Props)
	assert.Equal(t, "description", a.Values.Description)
	assert.Equal(t, "1", a.Values.ID)
	assert.Equal(t, "2,3,4,999", a.Values.OrderIDs)
	assert.Equal(t, "state", a.Values.State)
	assert.True(t, a.Values.IsAlive)
	assert.Equal(t, "\n\t\v extra \r\f", a.Values.Extra)
}

func TestVentureUpdate_Clean_2(t *testing.T) {
	a := VentureUpdate{
		Props: "description,state,extra",
	}
	a.Clean()
	assert.Equal(t, "description,state,extra", a.Props)
}

func TestVentureUpdate_Clean_3(t *testing.T) {
	a := VentureUpdate{}
	a.Clean()
	assert.Equal(t, VentureUpdate{}, a)
}

// ****************************************************************************
// VentureUpdate.Validate()
// ****************************************************************************

func TestVentureUpdate_Validate_1(t *testing.T) {
	a := VentureUpdate{
		Props: "description,state,extra",
		Values: Venture{
			ID:          "1",
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

func TestVentureUpdate_Validate_2(t *testing.T) {
	a := VentureUpdate{
		Props: "description",
		Values: Venture{
			ID:          "1",
			Description: "updated description",
		},
	}

	errMsgs := a.Validate()
	assert.Empty(t, errMsgs)
}

func TestVentureUpdate_Validate_3(t *testing.T) {
	a := VentureUpdate{
		Props: "",
		Values: Venture{
			ID: "1",
		},
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 1)
}

func TestVentureUpdate_Validate_4(t *testing.T) {
	a := VentureUpdate{
		Props: "description",
		Values: Venture{
			ID: "1",
		},
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 1)
}

func TestVentureUpdate_Validate_5(t *testing.T) {
	a := VentureUpdate{
		Props: "INVALID",
		Values: Venture{
			ID: "1",
		},
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 1)
}

func TestVentureUpdate_Validate_6(t *testing.T) {
	a := VentureUpdate{
		Props: "extra",
		Values: Venture{
			ID: "1",
		},
	}

	errMsgs := a.Validate()
	assert.Empty(t, errMsgs)
}

func TestVentureUpdate_Validate_7(t *testing.T) {
	a := VentureUpdate{
		Props: "order_ids",
		Values: Venture{
			ID:       "1",
			OrderIDs: "1,2,ILLEGAL,99,100",
		},
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 1)
}

func TestVentureUpdate_Validate_8(t *testing.T) {
	a := VentureUpdate{
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
