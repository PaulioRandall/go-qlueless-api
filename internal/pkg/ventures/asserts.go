package ventures

import (
	"io"
	"testing"

	ts "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	w "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/wrapped"
	ms "github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertVentureFromReader asserts that a Venture decoded from an io.Reader
// has the required fields populated and in the correct format
func AssertVentureFromReader(t *testing.T, r io.Reader) Venture {
	v, err := DecodeVenture(r)
	require.Nil(t, err)
	AssertGenericVenture(t, v)
	return v
}

// AssertVentureFromReader asserts that a Venture has the required fields
// populated and in the correct format
func AssertGenericVenture(t *testing.T, v Venture) {
	assert.NotEmpty(t, v.ID, "Venture.ID")
	assert.NotEmpty(t, v.Description, "Venture.Description")
	assert.NotEmpty(t, v.State, "Venture.State")
	if v.OrderIDs != "" {
		ts.AssertGenericIntCSV(t, v.OrderIDs)
	}
}

// AssertVentureSliceFromReader asserts that a Venture slice decoded from an
// io.Reader has the required fields populated and in the correct format
func AssertVentureSliceFromReader(t *testing.T, r io.Reader) []Venture {
	v, err := DecodeVentureSlice(r)
	require.Nil(t, err)
	AssertGenericVentureSlice(t, v)
	return v
}

// AssertGenericVentureSlice asserts that a Venture slice has the required
// fields populated and in the correct format
func AssertGenericVentureSlice(t *testing.T, vens []Venture) {
	for _, v := range vens {
		AssertGenericVenture(t, v)
	}
}

// AssertWrappedVentureSliceFromReader asserts that a Venture slice wrapped
// within a WrappedReply and decoded from an io.Reader has the required fields
// populated and in the correct format
func AssertWrappedVentureSliceFromReader(t *testing.T, r io.Reader) (w.WrappedReply, []Venture) {
	wr, err := w.DecodeWrappedReplyFromReader(r)
	require.Nil(t, err)
	w.AssertWrappedReply(t, wr)

	var v []Venture
	config := ms.DecoderConfig{
		TagName: "json",
		Result:  &v,
	}

	d, err := ms.NewDecoder(&config)
	require.Nil(t, err)

	err = d.Decode(wr.Data)
	require.Nil(t, err)

	AssertGenericVentureSlice(t, v)
	return wr, v
}

// AssertWrappedVentureFromReader asserts that a Venture wrapped within a
// WrappedReply and decoded from an io.Reader has the required fields populated
// and in the correct format
func AssertWrappedVentureFromReader(t *testing.T, r io.Reader) (w.WrappedReply, Venture) {
	wr, err := w.DecodeWrappedReplyFromReader(r)
	require.Nil(t, err)
	w.AssertWrappedReply(t, wr)

	var v Venture
	config := ms.DecoderConfig{
		TagName: "json",
		Result:  &v,
	}

	d, err := ms.NewDecoder(&config)
	require.Nil(t, err)

	err = d.Decode(wr.Data)
	require.Nil(t, err)

	AssertGenericVenture(t, v)
	return wr, v
}
