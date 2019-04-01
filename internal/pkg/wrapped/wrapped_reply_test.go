package wrapped

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ****************************************************************************
// DecodeWrappedReplyFromReader()
// ****************************************************************************

func TestDecodeWrappedReplyFromReader_1(t *testing.T) {
	aIn := strings.NewReader(`{
		"message": "message",
		"self": "/self",
		"data": {
			"k1": "v1",
			"k2": "v2"
		},
		"hints": "hints"
	}`)

	exp := WrappedReply{
		Message: "message",
		Self:    "/self",
		Data: map[string]interface{}{
			"k1": "v1",
			"k2": "v2",
		},
		Hints: "hints",
	}

	aOut, err := DecodeWrappedReplyFromReader(aIn)
	require.Nil(t, err)
	assert.Equal(t, exp, aOut)
}

func TestDecodeWrappedReplyFromReader_2(t *testing.T) {
	aIn := strings.NewReader(``)
	_, err := DecodeWrappedReplyFromReader(aIn)
	require.NotNil(t, err)
}

func TestDecodeWrappedReplyFromReader_3(t *testing.T) {
	aIn := strings.NewReader(`{}`)
	aOut, err := DecodeWrappedReplyFromReader(aIn)
	require.Nil(t, err)
	assert.Equal(t, WrappedReply{}, aOut)
}
