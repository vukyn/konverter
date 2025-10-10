package main

import (
	"konverter/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	app := server.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	log.Println("Waiting for shutdown signal...")
	<-sigChan

	server.Stop(app)

	log.Println("Server shutdown complete")
}
