package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func NewLogger(debug bool, logFilePath string) *zap.SugaredLogger {
	atom := zap.NewAtomicLevel()

	if debug {
		atom.SetLevel(zap.DebugLevel)
	} else {
		atom.SetLevel(zap.InfoLevel)
	}

	var writer io.Writer
	if logFilePath == "" {
		log.Fatal("Please set BF_LOG_PATH environment variable.")
	} else {
		logFilePath, err := filepath.Abs(logFilePath)
		if err != nil {
			log.Fatal(err)
		}
		logFilePath += "." + time.Now().Format("20060102150405")
		writer, err = os.Create(logFilePath)
		if err != nil {
			log.Fatal(err)
		}
	}

	encoder := zap.NewDevelopmentEncoderConfig()
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.AddSync(writer),
		atom))
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger)

	if debug {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger.Sugar()
}
