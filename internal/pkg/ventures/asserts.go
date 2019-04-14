package ventures

import (
	"io"
	"sort"
	"testing"

	ts "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	w "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/wrapped"
	ms "github.com/mitchellh/mapstructure"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
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
	assert.True(t, v.LastModified > 0, "Venture.LastModified")
	if v.OrderIDs != "" {
		ts.AssertGenericIntCSV(t, v.OrderIDs)
	}
}

// AssertVentureModEquals asserts that the two Ventures are equal with the
// exception of the last_modified field
func AssertVentureModEquals(t *testing.T, exp Venture, act Venture) {
	assert.Equal(t, exp.ID, act.ID, "Venture.ID")
	assert.Equal(t, exp.Description, act.Description, "Venture.Description")
	assert.Equal(t, exp.OrderIDs, act.OrderIDs, "Venture.OrderIDs")
	assert.Equal(t, exp.State, act.State, "Venture.State")
	assert.Equal(t, exp.IsDead, act.IsDead, "Venture.IsDead")
	assert.Equal(t, exp.Extra, act.Extra, "Venture.Extra")
}

// AssertVentureSliceModEquals asserts that the two Venture slices are equal
// with the exception of the last_modified fields
func AssertVentureSliceModEquals(t *testing.T, exp map[string]Venture, act []Venture) {
	require.Len(t, act, len(exp))
	for _, ven := range act {
		AssertVentureModEquals(t, exp[ven.ID], ven)
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
	wr, err := w.DecodeFromReader(r)
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
	wr, err := w.DecodeFromReader(r)
	require.Nil(t, err)
	w.AssertWrappedReply(t, wr)

	assert.NotEmpty(t, wr.Message)
	assert.NotEmpty(t, wr.Self)
	require.NotEmpty(t, wr.Data)

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

// AssertOrderlessSlicesEqual asserts that the two slices contain the same
// Ventures but differences in order should be ignored
func AssertOrderlessSlicesEqual(t *testing.T, exp []Venture, act []Venture) {
	sort.Sort(ByVenID(exp))
	sort.Sort(ByVenID(act))
	assert.Equal(t, exp, act)
}
