package logger

import (
	"log"
	"os"
)

func NewLogger() (Logger, Logger, Logger) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return InfoLogger, WarningLogger, ErrorLogger
}
