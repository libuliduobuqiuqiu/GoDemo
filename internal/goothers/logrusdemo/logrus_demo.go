package logrusdemo

import (
	"os"

	"github.com/sirupsen/logrus"
)

// 日志模块主要用法：设置日志级别、设置日志输出、设置日志钩子函数（打印堆栈，分割日志）

// var log = logrus.New()

// refer: https://github.com/rifflock/lfshook
// About logrus hook.
// Attempt to write log on the text file.
type logFileHook struct {
	file      *os.File
	formatter logrus.Formatter
}

func (l *logFileHook) Close() {
	if l.file != nil {
		l.file.Close()
	}
}

func (l *logFileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *logFileHook) Fire(h *logrus.Entry) error {
	data, err := l.formatter.Format(h)
	if err != nil {
		return err
	}

	_, err = l.file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func NewLogFileHook(path string, f logrus.Formatter) (*logFileHook, error) {
	w, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &logFileHook{file: w, formatter: f}, nil
}

func printFileNotExist(path string) error {
	_, err := os.Open(path)
	if err != nil {
		logrus.WithError(err).Error()
		return err
	}
	return nil
}

func OpenFile(path string) {
	err := printFileNotExist(path)
	if err != nil {
		logrus.WithError(err).Error()
	}
}

func UseLogrus() error {
	var log = logrus.StandardLogger()

	// writer, err := os.Create("tmp.log")
	// if err != nil {
	// 	return
	// }
	// log.SetOutput(writer)
	//
	// log.SetFormatter(&log.JSONFormatter{})
	//
	hook, err := NewLogFileHook("tmpHook.log", &logrus.JSONFormatter{})
	if err != nil {
		return err
	}

	log.AddHook(hook)
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	// log.SetLevel(logrus.WarnLevel)
	log.SetReportCaller(true)
	log.Info("Start logrus.")
	log.Warn("Test warning.")
	log.Error("Test Error.")

	log.WithFields(logrus.Fields{
		"code":     200,
		"message":  "success",
		"use_time": 50,
		"desc":     "About logrus.",
	}).Info("Test Info")

	OpenFile("/root/tmp.log")

	return nil
}
