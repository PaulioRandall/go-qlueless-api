//usr/bin/env go run "$0" "$@" "shared.go"; exit "$?"

package main

import (
	"fmt"
)

// main is the entry point for the script
func main() {
	clearTerminal()
	fmt.Println("[GOFMT -> BUILD -> TEST -> INSTALL]")

	root := findProjectRoot()
	makeBinDir(root)

	goFmt(root)
	goOpenAPI(root)
	goBuild(root)
	goTest(root)
	goTestApi(root)
	goInstall(root)

	return
}
