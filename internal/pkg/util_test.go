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

// When given a template with a single dynamic parameter, returns a map
// containing only the parameter
func TestPathVars___1(t *testing.T) {
	act, err := PathVars("/{id}", "/42")
	assert.Nil(t, err)
	assert.Len(t, act, 1)
	assert.Contains(t, act, "id")
	assert.Equal(t, "42", act["id"])
}

// When given a template with multiple dynamic parameters, returns a map
// containing the parameters
func TestPathVars___2(t *testing.T) {
	act, err := PathVars("/{abc}/{efg}", "/123/456")
	assert.Nil(t, err)
	assert.Len(t, act, 2)
	assert.Contains(t, act, "abc")
	assert.Contains(t, act, "efg")
	assert.Equal(t, "123", act["abc"])
	assert.Equal(t, "456", act["efg"])
}

// When given a tempalte with multiple dynamic parameters and static
// segments, returns a map containing the parameters with their values
func TestPathVars___3(t *testing.T) {
	act, err := PathVars("/master/{abc}/sub/{efg}", "/master/123/sub/456")
	assert.Nil(t, err)
	assert.Len(t, act, 2)
	assert.Contains(t, act, "abc")
	assert.Contains(t, act, "efg")
	assert.Equal(t, "123", act["abc"])
	assert.Equal(t, "456", act["efg"])
}

// When given a template and path with unequal number of segments, an error is
// returned
func TestPathVars___4(t *testing.T) {
	_, err := PathVars("/master/{abc}/sub", "/master/123")
	assert.NotNil(t, err)
}

// When given a template with an empty segment, an error is returned
func TestPathVars___5(t *testing.T) {
	_, err := PathVars("//{abc}", "//123")
	assert.NotNil(t, err)
}

// When given a template with an empty parameter, an error is returned
func TestPathVars___6(t *testing.T) {
	_, err := PathVars("/thing/{}", "/thing/123")
	assert.NotNil(t, err)
}

// When given a template and path where a static segment differs, an error is
// returned
func TestPathVars___7(t *testing.T) {
	_, err := PathVars("/master/{abc}", "/sub/123")
	assert.NotNil(t, err)
}

// When given a template with only static segments, an empty map is returned
func TestPathVars___8(t *testing.T) {
	act, err := PathVars("/master/sub", "/master/sub")
	assert.Nil(t, err)
	assert.Empty(t, act)
}

// When given a template with multiple instances of the same parameter, an error
// is returned
func TestPathVars___9(t *testing.T) {
	_, err := PathVars("/{abc}/{abc}", "/master/sub")
	assert.NotNil(t, err)
}
