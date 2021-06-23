package main

import (
	"fmt"
	"os"
	_ "reflect"
	_ "runtime"

	"time"

	"github.com/sirupsen/logrus"
	. "github.com/sirupsen/logrus"
)

const logDirectory = "./logs"

var logFile = fmt.Sprintf("installer%d.log", time.Now().Unix())

var LOGGER *logrus.Logger = newLogger()

func newLogger() *Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	mkerr := os.MkdirAll(logDirectory, os.ModePerm)
	if mkerr != nil {
		panic(mkerr)
	}

	var filePath = fmt.Sprintf("%s/%s", logDirectory, logFile)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)

	if err != nil {
		panic(err)
	}

	logger.SetOutput(file)
	logger.SetFormatter(&logrus.TextFormatter{})
	return logger
}
