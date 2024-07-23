package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-shutdown:
			fmt.Println("Shutting down")
			return
		case <-time.After(1 * time.Second):
			fmt.Printf("Test app output. version: %s, date: %s, commit: %s\n", version, date, commit)
			fmt.Printf("OS: %s\n", runtime.GOOS)
		}
	}
}
