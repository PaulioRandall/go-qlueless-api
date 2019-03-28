package ventures

import (
	"testing"

	ts "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	"github.com/stretchr/testify/assert"
)

func AssertGenericVenture(t *testing.T, v Venture) {
	assert.NotEmpty(t, v.ID)
	assert.NotEmpty(t, v.Description)
	assert.NotEmpty(t, v.State)
	if v.OrderIDs != "" {
		ts.AssertGenericIntCSV(t, v.OrderIDs)
	}
	assert.True(t, v.IsAlive)
}

func AssertGenericVentures(t *testing.T, vens []Venture) {
	for _, v := range vens {
		AssertGenericVenture(t, v)
	}
}
