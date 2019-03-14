package pkg

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog_if_err___1(t *testing.T) {
	act := Log_if_err(nil)
	assert.False(t, act)
	// Output:
	//
}

func TestLog_if_err___2(t *testing.T) {
	var err error = errors.New("Computer says no!")
	act := Log_if_err(err)
	assert.True(t, act)
	// Output:
	// Computer says no!
}
