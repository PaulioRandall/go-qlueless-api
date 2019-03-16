package orders

import (
	"testing"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

// When invoked, should not return nil
func TestLoadOrders___1(t *testing.T) {
	assert.NotNil(t, LoadOrders())
}

// When invoked, should return array of valid orders
func TestLoadOrders___2(t *testing.T) {
	var act []WorkItem = LoadOrders()
	for _, o := range act {
		CheckOrder(t, o)
	}
}
