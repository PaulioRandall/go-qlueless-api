package pkg

import (
	"bytes"
	"log"
	"strconv"
	"unicode"
)

// IsInt returns true if the string contains an integer
func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// IsBlank returns true if the string is empty or only contains whitespace
func IsBlank(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// LogIfErr checks if the input err is NOT nil returning true if it is.
// When true the error is logged so all the calling handler needs to do is
// clean up then invoke Http_500(*http.ResponseWriter) before returning
func LogIfErr(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

// StripWhitespace removes all white space from a string
func StripWhitespace(s string) string {
	var buf bytes.Buffer
	for _, r := range s {
		if !unicode.IsSpace(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
