package utils

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var CustomLogger *logrus.Logger

func InitCustomLogger() {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	// Parse the log level
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
		fmt.Printf("Invalid log level '%s', defaulting to info\n", logLevel)
	}

	// Setup log file writer
	logFile := os.Getenv("LOG_FILE")
	if logFile == "" {
		logFile = "logs/api.log"
	}
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}
	multiWriter := io.MultiWriter(os.Stdout, file)

	// Setup logger
	CustomLogger = logrus.New()
	CustomLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors:          true,
		DisableLevelTruncation: true,
		FullTimestamp:          true,
		PadLevelText:           true,
		QuoteEmptyFields:       true,
		TimestampFormat:        time.RFC3339,
	})
	CustomLogger.SetLevel(level)
	CustomLogger.SetOutput(multiWriter)
}
