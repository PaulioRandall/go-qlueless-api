package pkg

import (
	"errors"
	"fmt"
	"testing"
)

func fail(t *testing.T, name string, exp interface{}, act interface{}) {
	var m = fmt.Sprintf("%s... Expected: %v... Actual: %v", name, exp, act)
	t.Errorf(m)
}

func TestLog_if_err___1(t *testing.T) {
	act := Log_if_err(nil)
	if act {
		fail(t, "Log_if_err(nil)", false, act)
	}
	// Output:
	//
}

func TestLog_if_err___2(t *testing.T) {
	var err error = errors.New("Computer says no!")
	act := Log_if_err(err)
	if !act {
		fail(t, "Log_if_err(err)", true, act)
	}
	// Output:
	// Computer says no!
}
