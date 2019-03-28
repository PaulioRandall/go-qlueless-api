package venture

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVenture_Clean_1(t *testing.T) {
	a := Venture{
		Description: "\n\t\v description \r\f ",
		VentureID:   "\n\t\v 1 \r\f ",
		OrderIDs:    "\n\t\v 2 \r\f , 3,4,\v 999 \f\t",
		State:       "\n\t\v state \r\f ",
		IsAlive:     true,
		Extra:       "\n\t\v extra \r\f",
	}

	a.Clean()

	assert.Equal(t, "description", a.Description)
	assert.Equal(t, "1", a.VentureID)
	assert.Equal(t, "2,3,4,999", a.OrderIDs)
	assert.Equal(t, "state", a.State)
	assert.True(t, a.IsAlive)
	assert.Equal(t, "\n\t\v extra \r\f", a.Extra)
}

func TestVenture_Clean_2(t *testing.T) {
	a := Venture{
		Description: "description",
		VentureID:   "1",
		OrderIDs:    "2,3,4,999",
		State:       "state",
	}

	a.Clean()

	assert.Equal(t, "description", a.Description)
	assert.Equal(t, "1", a.VentureID)
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
		VentureID:   "1",
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
		VentureID:   "1",
		State:       "state",
	}

	errMsgs := a.Validate()
	assert.Empty(t, errMsgs)
}

func TestVenture_Validate_3(t *testing.T) {
	a := Venture{
		Description: "",
		VentureID:   "",
		State:       "",
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 3)
}

func TestVenture_Validate_4(t *testing.T) {
	a := Venture{
		Description: "valid",
		VentureID:   "invalid",
		OrderIDs:    "3,invalid,4",
		State:       "valid",
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 2)
}
