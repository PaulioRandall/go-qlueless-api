//usr/bin/env go run "$0" "$@" "shared.go"; exit "$?"

package main

import (
	"fmt"
)

// main is the entry point for the application
func main() {
	fmt.Println("[BUILD -> TEST -> API]")
	root := findProjectRoot()
	makeBinDir(root)

	goBuild(root)
	goTest(root)
	goTestApi(root)
	return
}
