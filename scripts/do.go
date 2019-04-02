//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// main is the entry point for the application
func main() {

	root := _findProjectRoot()
	os.MkdirAll(root+"/bin", os.ModePerm)

	_goBuild(root)
	return
}

// _findProjectRoot returns the absolute path to the projects root
// directory
func _findProjectRoot() string {
	scripts, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	scripts = strings.TrimSpace(scripts)
	root := filepath.Clean(scripts + "/..")
	fmt.Println("[root] " + root)

	return root
}

// _goBuild builds the application and places the result binary
// in '{project-root}/bin'
func _goBuild(root string) {
	cmd := root + "/cmd"
	goFiles, err := filepath.Glob(cmd + "/*.go")
	if err != nil {
		panic(err)
	}

	args := []string{
		"build",
		"-o",
		root + "/bin/go-qlueless-assembly-api",
	}
	args = append(args, goFiles...)
	_go(cmd, args)
}

// _go runs a Go command
func _go(dir string, args []string) {
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
