//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	comfiler "github.com/PaulioRandall/go-cookies/comfiler"
	cookies "github.com/PaulioRandall/go-cookies/cookies"
)

// main is the entry point for this script. It wraps the standard Go format,
// build, test, run, and install operations specifically for this project.
func main() {
	clearTerminal()
	printTime()

	started := cookies.ToUnixMilli(time.Now().UTC())
	root := getwd()
	makeBinDir(root)

	// Don't abstract the build workflows! They are more readable and extendable
	// this way.
	switch getArgument() {
	case "build":
		goFmt(root)
		goOpenAPI(root)
		goBuild(root)
		goTest(root)
		goTestApi(root)

	case "run":
		goFmt(root)
		goOpenAPI(root)
		goBuild(root)
		goTest(root)
		goTestApi(root)
		goRun(root)

	case "install":
		goFmt(root)
		goOpenAPI(root)
		goBuild(root)
		goTest(root)
		goTestApi(root)
		goInstall(root)

	default:
		badSyntax()
	}

	finished := cookies.ToUnixMilli(time.Now())
	printTime()
	printRunTime(started, finished)
}

// clearTerminal clears the terminal.
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

// printTime prints the current time to terminal.
func printTime() {
	fmt.Printf("Now\t%v\n", time.Now().UTC())
}

// printOk prints an OK message in the style of 'go test'.
func printOk(url string, results ...string) {
	if len(results) == 0 {
		fmt.Printf("ok\t%s\n", url)
		return
	}

	r := strings.Join(results, ", ")
	fmt.Printf("ok\t%s\t(%s)\n", url, r)
}

// getwd returns the absolute path to the projects root directory.
func getwd() string {
	fmt.Println("...finding project root...")

	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	printOk(root)
	return root
}

// makeBinDir creates a /bin directory in the 'root' directory if it doesn't
// already exist.
func makeBinDir(root string) {
	bin := filepath.Join(root, "bin")
	err := os.Mkdir(bin, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}

// getArgument returns the argument passed that represents the operation to
// perform.
func getArgument() string {
	args := os.Args[1:]
	if len(args) != 1 {
		badSyntax()
	}
	return args[0]
}

// badSyntax prints the scripts syntax to console then exits the application
// with code 1.
func badSyntax() {
	syntax := `syntax options:
1) ./godo.go build  		Builds and tests
2) ./godo.go run    		Builds, tests, and runs
3) ./godo.go install		Builds, tests, and installs`

	fmt.Println(syntax + "\n")
	os.Exit(1)
}

// goFmt recursively formats all Go code within the 'root' directory.
func goFmt(root string) {
	fmt.Println("...formatting Go code...")
	target := filepath.Join(root, "...")
	goExe(root, "fmt", target)
}

// goOpenAPI builds the OpenAPI specification and places a copy of the
// resultant file in 'root/bin'.
func goOpenAPI(root string) {
	fmt.Println("...compiling OpenAPI specification...")

	api := filepath.Join(root, "api")
	template := filepath.Join(api, "oai-template.json")
	output := filepath.Join(root, "bin", "openapi.json")

	tmp := comfiler.Comfile{
		Template:  template,
		Resources: api,
	}

	err := tmp.Compile(output)
	if err != nil {
		panic(err)
	}
	printOk(output, "created")

	cl := filepath.Join(root, "CHANGELOG.md")
	clBin := filepath.Join(root, "bin", "CHANGELOG.md")
	copyFile(cl, clBin)
	printOk(clBin, "copied")
}

// goBuild builds the application and places the result binary in 'root/bin'.
func goBuild(root string) {
	fmt.Println("...building application...")

	api := filepath.Join(root, "api")
	goFiles := globGoFiles(api)

	output := filepath.Join(root, "bin", "go-qlueless-api")
	args := []string{"build", "-o", output}
	args = append(args, goFiles...)
	goExe(api, args...)

	printOk(output, "created")
}

// globGoFiles returns the names of all Go files in the API directory.
func globGoFiles(apiDir string) []string {
	glob := filepath.Join(apiDir, "*.go")
	goFiles, err := filepath.Glob(glob)
	if err != nil {
		panic(err)
	}
	return goFiles
}

// goTest runs the applications internal tests (unit).
func goTest(root string) {
	fmt.Println("...internal code testing...")

	api := filepath.Join(root, "api", "...")
	goExe(root, "test", api)

	shared := filepath.Join(root, "shared", "...")
	goExe(root, "test", shared)
}

// goTestApi runs the applications API tests.
func goTestApi(root string) {
	fmt.Println("...web API testing, this may take a few moments...")
	tests := filepath.Join(root, "test", "...")
	goExe(root, "test", "-count=1", "-p=1", "-failfast", tests)
}

// goInstall installs the compiled application.
func goInstall(root string) {
	fmt.Println("...installing application...")

	api := filepath.Join(root, "api")
	goFiles := globGoFiles(api)

	output := filepath.Join(root, "bin", "go-qlueless-api")
	args := append([]string{"install"}, goFiles...)
	goExe(api, args...)

	printOk(output, "installed")
}

// goRun runs the compiled application from the /bin directory.
func goRun(root string) {
	fmt.Println("...running application...")
	cmd := exec.Command("go-qlueless-api")
	cmd.Dir = filepath.Join(root, "bin")
	runCmd(cmd)
}

// runCmd runs a command and panics on error.
func runCmd(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if _, ok := err.(*exec.ExitError); ok {
		panic(err.Error())
	} else if err != nil {
		panic(err)
	}
}

// goExe runs a Go command
func goExe(dir string, args ...string) {
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	runCmd(cmd)
}

// copyFile copies a file from 'src' to 'dst'.
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

// printRunTime prints the time taken for specific godo command to fully run.
func printRunTime(start, finish int64) {
	ms := finish - start
	ns := ms * int64(time.Millisecond)
	secs := float64(ns) / float64(time.Second)
	fmt.Printf("Stats\t%.2f seconds\n", secs)
}
