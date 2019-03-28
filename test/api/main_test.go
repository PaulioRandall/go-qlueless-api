package api

import (
	"fmt"
	"os"
	"testing"
)

const (
	CORS_METHODS_PATTERN = "^((\\s*[A-Z]*\\s*,)+)*(\\s*[A-Z]*\\s*)$" // Example: 'GET, POST   ,OPTIONS'
)

func exit(exitCode *int) {
	os.Exit(*exitCode)
}

func TestMain(m *testing.M) {
	var exitCode int = 1
	cmd := startServer()

	defer exit(&exitCode)
	defer attemptRecover(&exitCode)
	defer stopServer(cmd)

	exitCode = m.Run()
	adminPrint(fmt.Sprintf("Test exit code: %d", exitCode))
}
