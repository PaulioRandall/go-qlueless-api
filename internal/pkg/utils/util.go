package utils

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

const (
	POSITIVE_INT_CSV_PATTERN = "^([1-9][0-9]*,)*([1-9][0-9]*)$"
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

// AppendIfEmpty appends 'm' to 'r' if 's' is empty
func AppendIfEmpty(s string, r []string, m string) []string {
	if s == "" {
		return append(r, m)
	}
	return r
}

// AppendIfNotPositiveInt appends 'm' to 'r' if 's' is not a positive integer
func AppendIfNotPositiveInt(s string, r []string, m string) []string {
	i, err := strconv.Atoi(s)
	if err != nil || i < 1 {
		return append(r, m)
	}
	return r
}

// AppendIfNotPositiveIntCSV appends 'm' to 'r' if 's' is not a CSV of
// positive integers
func AppendIfNotPositiveIntCSV(s string, r []string, m string) []string {
	if IsPositiveIntCSV(s) {
		return r
	}
	return append(r, m)
}

// IsPositiveIntCSV returns true if the input is a CSV of positive integers
func IsPositiveIntCSV(s string) bool {
	match, _ := regexp.MatchString(POSITIVE_INT_CSV_PATTERN, s)
	return match
}

// DeleteIfExists deletes the file at the path specified if it exists
//
// @UNTESTED
func DeleteIfExists(path string) {
	err := os.Remove(path)
	switch {
	case err == nil, os.IsNotExist(err):
	default:
		log.Fatal(err)
	}
}
