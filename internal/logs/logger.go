package logs

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	// Open/Create file to store logs
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	// Create new logger
	logger = log.New(file, "[ALS] ", log.LstdFlags|log.Lshortfile)
}

// LogInfo register an info record inside log file
func LogInfo(message string) {
	logger.Println("[INFO] ", message)
}

// LogError register an error record inside log file
func LogError(err error) {
	logger.Println("[ERROR] ", err)
}
