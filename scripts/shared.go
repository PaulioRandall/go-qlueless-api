package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// makeBinDir makes a bin directory in the project root directory if it isn't
// already there
func makeBinDir(root string) {
	os.MkdirAll(root+"/bin", os.ModePerm)
}

// findProjectRoot returns the absolute path to the projects root
// directory
func findProjectRoot() string {
	scripts, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	scripts = strings.TrimSpace(scripts)
	root := filepath.Clean(scripts + "/..")
	fmt.Println("...found project root...")
	fmt.Println("ok\t" + root)

	return root
}

// goOpenAPI builds the OpenAPI specification and places a copy of the
// resultant file in '{project-root}/bin'
func goOpenAPI(root string) {
	fmt.Println("...compiling OpenAPI specification...")

	api := root + "/api"
	oai := api + "/openapi"
	goExe(oai, []string{"run", oai + "/injector.go"})

	oaiOut := oai + "/openapi.json"
	fmt.Println("ok\t" + oaiOut + "\t(created)")

	oaiBin := root + "/bin/openapi.json"
	copyFile(oaiOut, oaiBin)
	fmt.Println("ok\t" + oaiBin + "\t(copied)")

	cl := api + "/CHANGELOG.md"
	clBin := root + "/bin/CHANGELOG.md"
	copyFile(cl, clBin)
	fmt.Println("ok\t" + clBin + "\t(copied)")
}

// goBuild builds the application and places the result binary in
// '{project-root}/bin'
func goBuild(root string) {
	fmt.Println("...building application...")

	cmd := root + "/cmd"
	goFiles, err := filepath.Glob(cmd + "/*.go")
	if err != nil {
		panic(err)
	}

	output := root + "/bin/go-qlueless-assembly-api"
	args := []string{"build", "-o", output}
	args = append(args, goFiles...)
	goExe(cmd, args)

	fmt.Println("ok\t" + output + "\t(created)")
}

// goTest runs the application unit tests
func goTest(root string) {
	fmt.Println("...testing application...")
	goExe(".", []string{"test", root + "/cmd/..."})
	goExe(".", []string{"test", root + "/internal/..."})
}

// goTestApi runs the application API tests
func goTestApi(root string) {
	fmt.Println("...testing API, this may take a few moments...")
	goExe(".", []string{"test", "-count=1", root + "/test/..."})
}

// goExe runs a Go command
func goExe(dir string, args []string) {
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// copyFile copies a file from one location to another
func copyFile(src string, dst string) {
	in, err := ioutil.ReadFile(src)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(dst, in, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
