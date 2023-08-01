package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func createLogFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

var Logger *log.Logger

func init() {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working directory:", err)
	}

	// Get the date
	date := time.Now().Format("2006-01-02")

	// Get hours, minutes and seconds
	// time := time.Now().Format("15-04-05")

	// Create the log file
	filePath := filepath.Join(dir, "logs", date + ".log")

	// If file doesn't exist, create it, or append to the file
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to create log file:", err)
	}

	// Initialize the logger
	Logger = log.New(file, "\nxCall Event Watcher: ", log.LstdFlags)
	// If you want to log to the console, use this line instead:
	// logger = log.New(os.Stdout, "MyApp: ", log.LstdFlags)
}
