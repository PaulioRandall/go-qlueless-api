package test

import (
	"fmt"
	"os"
	"testing"
)

// _exit is file private function that accepts a referenced integer instead of
// an actual one
func _exit(exitCode *int) {
	os.Exit(*exitCode)
}

// TestMain is the entry point for the 'go test'
func TestMain(m *testing.M) {
	var exitCode int = 1

	defer _exit(&exitCode)
	defer attemptRecover(&exitCode)

	exitCode = m.Run()
	adminPrint(fmt.Sprintf("Test exit code: %d", exitCode))
}
