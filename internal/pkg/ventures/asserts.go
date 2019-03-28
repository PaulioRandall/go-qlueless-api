package ventures

import (
	"testing"

	ts "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/asserts"
	"github.com/stretchr/testify/assert"
)

func AssertGenericVenture(t *testing.T, ven Venture) {
	assert.Equal(t, ven.VentureID, "1")
	assert.NotEmpty(t, ven.Description)
	assert.NotEmpty(t, ven.State)
	if ven.OrderIDs != "" {
		ts.AssertGenericIntCSV(t, ven.OrderIDs)
	}
}

func AssertGenericVentures(t *testing.T, vens []Venture) {
	for _, v := range vens {
		AssertGenericVenture(t, v)
	}
}
