package ventures

import (
	"io"
	"testing"

	p "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	ts "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	ms "github.com/mitchellh/mapstructure"
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
	assert.NotEmpty(t, v.ID, "Venture.ID")
	assert.NotEmpty(t, v.Description, "Venture.Description")
	assert.NotEmpty(t, v.State, "Venture.State")
	if v.OrderIDs != "" {
		ts.AssertGenericIntCSV(t, v.OrderIDs)
	}
	assert.True(t, v.IsAlive, "Venture.IsAlive")
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

func AssertWrappedVentureSliceFromReader(t *testing.T, r io.Reader) p.WrappedReply {
	wr, err := p.DecodeWrappedReplyFromReader(r)
	require.Nil(t, err)

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
	return wr
}
