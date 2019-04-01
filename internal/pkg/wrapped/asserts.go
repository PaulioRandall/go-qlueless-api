package wrapped

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

// AssertGenericError checks that an error response body has the required fields
// populated
func AssertGenericError(t *testing.T, r WrappedReply) {
	assert.NotEmpty(t, r.Message)
	assert.NotEmpty(t, r.Self)
	assert.Empty(t, r.Data)
}

// AssertWrappedReply checks that a wrapped response body has the required
// fields populated
func AssertWrappedReply(t *testing.T, r WrappedReply) {
	assert.NotEmpty(t, r.Message)
	assert.NotEmpty(t, r.Self)
	assert.NotEmpty(t, r.Data)
}
