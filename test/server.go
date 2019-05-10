package test

import (
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

var cmd *exec.Cmd = nil
var longPause sync.Once

// adminPrint prints a general test message distinguishable from ordinary test
// log messages
func adminPrint(m string) {
	log.Println("[API-TEST] " + m)
}

// attemptRecover attempts to recover from a panic setting the exit code to 1
// if recovery from a panic was made
func attemptRecover(exitCode *int) {
	if r := recover(); r != nil {
		adminPrint("Recovered from panic, exiting a little more gracefully")
		*exitCode = 1
	}
}

// StartServer starts the application server
func StartServer(binPath string) {
	if binPath == "" {
		binPath = "../bin"
	}

	cmd = &exec.Cmd{
		Path:   "./go-qlueless-api",
		Dir:    binPath,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	adminPrint("Starting server: " + cmd.Path)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	adminPrint("Pause to let server start")
	longPause.Do(func() {
		time.Sleep(300 * time.Millisecond)
	})
	time.Sleep(75 * time.Millisecond)
}

// StopServer stops the application server
func StopServer() {
	adminPrint("Killing server: " + cmd.Path)
	err := cmd.Process.Kill()
	if err != nil {
		log.Fatal("StopServer(): ", err)
	}
	cmd = nil
}
