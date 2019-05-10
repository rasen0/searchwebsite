package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/go-kit/kit/log"
	logrus2 "github.com/go-kit/kit/log/logrus"
	"github.com/sirupsen/logrus"
)

var logger *Log

type Log struct {
	log.Logger
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
		logger = &Log{Logger: logrus2.NewLogrusLogger(logrusLogger), File: fd}
	})
}

func main() {
	logger.Log("name", "jack")
	logger.File.Close()

	exit := make(chan os.Signal)
	fmt.Println("qidong")
	signal.Notify(exit,os.Kill,os.Interrupt,syscall.SIGTERM)
	for c := range exit{
		switch c {
		case os.Kill,os.Interrupt,syscall.SIGKILL,syscall.SIGTERM,syscall.SIGQUIT:
			os.Exit(0)
		default:
			fmt.Println("watch sig")
		}
	}
}
