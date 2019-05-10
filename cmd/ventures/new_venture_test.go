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
		"orders": "2,3,4,999",
		"state": "state",
		"extra": "extra"
	}`)

	exp := NewVenture{
		Description: "description",
		Orders:      "2,3,4,999",
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
		Orders:      "\n\t\v 2 \r\f , 3,4,\v 999 \f\t",
		State:       "\n\t\v state \r\f ",
		Extra:       "\n\t\v extra \r\f",
	}

	a.Clean()

	assert.Equal(t, "description", a.Description)
	assert.Equal(t, "2,3,4,999", a.Orders)
	assert.Equal(t, "state", a.State)
	assert.Equal(t, "\n\t\v extra \r\f", a.Extra)
}

func TestNewVenture_Clean_2(t *testing.T) {
	a := NewVenture{
		Description: "description",
		Orders:      "2,3,4,999",
		State:       "state",
	}

	a.Clean()

	assert.Equal(t, "description", a.Description)
	assert.Equal(t, "2,3,4,999", a.Orders)
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
		Orders:      "2,3,4,999",
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
		Orders:      "3,invalid,4",
		State:       "valid",
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 1)
}

func TestNewVenture_Validate_5(t *testing.T) {
	a := NewVenture{
		Description: "",
		State:       "",
		Orders:      "3,invalid,4",
	}

	errMsgs := a.Validate()
	assert.Len(t, errMsgs, 3)
}
