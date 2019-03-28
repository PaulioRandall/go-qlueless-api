package ventures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestVenture_Validate_1(t *testing.T) {
	a := Venture{
		Description: "description",
		ID:          "1",
		OrderIDs:    "2,3,4,999",
		State:       "state",
		IsAlive:     true,
		Extra:       "\n\t\v extra \r\f",
	}

	errMsgs := a.Validate()
	assert.Empty(t, errMsgs)
}

func TestVenture_Validate_2(t *testing.T) {
	a := Venture{
		Description: "description",
		ID:          "1",
		State:       "state",
	}

	errMsgs := a.Validate()
	assert.Empty(t, errMsgs)
}

func TestVenture_Validate_3(t *testing.T) {
	a := Venture{
		Description: "",
		ID:          "",
		State:       "",
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 3)
}

func TestVenture_Validate_4(t *testing.T) {
	a := Venture{
		Description: "valid",
		ID:          "invalid",
		OrderIDs:    "3,invalid,4",
		State:       "valid",
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 2)
}

func TestSplitOrderIDs_1(t *testing.T) {
	a := Venture{
		OrderIDs: "1,2,3",
	}

	s := a.SplitOrderIDs()

	assert.Len(t, s, 3)
	exp := []string{"1", "2", "3"}
	assert.Equal(t, exp, s)
}

func TestSplitOrderIDs_2(t *testing.T) {
	a := Venture{
		OrderIDs: "1",
	}

	s := a.SplitOrderIDs()

	assert.Len(t, s, 1)
	exp := []string{"1"}
	assert.Equal(t, exp, s)
}

func TestSplitOrderIDs_3(t *testing.T) {
	a := Venture{}
	s := a.SplitOrderIDs()
	assert.Empty(t, s)
}

func TestSetOrderIDs_1(t *testing.T) {
	a := Venture{}
	ids := []string{"1", "2", "3"}
	a.SetOrderIDs(ids)

	assert.Equal(t, "1,2,3", a.OrderIDs)
}

func TestSetOrderIDs_2(t *testing.T) {
	a := Venture{}
	ids := []string{"1"}
	a.SetOrderIDs(ids)

	assert.Equal(t, "1", a.OrderIDs)
}

func TestSetOrderIDs_3(t *testing.T) {
	a := Venture{}
	ids := []string{}
	a.SetOrderIDs(ids)

	assert.Equal(t, "", a.OrderIDs)
}
