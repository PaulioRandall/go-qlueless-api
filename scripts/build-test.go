//usr/bin/env go run "$0" "$@" "shared.go" "injector.go"; exit "$?"

package main

import (
	"fmt"
)

// main is the entry point for the script
func main() {
	clearTerminal()
	fmt.Println("[GOFMT -> BUILD -> TEST]")

	root := findProjectRoot()
	makeBinDir(root)

	goFmt(root)
	goOpenAPI(root)
	goBuild(root)
	goTest(root)
	goTestApi(root)

	return
}
