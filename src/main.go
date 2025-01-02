package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"reminder/src/app"
	"syscall"
)

func main() {
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
