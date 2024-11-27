package logrusdemo

import (
	"github.com/sirupsen/logrus"
	"godemo/internal/golib/osdemo"
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

func UseLogrus() error {
	var log = logrus.StandardLogger()

	// Set Output File
	// writer, err := os.Create("tmp.log")
	// if err != nil {
	// 	return
	// }
	// log.SetOutput(writer)

	hook, err := NewLogFileHook("tmpHook.log", &logrus.JSONFormatter{})
	if err != nil {
		return err
	}

	log.AddHook(&stackHook{})
	log.AddHook(hook)
	log.SetLevel(logrus.ErrorLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{})

	err = osdemo.PrintFilePath("/roots")
	if err != nil {
		logrus.WithError(err).Error()
	}

	return nil
}
