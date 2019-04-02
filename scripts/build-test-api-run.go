//usr/bin/env go run "$0" "$@" "shared.go"; exit "$?"

package main

import (
	"fmt"
)

// main is the entry point for the script
func main() {
	fmt.Println("[BUILD -> TEST -> API -> RUN]")

	root := findProjectRoot()
	makeBinDir(root)

	goOpenAPI(root)
	goBuild(root)
	goTest(root)
	goTestApi(root)
	goRun(root)

	return
}
