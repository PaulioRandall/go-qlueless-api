package ventures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ****************************************************************************
// Venture.Clean
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
// Venture.Validate
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
// Venture.SplitOrderIDs
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
// Venture.SetOrderIDs
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
// VentureUpdate.SplitProps
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
// VentureUpdate.Validate
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

// TODO: DO more Validate() tests