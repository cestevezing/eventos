package logs

import (
	"log"
	"os"
)

var (
	Logger *log.Logger
)

func InitLogger() {
	NewLogger("internal/logs/errors.log")
}

func NewLogger(path string) {
	var errorLogPath = path
	var errorFile, err = os.OpenFile(errorLogPath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Error", err)
	}
	Logger = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
