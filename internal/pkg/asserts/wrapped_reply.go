package asserts

import (
	"testing"

	p "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

func AssertGenericError(t *testing.T, r p.WrappedReply) {
	assert.NotEmpty(t, r.Message)
	assert.NotEmpty(t, r.Self)
	assert.Empty(t, r.Data)
	assert.Empty(t, r.Hints)
}
