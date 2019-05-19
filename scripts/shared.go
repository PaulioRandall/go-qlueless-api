package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	comfiler "github.com/PaulioRandall/go-cookies/comfiler"
)

// clearTerminal clears the terminal
func clearTerminal() {
	p := runtime.GOOS
	switch p {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		panic("Platform '" + p + "' not currently supported")
	}
}

// makeBinDir makes a bin directory in the project root directory if it isn't
// already there
func makeBinDir(root string) {
	os.MkdirAll(root+"/bin", os.ModePerm)
}

// findProjectRoot returns the absolute path to the projects root
// directory
func findProjectRoot() string {
	fmt.Println("...finding project root...")

	scripts, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	scripts = strings.TrimSpace(scripts)
	root := filepath.Clean(scripts + "/..")
	fmt.Println("ok\t" + root)

	return root
}

// goFmt formats all Go code
func goFmt(root string) {
	fmt.Println("...formatting Go code...")
	goExe(".", []string{"fmt", root + "/..."})
}

// goOpenAPI builds the OpenAPI specification and places a copy of the
// resultant file in '{project-root}/bin'
func goOpenAPI(root string) {
	fmt.Println("...compiling OpenAPI specification...")

	api := root + "/api"
	oai := api + "/openapi"
	spec := oai + "/openapi.json"

	tmp := comfiler.Comfile{
		Template:  oai + "/template.json",
		Resources: oai + "/resources",
	}

	tmp.Compile(spec)
	fmt.Println("ok\t" + spec + "\t(created)")

	oaiBin := root + "/bin/openapi.json"
	copyFile(spec, oaiBin)
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

	output := root + "/bin/go-qlueless-api"
	args := []string{"build", "-o", output}
	args = append(args, goFiles...)
	goExe(cmd, args)

	fmt.Println("ok\t" + output + "\t(created)")
}

// goTest runs the application unit tests
func goTest(root string) {
	fmt.Println("...internal code testing...")
	goExe(".", []string{"test", root + "/cmd/..."})
	goExe(".", []string{"test", root + "/internal/..."})
}

// goTestApi runs the application API tests
func goTestApi(root string) {
	fmt.Println("...web API testing, this may take a few moments...")
	goExe(".", []string{"test", "-count=1", "-p=1", "-failfast", root + "/test/..."})
}

// goInstall install the compiled application
func goInstall(root string) {
	fmt.Println("...installing application...")

	cmd := root + "/cmd"
	goFiles, err := filepath.Glob(cmd + "/*.go")
	if err != nil {
		panic(err)
	}

	output := root + "/bin/go-qlueless-api"
	args := []string{"install"}
	args = append(args, goFiles...)
	goExe(cmd, args)

	fmt.Println("ok\t" + output + "\t(installed)")
}

// goRun runs the compiled application from the /bin directory
func goRun(root string) {
	fmt.Println("...running application...")
	cmd := exec.Command("./go-qlueless-api")
	cmd.Dir = root + "/bin"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
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
