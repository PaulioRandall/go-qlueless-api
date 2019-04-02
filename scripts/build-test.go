//usr/bin/env go run "$0" "$@" "shared.go"; exit "$?"

package main

import (
	"fmt"
)

// main is the entry point for the application
func main() {
	fmt.Println("[BUILD -> TEST]")
	root := findProjectRoot()
	makeBinDir(root)

	goOpenAPI(root)
	goBuild(root)
	goTest(root)

	return
}
