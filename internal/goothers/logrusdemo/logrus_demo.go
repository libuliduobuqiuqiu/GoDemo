package logrusdemo

import (
	"godemo/internal/golib/osdemo"
	"io"

	"github.com/sirupsen/logrus"
)

func logrusPrintLog() {
	logrus.Info("Start logrus.")
	logrus.Warn("Test warning.")
	logrus.Error("Test Error.")
	logrus.WithFields(logrus.Fields{
		"code":     200,
		"message":  "success",
		"use_time": 50,
		"desc":     "About logrus.",
	}).Info("Test Info")
}

func NewLogrusLogger() *logrus.Logger {
	var logger = logrus.StandardLogger()
	rotateLogHook := NewRotateLogFileHook(nil)
	logger.AddHook(rotateLogHook)

	// consoleLogHook := NewConsoleHook()
	// logger.AddHook(consoleLogHook)

	logger.SetOutput(io.Discard)
	logger.SetLevel(logrus.ErrorLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}

func UseLogrus() error {
	var log = logrus.StandardLogger()

	// Set Output File
	// writer, err := os.Create("tmp.log")
	// if err != nil {
	// 	return
	// }
	// log.SetOutput(writer)

	logHook, err := NewLogFileHook("tmpHook.log", &logrus.JSONFormatter{})
	if err != nil {
		return err
	}
	rotateLogHook := NewRotateLogFileHook(nil)
	consoleHook := NewConsoleHook()

	log.AddHook(&stackHook{})
	log.AddHook(logHook)
	log.AddHook(rotateLogHook)
	log.AddHook(consoleHook)

	log.SetLevel(logrus.InfoLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{})

	// 取消直接打印到前端
	log.SetOutput(io.Discard)

	err = osdemo.PrintFilePath("/roots")
	if err != nil {
		logrus.WithError(err).Error()
	}

	log.Info("success")

	return nil
}
