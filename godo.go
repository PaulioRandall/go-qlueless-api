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

	comfiler "github.com/PaulioRandall/go-cookies/comfiler"
)

// main is the entry point for this script. Use this script to perform the
// standard Go format, build, test, run, and install operations without
// needing to know each operations arguments.
func main() {
	clearTerminal()

	root := findProjectRoot()
	makeBinDir(root)

	args := os.Args[1:]
	if len(args) != 1 {
		badSyntax()
	}

	// Don't abstract the build processes! They are more readable and extendable
	// this way.
	switch args[0] {
	case "build":
		fmt.Println("[GOFMT -> BUILD -> TEST]")

		goFmt(root)
		goOpenAPI(root)
		goBuild(root)
		goTest(root)
		goTestApi(root)

	case "run":
		fmt.Println("[GOFMT -> BUILD -> TEST -> RUN]")

		goFmt(root)
		goOpenAPI(root)
		goBuild(root)
		goTest(root)
		goTestApi(root)
		goRun(root)

	case "install":
		fmt.Println("[GOFMT -> BUILD -> TEST -> INSTALL]")

		goFmt(root)
		goOpenAPI(root)
		goBuild(root)
		goTest(root)
		goTestApi(root)
		goInstall(root)

	default:
		badSyntax()
	}
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

// findProjectRoot returns the absolute path to the projects root directory.
func findProjectRoot() string {
	fmt.Println("...finding project root...")

	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	root = strings.TrimSpace(root)
	root = filepath.Clean(root)

	fmt.Println("ok\t" + root)
	return root
}

// makeBinDir creates a /bin directory in the 'root' directory if it doesn't
// already exist.
func makeBinDir(root string) {
	err := os.Mkdir(root+"/bin", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}

// goFmt recursively formats all Go code within the 'root' directory.
func goFmt(root string) {
	fmt.Println("...formatting Go code...")
	goExe(root, []string{"fmt", root + "/..."})
}

// goOpenAPI builds the OpenAPI specification and places a copy of the
// resultant file in 'root/bin'.
func goOpenAPI(root string) {
	fmt.Println("...compiling OpenAPI specification...")

	api := root + "/api"
	spec := root + "/bin/openapi.json"

	tmp := comfiler.Comfile{
		Template:  api + "/oai-template.json",
		Resources: api,
	}

	err := tmp.Compile(spec)
	if err != nil {
		panic(err)
	}

	fmt.Println("ok\t" + spec + "\t(created)")

	cl := root + "/CHANGELOG.md"
	clBin := root + "/bin/CHANGELOG.md"
	copyFile(cl, clBin)
	fmt.Println("ok\t" + clBin + "\t(copied)")
}

// goBuild builds the application and places the result binary in 'root/bin'.
func goBuild(root string) {
	fmt.Println("...building application...")

	api := root + "/api"
	goFiles, err := filepath.Glob(api + "/*.go")
	if err != nil {
		panic(err)
	}

	output := root + "/bin/go-qlueless-api"
	args := []string{"build", "-o", output}
	args = append(args, goFiles...)
	goExe(api, args)

	fmt.Println("ok\t" + output + "\t(created)")
}

// goTest runs the applications internal tests (unit).
func goTest(root string) {
	fmt.Println("...internal code testing...")
	goExe(root, []string{"test", root + "/api/..."})
	goExe(root, []string{"test", root + "/shared/..."})
}

// goTestApi runs the applications API tests.
func goTestApi(root string) {
	fmt.Println("...web API testing, this may take a few moments...")
	goExe(root, []string{"test", "-count=1", "-p=1", "-failfast", root + "/test/..."})
}

// goInstall installs the compiled application.
func goInstall(root string) {
	fmt.Println("...installing application...")

	api := root + "/api"
	goFiles, err := filepath.Glob(api + "/*.go")
	if err != nil {
		panic(err)
	}

	output := root + "/bin/go-qlueless-api"
	args := []string{"install"}
	args = append(args, goFiles...)
	goExe(api, args)

	fmt.Println("ok\t" + output + "\t(installed)")
}

// goRun runs the compiled application from the /bin directory.
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
