package structlog

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
)

var Logger *Log

type Log struct {
	*logrus.Logger
	*os.File
}


func init() {
	singleLogger := sync.Once{}
	singleLogger.Do(func() {
		os.Mkdir("logs", os.ModePerm)
		logPath := filepath.Join(".", "logs", "log.txt")
		fd, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Fprint(os.Stderr, "open or create log file fail. err:", err)
		}
		logrusLogger := logrus.New()
		logrusLogger.Out = fd
		logrusLogger.Formatter = &logrus.TextFormatter{TimestampFormat: "02-01-2006 15:04:05", FullTimestamp: true}
		Logger = &Log{Logger: logrusLogger, File: fd}
	})
}