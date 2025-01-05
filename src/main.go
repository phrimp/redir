package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"redir/src/app"
	"syscall"
)

func main() {
	// Check for command-line arguments
	args := os.Args[1:]
	if len(args) > 0 {
		// Handle commands sent to the running instance
		if err := app.HandleCommand(args); err != nil {
			fmt.Printf("Error handling command: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Run as a background service
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.StartApp(); err != nil {
			log.Fatalf("Error Running App: %v", err)
		}
	}()

	<-quit
	fmt.Println("Shutting down application...")
	app.Shutdown()
}
