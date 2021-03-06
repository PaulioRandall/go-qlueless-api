package ventures

import (
	"io"
	"sort"
	"testing"

	"github.com/PaulioRandall/go-cookies/cookies"
	"github.com/PaulioRandall/go-qlueless-api/shared/wrapped"
	ms "github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RequireSliceOfVentures asserts that the value decoded from 'r' is a slice of
// valid Ventures; the slice is returned.
func RequireSliceOfVentures(t *testing.T, r io.Reader) []Venture {
	v, err := DecodeVentureSlice(r)
	require.Nil(t, err, "Error decoding slice of ventures from reader")
	return v
}

// AssertVentureSlice asserts that all Ventures 'vens' have valid content.
func AssertVentureSlice(t *testing.T, vens []Venture) {
	for _, v := range vens {
		AssertVenture(t, v)
	}
}

// AssertVenture asserts that a Venture 'v' has valid content.
func AssertVenture(t *testing.T, v Venture) {
	assert.NotEmpty(t, v.ID, "Venture.ID")
	assert.NotEmpty(t, v.Description, "Venture.Description")
	assert.NotEmpty(t, v.State, "Venture.State")
	assert.True(t, v.LastModified > 0, "Venture.LastModified")
	if v.Orders != "" {
		assert.True(t, cookies.IsUintCSV(v.Orders), "Venture.Orders")
	}
}

// AssertVenturesEqual asserts that 'exp' and 'act' contain the same
// Ventures. If 'orderless' is true then the slices are ordered by ID first.
func AssertVenturesEqual(t *testing.T, exp []Venture, act []Venture, orderless bool) {
	if orderless {
		sort.Sort(ByVenID(exp))
		sort.Sort(ByVenID(act))
	}
	assert.Equal(t, exp, act)
}

//
// OLD
//

// AssertVentureFromReader asserts that a Venture decoded from an io.Reader
// has the required fields populated and in the correct format
func AssertVentureFromReader(t *testing.T, r io.Reader) Venture {
	v, err := DecodeVenture(r)
	require.Nil(t, err)
	AssertGenericVenture(t, v)
	return v
}

// AssertGenericVenture asserts that a Venture has the required fields
// populated and in the correct format
func AssertGenericVenture(t *testing.T, v Venture) {
	assert.NotEmpty(t, v.ID, "Venture.ID")
	assert.NotEmpty(t, v.Description, "Venture.Description")
	assert.NotEmpty(t, v.State, "Venture.State")
	assert.True(t, v.LastModified > 0, "Venture.LastModified")
	if v.Orders != "" {
		assert.True(t, cookies.IsUintCSV(v.Orders), "Venture.Orders")
	}
}

// AssertVentureModEquals asserts that the two Ventures are equal with the
// exception of the last_modified field
func AssertVentureModEquals(t *testing.T, exp Venture, act Venture) {
	assert.Equal(t, exp.ID, act.ID, "Venture.ID")
	assert.Equal(t, exp.Description, act.Description, "Venture.Description")
	assert.Equal(t, exp.Orders, act.Orders, "Venture.Orders")
	assert.Equal(t, exp.State, act.State, "Venture.State")
	assert.Equal(t, exp.Dead, act.Dead, "Venture.Dead")
	assert.Equal(t, exp.Extra, act.Extra, "Venture.Extra")
}

// AssertEqualsModified asserts that the before Venture equals the after
// Venture except for the 'LastModified' field which must be greater then the
// after Venture.
func AssertEqualsModified(t *testing.T, before *Venture, after *Venture) {
	assert.Equal(t, before.ID, after.ID, "Venture.ID")
	assert.Equal(t, before.Description, after.Description, "Venture.Description")
	assert.Equal(t, before.Orders, after.Orders, "Venture.Orders")
	assert.Equal(t, before.State, after.State, "Venture.State")
	assert.Equal(t, before.Dead, after.Dead, "Venture.Dead")
	assert.Equal(t, before.Extra, after.Extra, "Venture.Extra")
	assert.True(t, before.LastModified < after.LastModified, "Venture.LastModified")
}

// AssertGenericReplyFromReader asserts that reading from an io.Reader produces
// a generic reply.
func AssertGenericReplyFromReader(t *testing.T, r io.Reader) {
	wr, err := wrapped.DecodeFromReader(r)
	require.Nil(t, err)
	wrapped.AssertGenericReply(t, wr)
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
func AssertWrappedVentureSliceFromReader(t *testing.T, r io.Reader) (wrapped.WrappedReply, []Venture) {
	wr, err := wrapped.DecodeFromReader(r)
	require.Nil(t, err)
	wrapped.AssertWrappedReply(t, wr)

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
func AssertWrappedVentureFromReader(t *testing.T, r io.Reader) (wrapped.WrappedReply, Venture) {
	wr, err := wrapped.DecodeFromReader(r)
	require.Nil(t, err)
	wrapped.AssertWrappedReply(t, wr)

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
