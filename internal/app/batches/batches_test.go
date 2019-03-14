package batches

import (
	"testing"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

// When invoked, should not return nil
func TestLoad_batches___1(t *testing.T) {
	assert.NotNil(t, Load_batches())
}

// When invoked, should return array of valid batches
func TestLoad_batches___2(t *testing.T) {
	var act []shr.WorkItem = Load_batches()
	for _, b := range act {
		shr.CheckBatch(t, b)
	}
}
