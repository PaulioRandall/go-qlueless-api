///bin/true; exec /usr/bin/env go run "$0" "$@"

// Generates an openapi.json file from a template and fragment files
package main

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

// The OpenAPI struct represents a single template substitution activity
type OpenAPI struct {
	Template  string // Path to a template
	Resources string // Path to the root folder of injectable resource
}

// Check is a shorthand function for panic if err is not nil
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// FromFile takes a filename that is relative to the OpenAPI instance and
// returns its content with each indented with the supplied number of tabs
func (o OpenAPI) FromFile(filename string, indent int) string {
	path := o.Resources + filename
	bytes, err := ioutil.ReadFile(path)
	check(err)

	s := string(bytes)
	lines := strings.Split(s, "\n")
	prefix := strings.Repeat("\t", indent)

	for i, l := range lines {
		lines[i] = prefix + l
	}

	r := strings.Join(lines, "\n")
	return r
}

// Main is the entry point for the OpenAPI specification generator
func main() {
	var err error

	o := OpenAPI{
		Template:  "./template.json",
		Resources: "./resources",
	}

	t, err := template.ParseFiles(o.Template)
	check(err)

	f, err := os.Create("./openapi.json")
	check(err)

	defer f.Close()
	t.Execute(f, o)
}
