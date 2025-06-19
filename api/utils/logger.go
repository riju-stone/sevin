package utils

import (
	"fmt"
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

	CustomLogger = logrus.New()
	CustomLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
	CustomLogger.SetLevel(level)
	CustomLogger.SetOutput(os.Stdout)
}
