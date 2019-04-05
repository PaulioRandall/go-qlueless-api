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

// startServer starts the application server
func startServer() {
	cmd = &exec.Cmd{
		Path:   "./go-qlueless-assembly-api",
		Dir:    "../bin",
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
		time.Sleep(250 * time.Millisecond)
	})
	time.Sleep(50 * time.Millisecond)
}

// stopServer stops the application server
func stopServer() {
	adminPrint("Killing server: " + cmd.Path)
	err := cmd.Process.Kill()
	if err != nil {
		log.Fatal("StopServer(): ", err)
	}
	cmd = nil
}
