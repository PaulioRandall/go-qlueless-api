package wrapped

import (
	"strings"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// ****************************************************************************
// DecodeFromReader()
// ****************************************************************************

func TestDecodeFromReader_1(t *testing.T) {
	aIn := strings.NewReader(`{
		"message": "message",
		"self": "/self",
		"data": {
			"k1": "v1",
			"k2": "v2"
		}
	}`)

	exp := WrappedReply{
		Message: "message",
		Self:    "/self",
		Data: map[string]interface{}{
			"k1": "v1",
			"k2": "v2",
		},
	}

	aOut, err := DecodeFromReader(aIn)
	require.Nil(t, err)
	assert.Equal(t, exp, aOut)
}

func TestDecodeFromReader_2(t *testing.T) {
	aIn := strings.NewReader(``)
	_, err := DecodeFromReader(aIn)
	require.NotNil(t, err)
}

func TestDecodeFromReader_3(t *testing.T) {
	aIn := strings.NewReader(`{}`)
	aOut, err := DecodeFromReader(aIn)
	require.Nil(t, err)
	assert.Equal(t, WrappedReply{}, aOut)
}
