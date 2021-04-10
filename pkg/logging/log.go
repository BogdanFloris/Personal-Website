package logging

import (
	"io"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func InitLoggers() {
	file, err := os.OpenFile(os.Getenv("BF_LOG_PATH"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	var loggerWriter io.Writer
	if err != nil {
		loggerWriter = os.Stdout
	} else {
		loggerWriter = file
	}

	InfoLogger = log.New(loggerWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(loggerWriter, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(loggerWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
