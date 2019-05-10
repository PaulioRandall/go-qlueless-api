package main

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

// The Template struct represents a single substitution activity
type Template struct {
	Template  string // Path to the template file
	Resources string // Path to the root folder of injectable resource
	Output    string // Path to the output file
}

// Inject takes a filename that is relative to the Template instance and
// returns its content with each line indented with the specified number of tabs
func (o Template) Inject(filename string, indent int) string {
	path := o.Resources + filename
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	s := string(bytes)
	lines := strings.Split(s, "\n")
	prefix := strings.Repeat("\t", indent)

	for i, l := range lines {
		lines[i] = prefix + l
	}

	r := strings.Join(lines, "\n")
	return r
}

// compileOpenAPI performs the OpenAPI specification generation
func compileOpenAPI(o Template) error {
	var err error

	t, err := template.ParseFiles(o.Template)
	if err != nil {
		return err
	}

	f, err := os.Create(o.Output)
	if err != nil {
		return err
	}

	defer f.Close()
	err = t.Execute(f, o)
	if err != nil {
		return err
	}

	return nil
}
