package logger

import (
	"log"
	"os"
	"path/filepath"
)

var logger *log.Logger
var logChan chan string

func init() {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working directory:", err)
	}

	// Create the log file
	filePath := filepath.Join(dir, "logs", "transactions.log")

	// If file doesn't exist, create it, or append to the file
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to create log file:", err)
	}

	// Initialize the logger
	logger = log.New(file, "xCall Event Watcher: ", log.LstdFlags)
	// If you want to log to the console, use this line instead:
	// logger = log.New(os.Stdout, "MyApp: ", log.LstdFlags)

	logChan = make(chan string, 100) // Buffer size 100 (adjust based on your needs)
	go logHandler()
}

func logHandler() {
	for msg := range logChan {
		logger.Println(msg)
	}
}

func LogMessage(msg string) {
	logChan <- msg
}