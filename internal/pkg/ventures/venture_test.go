package ventures

import (
	"strings"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// ****************************************************************************
// DecodeNewVenture()
// ****************************************************************************

func TestDecodeNewVenture_1(t *testing.T) {
	a := strings.NewReader(`{
		"description": "description",
		"order_ids": "2,3,4,999",
		"state": "state",
		"extra": "extra"
	}`)

	exp := NewVenture{
		Description: "description",
		OrderIDs:    "2,3,4,999",
		State:       "state",
		Extra:       "extra",
	}

	b, err := DecodeNewVenture(a)
	require.Nil(t, err)
	assert.Equal(t, exp, b)
}

func TestDecodeNewVenture_2(t *testing.T) {
	a := strings.NewReader(`{}`)

	exp := NewVenture{}

	b, err := DecodeNewVenture(a)
	require.Nil(t, err)
	assert.Empty(t, exp, b)
}

func TestDecodeNewVenture_3(t *testing.T) {
	a := strings.NewReader(``)
	_, err := DecodeNewVenture(a)
	require.NotNil(t, err)
}

// ****************************************************************************
// NewVenture.Clean()
// ****************************************************************************
func TestNewVenture_Clean_1(t *testing.T) {
	a := NewVenture{
		Description: "\n\t\v description \r\f ",
		OrderIDs:    "\n\t\v 2 \r\f , 3,4,\v 999 \f\t",
		State:       "\n\t\v state \r\f ",
		Extra:       "\n\t\v extra \r\f",
	}

	a.Clean()

	assert.Equal(t, "description", a.Description)
	assert.Equal(t, "2,3,4,999", a.OrderIDs)
	assert.Equal(t, "state", a.State)
	assert.Equal(t, "\n\t\v extra \r\f", a.Extra)
}

func TestNewVenture_Clean_2(t *testing.T) {
	a := NewVenture{
		Description: "description",
		OrderIDs:    "2,3,4,999",
		State:       "state",
	}

	a.Clean()

	assert.Equal(t, "description", a.Description)
	assert.Equal(t, "2,3,4,999", a.OrderIDs)
	assert.Equal(t, "state", a.State)
}

func TestNewVenture_Clean_3(t *testing.T) {
	a := NewVenture{}
	b := NewVenture{}
	a.Clean()
	assert.Equal(t, b, a)
}

// ****************************************************************************
// NewVenture.Validate()
// ****************************************************************************

func TestNewVenture_Validate_1(t *testing.T) {
	a := NewVenture{
		Description: "description",
		OrderIDs:    "2,3,4,999",
		State:       "state",
		Extra:       "\n\t\v extra \r\f",
	}

	errMsgs := a.Validate()
	assert.Empty(t, errMsgs)
}

func TestNewVenture_Validate_2(t *testing.T) {
	a := NewVenture{
		Description: "description",
		State:       "state",
	}

	errMsgs := a.Validate()
	assert.Empty(t, errMsgs)
}

func TestNewVenture_Validate_3(t *testing.T) {
	a := NewVenture{
		Description: "",
		State:       "",
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 2)
}

func TestNewVenture_Validate_4(t *testing.T) {
	a := NewVenture{
		Description: "valid",
		OrderIDs:    "3,invalid,4",
		State:       "valid",
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 1)
}

func TestNewVenture_Validate_5(t *testing.T) {
	a := NewVenture{
		Description: "",
		State:       "",
		OrderIDs:    "3,invalid,4",
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 3)
}

// ****************************************************************************
// DecodeVenture()
// ****************************************************************************

func TestDecodeVenture_1(t *testing.T) {
	a := strings.NewReader(`{
		"description": "description",
		"venture_id": "1",
		"order_ids": "2,3,4,999",
		"state": "state",
		"is_alive": true,
		"extra": "extra"
	}`)

	exp := Venture{
		Description: "description",
		ID:          "1",
		OrderIDs:    "2,3,4,999",
		State:       "state",
		IsAlive:     true,
		Extra:       "extra",
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
			"venture_id": "1",
			"order_ids": "2,3,4",
			"state": "1",
			"is_alive": true,
			"extra": "1"
		},
		{
			"description": "2",
			"venture_id": "2",
			"order_ids": "5,6,7",
			"state": "2",
			"is_alive": false,
			"extra": "2"
		}
	]`)

	exp := []Venture{
		Venture{
			Description: "1",
			ID:          "1",
			OrderIDs:    "2,3,4",
			State:       "1",
			IsAlive:     true,
			Extra:       "1",
		},
		Venture{
			Description: "2",
			ID:          "2",
			OrderIDs:    "5,6,7",
			State:       "2",
			IsAlive:     false,
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
		OrderIDs:    "\n\t\v 2 \r\f , 3,4,\v 999 \f\t",
		State:       "\n\t\v state \r\f ",
		IsAlive:     true,
		Extra:       "\n\t\v extra \r\f",
	}

	a.Clean()

	assert.Equal(t, "description", a.Description)
	assert.Equal(t, "1", a.ID)
	assert.Equal(t, "2,3,4,999", a.OrderIDs)
	assert.Equal(t, "state", a.State)
	assert.True(t, a.IsAlive)
	assert.Equal(t, "\n\t\v extra \r\f", a.Extra)
}

func TestVenture_Clean_2(t *testing.T) {
	a := Venture{
		Description: "description",
		ID:          "1",
		OrderIDs:    "2,3,4,999",
		State:       "state",
	}

	a.Clean()

	assert.Equal(t, "description", a.Description)
	assert.Equal(t, "1", a.ID)
	assert.Equal(t, "2,3,4,999", a.OrderIDs)
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
		Description: "description",
		ID:          "1",
		OrderIDs:    "2,3,4,999",
		State:       "state",
		IsAlive:     true,
		Extra:       "\n\t\v extra \r\f",
	}

	errMsgs := a.Validate(false)
	assert.Empty(t, errMsgs)
}

func TestVenture_Validate_2(t *testing.T) {
	a := Venture{
		Description: "description",
		ID:          "1",
		State:       "state",
	}

	errMsgs := a.Validate(false)
	assert.Empty(t, errMsgs)
}

func TestVenture_Validate_3(t *testing.T) {
	a := Venture{
		Description: "",
		ID:          "",
		State:       "",
	}

	errMsgs := a.Validate(false)
	assert.Len(t, errMsgs, 3)
}

func TestVenture_Validate_4(t *testing.T) {
	a := Venture{
		Description: "valid",
		ID:          "invalid",
		OrderIDs:    "3,invalid,4",
		State:       "valid",
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
		OrderIDs:    "3,invalid,4",
	}

	errMsgs := a.Validate(true)
	assert.Len(t, errMsgs, 3)
}

// ****************************************************************************
// Venture.SplitOrderIDs()
// ****************************************************************************

func TestVenture_SplitOrderIDs_1(t *testing.T) {
	a := Venture{
		OrderIDs: "1,2,3",
	}

	s := a.SplitOrderIDs()
	exp := []string{"1", "2", "3"}
	assert.Equal(t, exp, s)
}

func TestVenture_SplitOrderIDs_2(t *testing.T) {
	a := Venture{
		OrderIDs: "1",
	}

	s := a.SplitOrderIDs()
	exp := []string{"1"}
	assert.Equal(t, exp, s)
}

func TestVenture_SplitOrderIDs_3(t *testing.T) {
	a := Venture{}
	s := a.SplitOrderIDs()
	assert.Empty(t, s)
}

// ****************************************************************************
// Venture.SetOrderIDs()
// ****************************************************************************

func TestVenture_SetOrderIDs_1(t *testing.T) {
	a := Venture{}
	ids := []string{"1", "2", "3"}
	a.SetOrderIDs(ids)

	assert.Equal(t, "1,2,3", a.OrderIDs)
}

func TestVenture_SetOrderIDs_2(t *testing.T) {
	a := Venture{}
	ids := []string{"1"}
	a.SetOrderIDs(ids)

	assert.Equal(t, "1", a.OrderIDs)
}

func TestVenture_SetOrderIDs_3(t *testing.T) {
	a := Venture{}
	ids := []string{}
	a.SetOrderIDs(ids)

	assert.Equal(t, "", a.OrderIDs)
}

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
