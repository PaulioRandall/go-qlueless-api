package batches

import (
	"testing"

	shr "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

// When invoked, should not return nil
func TestLoadBatches___1(t *testing.T) {
	assert.NotNil(t, LoadBatches())
}

// When invoked, should return array of valid batches
func TestLoadBatches___2(t *testing.T) {
	var act []shr.WorkItem = LoadBatches()
	for _, b := range act {
		shr.CheckBatch(t, b)
	}
}
