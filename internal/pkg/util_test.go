package pkg

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// When given a valid int string, true is returned
func TestIsInt___1(t *testing.T) {
	assert.True(t, IsInt("123"))
}

// When given an invalid int string, false is returned
func TestIsInt___2(t *testing.T) {
	assert.False(t, IsInt("abc"))
}

// When given an invalid int string, false is returned
func TestIsInt___3(t *testing.T) {
	assert.False(t, IsInt("123abc"))
}

// When given an empty string, true is returned
func TestIsBlank___1(t *testing.T) {
	act := IsBlank("")
	assert.True(t, act)
}

// When given a string with whitespace, true is returned
func TestIsBlank___2(t *testing.T) {
	act := IsBlank("\r\n \t\f")
	assert.True(t, act)
}

// When given a string with no whitespaces, false is returned
func TestIsBlank___3(t *testing.T) {
	act := IsBlank("Captain Vimes")
	assert.False(t, act)
}

// When given nil, returns false and doesn't print anything
func TestLogIfErr___1(t *testing.T) {
	act := LogIfErr(nil)
	assert.False(t, act)
	// Output:
	//
}

// When given an error, returns true and prints the error message
func TestLogIfErr___2(t *testing.T) {
	err := errors.New("Computer says no!")
	act := LogIfErr(err)
	assert.True(t, act)
	// Output:
	// Computer says no!
}

// When given a string containing a variation of whitespace characters in
// between non-whitespace characters, the whitespace is removed
func TestStripWhitespace___1(t *testing.T) {
	act := StripWhitespace("Rince \n\t\f\r wind")
	assert.Equal(t, "Rincewind", act)
}

// When given a string containing a variation of whitespace characters
// leading non-whitespace characters, the whitespace is removed
func TestStripWhitespace___2(t *testing.T) {
	act := StripWhitespace("\t \n\t \r\n\n\fRincewind")
	assert.Equal(t, "Rincewind", act)
}

// When given a string containing a variation of whitespace characters
// trailing non-whitespace characters, the whitespace is removed
func TestStripWhitespace___3(t *testing.T) {
	act := StripWhitespace("Rincewind\r\n \t\t\t\t \f \r  \v\v")
	assert.Equal(t, "Rincewind", act)
}

// When given a string containing a variation of non-whitespace characters
// interspread between whitespace characters, the whitespace is removed
func TestStripWhitespace___4(t *testing.T) {
	act := StripWhitespace("\r\nRi \tn\tc\t\t ew\f \r  in\vd\v")
	assert.Equal(t, "Rincewind", act)
}
