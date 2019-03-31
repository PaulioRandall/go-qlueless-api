package ventures

import (
	"io"
	"testing"

	ts "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertVentureFromReader(t *testing.T, r io.Reader) Venture {
	v, err := DecodeVenture(r)
	require.Nil(t, err)
	AssertGenericVenture(t, v)
	return v
}

func AssertGenericVenture(t *testing.T, v Venture) {
	assert.NotEmpty(t, v.ID)
	assert.NotEmpty(t, v.Description)
	assert.NotEmpty(t, v.State)
	if v.OrderIDs != "" {
		ts.AssertGenericIntCSV(t, v.OrderIDs)
	}
	assert.True(t, v.IsAlive)
}

func AssertVentureSliceFromReader(t *testing.T, r io.Reader) []Venture {
	v, err := DecodeVentureSlice(r)
	require.Nil(t, err)
	AssertGenericVentureSlice(t, v)
	return v
}

func AssertGenericVentureSlice(t *testing.T, vens []Venture) {
	for _, v := range vens {
		AssertGenericVenture(t, v)
	}
}
