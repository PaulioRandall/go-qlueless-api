package pkg

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

// When given an empty string, the message is appended
func TestAppendIfEmpty___1(t *testing.T) {
	act := AppendIfEmpty("", []string{}, "abc")
	require.Len(t, act, 1)
	assert.Contains(t, act, "abc")
}

// When given an empty string, the message is appended
func TestAppendIfEmpty___2(t *testing.T) {
	act := AppendIfEmpty("", []string{"xyz"}, "abc")
	require.Len(t, act, 2)
	assert.Contains(t, act, "xyz")
	assert.Contains(t, act, "abc")
}

// When given a non-empty string, no appending occurs
func TestAppendIfEmpty___3(t *testing.T) {
	act := AppendIfEmpty("NOT-EMPTY", []string{}, "abc")
	assert.Len(t, act, 0)
}

// When given a positive integer, the message is NOT appended
func TestAppendIfNotPositiveInt___1(t *testing.T) {
	act := AppendIfNotPositiveInt("5", []string{}, "abc")
	assert.Len(t, act, 0)
}

// When given zero, the message is appended
func TestAppendIfNotPositiveInt___2(t *testing.T) {
	act := AppendIfNotPositiveInt("0", []string{}, "abc")
	require.Len(t, act, 1)
	assert.Contains(t, act, "abc")
}

// When given a negative number, the message is appended
func TestAppendIfNotPositiveInt___3(t *testing.T) {
	act := AppendIfNotPositiveInt("-5", []string{}, "abc")
	require.Len(t, act, 1)
	assert.Contains(t, act, "abc")
}

// When given a multiple non-positive inputs, all messages are appended
func TestAppendIfNotPositiveInt___4(t *testing.T) {
	act := []string{}
	act = AppendIfNotPositiveInt("-1", act, "abc")
	act = AppendIfNotPositiveInt("-1", act, "efg")
	act = AppendIfNotPositiveInt("-1", act, "hij")
	require.Len(t, act, 3)
	assert.Contains(t, act, "abc")
	assert.Contains(t, act, "efg")
	assert.Contains(t, act, "hij")
}

func TestAppendIfNotPositiveIntCSV_1(t *testing.T) {
	act := AppendIfNotPositiveIntCSV("1,2,99,4,3", []string{}, "abc")
	assert.Len(t, act, 0)
}

func TestAppendIfNotPositiveIntCSV_2(t *testing.T) {
	act := AppendIfNotPositiveIntCSV("4", []string{}, "abc")
	assert.Len(t, act, 0)
}

func TestAppendIfNotPositiveIntCSV_3(t *testing.T) {
	act := AppendIfNotPositiveIntCSV("0", []string{}, "abc")
	assert.Len(t, act, 1)
	assert.Equal(t, "abc", act[0])
}

func TestAppendIfNotPositiveIntCSV_4(t *testing.T) {
	act := AppendIfNotPositiveIntCSV("-99", []string{}, "abc")
	assert.Len(t, act, 1)
	assert.Equal(t, "abc", act[0])
}

func TestAppendIfNotPositiveIntCSV_5(t *testing.T) {
	act := AppendIfNotPositiveIntCSV("3,2,1,0", []string{}, "abc")
	assert.Len(t, act, 1)
	assert.Equal(t, "abc", act[0])
}

func TestAppendIfNotPositiveIntCSV_6(t *testing.T) {
	act := AppendIfNotPositiveIntCSV(",1,2", []string{}, "abc")
	assert.Len(t, act, 1)
	assert.Equal(t, "abc", act[0])
}

func TestAppendIfNotPositiveIntCSV_7(t *testing.T) {
	act := AppendIfNotPositiveIntCSV("", []string{}, "abc")
	assert.Len(t, act, 1)
	assert.Equal(t, "abc", act[0])
}
