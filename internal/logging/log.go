package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func InitLoggers() {
	logFile := os.Getenv("BF_LOG_PATH")
	layout := "2006-01-02T15:04:05-0700"
	if len(logFile) != 0 {
		logFile += fmt.Sprintf("_%s.log", time.Now().Format(layout))
	}
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
