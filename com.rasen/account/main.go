package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	"com.rasen/account/netregistry"
	"github.com/sirupsen/logrus"
)

var logger *Log

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
		logger = &Log{Logger: logrusLogger, File: fd}
	})
}

func main(){
	logger.WithFields(logrus.Fields{"state":"run"}).Info("start account server")
	defer logger.File.Close()
	// 开启web服务
	go	netregistry.RegistryHub()

	// 退出
	exit := make(chan os.Signal)
	signal.Notify(exit,os.Kill,os.Interrupt,syscall.SIGTERM,syscall.SIGQUIT)
	for c := range exit{
		switch c {
		case os.Kill,os.Interrupt,syscall.SIGTERM,syscall.SIGQUIT:
			logger.Info("exit")
			os.Exit(0)
		default:
			fmt.Println("watch sig")
		}
	}
}
