package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/antklim/go-calc"
)

func main() {
	errors := make(chan error, 1)
	go func() {
		errors <- calc.Start(":8080")
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errors:
		log.Fatalf("failed to start calc service: %+v", err)
	case <-osSignals:
		log.Println("shutting down calc service")
	}
}
