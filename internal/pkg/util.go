package pkg

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
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

// pathVars_NewMatchErr is used by PathVars(..) to create an error detailing a
// mismatch between the template path and the request path
func pathVars_NewMatchErr(t string, p string) error {
	return errors.New(fmt.Sprintf("Template does not match path:"+
		" '%s' !~ '%s'", t, p))
}

// pathVars_AddParam is used by PathVars(..) to validate and add parameters to
// a resultant map
func pathVars_AddParam(m map[string]string, i int, t_seg string, p_seg string) error {
	if len(t_seg) == 2 {
		return errors.New(fmt.Sprintf("Empty name found for segment %d", i))
	}
	k := t_seg[1 : len(t_seg)-1]
	v := p_seg
	if _, exists := m[k]; exists {
		return errors.New(fmt.Sprintf("Multiple instances of the"+
			" parameter '%s' found", k))
	}
	m[k] = v
	return nil
}

// PathVars returns a map of all dynamic parameters within a URL path, templates
// must not use multiple instances of the same parameter key.
// Usage: 'PathVars("/things/{id}", "/things/42")' -> '{"id": "42"}'
// Usage: 'PathVars("/abc{things}/{id}", "/abc{things}/42")' -> '{"id": "42"}'
func PathVars(template string, path string) (map[string]string, error) {
	m := map[string]string{}
	t_parts := strings.Split(template, "/")[1:]
	p_parts := strings.Split(path, "/")[1:]

	if len(t_parts) != len(p_parts) {
		return nil, pathVars_NewMatchErr(template, path)
	}

	for i, seg := range t_parts {
		switch {
		case seg == "":
			return nil, errors.New(fmt.Sprintf("Empty segment in template at %d", i))
		case seg[0] == '{' && seg[len(seg)-1] == '}':
			err := pathVars_AddParam(m, i, seg, p_parts[i])
			if err != nil {
				return nil, err
			}
		case seg != p_parts[i]:
			return nil, pathVars_NewMatchErr(template, path)
		}
	}

	return m, nil
}
