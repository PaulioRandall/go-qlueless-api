package api

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func adminPrint(m string) {
	log.Println("[API-TEST] " + m)
}

func attemptRecover(exitCode *int) {
	if r := recover(); r != nil {
		adminPrint("Recovered from panic, exiting a little more gracefully")
		*exitCode = 1
	}
}

func startServer() *exec.Cmd {
	cmd := &exec.Cmd{
		Path:   "./go-qlueless-assembly-api",
		Dir:    "../../bin",
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	adminPrint("Starting server: " + cmd.Path)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	adminPrint("Pause to let server start")
	time.Sleep(500 * time.Millisecond)
	return cmd
}

func stopServer(cmd *exec.Cmd) {
	adminPrint("Killing server: " + cmd.Path)
	err := cmd.Process.Kill()
	if err != nil {
		log.Fatal("StopServer(): ", err)
	}
}
