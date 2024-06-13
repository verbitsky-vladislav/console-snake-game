package utils

import (
	"log"
	"os"
)

var (
	LogInfo  *log.Logger
	LogError *log.Logger
)

func init() {
	file, err := os.OpenFile("snake-game.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	LogInfo = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	LogError = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
