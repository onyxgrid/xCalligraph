package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/paulrouge/xCalligraph/internal/evm"
	"github.com/paulrouge/xCalligraph/internal/icon"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Berlin := os.Getenv("BERLIN")
	Sepolia := os.Getenv("SEPOLIA")

	if Berlin != "true" && Sepolia != "true" {
		panic("Berlin nor Sepolia is set to true. Please check your .env file.")
	}

	// Start go routines for checking the Berlin chain.
	if Berlin == "true" {
		go icon.CallExecuteCall()
		go icon.CheckBlocks()
		go icon.HandleBlock()
		go icon.HandleTransaction()
	}

	// Start go routines for checking the Sepolia chain.
	if Sepolia == "true" {
		go evm.CallExecuteCall()
		go evm.CheckBlocks()
		go evm.HandleBlock()
		go evm.HandleTransaction()
	}

	// Keep the main thread alive.
	select {}
}
