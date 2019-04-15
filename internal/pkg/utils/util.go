package utils

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
	"unicode"
)

const (
	POSITIVE_INT_CSV_PATTERN = "^([1-9][0-9]*,)*([1-9][0-9]*)$"
)

// LogIfErr checks if the input error is NOT nil. When true, the error is logged
// and the check result is returned.
func LogIfErr(err error) bool {
	if err != nil {
		log.Println("[ERROR] " + err.Error())
		return true
	}
	return false
}

// WarnIfErr checks if the input error is NOT nil. When true, the error is
// logged as a warning and the check result is returned.
func WarnIfErr(err error) bool {
	if err != nil {
		log.Println("[WARNING] " + err.Error())
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

// UnixMilliNow returns the current time as Unix milliseconds
//
// @UNTESTED
func UnixMilliNow() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
