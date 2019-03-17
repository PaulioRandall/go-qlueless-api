package batches

import (
	"testing"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

// When invoked, should not return nil
func TestLoadBatches___1(t *testing.T) {
	assert.NotNil(t, LoadBatches())
}

// When invoked, should return array of valid batches
func TestLoadBatches___2(t *testing.T) {
	act := LoadBatches()
	for _, b := range act {
		CheckBatch(t, b)
	}
}
