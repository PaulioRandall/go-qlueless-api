package orders

import (
	"testing"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

// When invoked, should not return nil
func TestLoad_orders___1(t *testing.T) {
	assert.NotNil(t, Load_orders())
}

// When invoked, should return array of valid orders
func TestLoad_orders___2(t *testing.T) {
	var act []shr.WorkItem = Load_orders()
	for _, o := range act {
		shr.CheckOrder(t, o)
	}
}
