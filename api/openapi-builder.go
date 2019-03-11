///bin/true; exec /usr/bin/env go run "$0" "$@"
package main

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type OpenAPI struct {
	Template  string
	Resources string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (o OpenAPI) FromFile(filename string, indent int) string {
	path := o.Resources + filename
	bytes, err := ioutil.ReadFile(path)
	check(err)

	s := string(bytes)
	lines := strings.Split(s, "\n")

	for i, l := range lines {
		lines[i] = strings.Repeat("\t", indent) + l
	}

	r := strings.Join(lines, "\n")
	return r
}

func main() {
	var err error

	o := OpenAPI{
		Template:  "./openapi-template.json",
		Resources: "./openapi-resources",
	}
	t, err := template.ParseFiles(o.Template)
	check(err)

	f, err := os.Create("./openapi.json")
	check(err)
	defer f.Close()

	t.Execute(f, o)
}
