package main

import (
	"log"
	"os"
	"strconv"
	"syscall"
)

var (
	pid = os.Getenv("RECEIVER_PID")
)

func main() {
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		log.Fatal(err)
	}

	p, err := os.FindProcess(pidInt)
	if err != nil {
		log.Fatalf("failed to find process with id: %v", pid)
	}
	log.Printf("sending interrupt to process: %v\n", pid)
	err = p.Signal(syscall.SIGTERM)
}
