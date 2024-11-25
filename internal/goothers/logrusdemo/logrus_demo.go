package logrusdemo

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func UseLogrus() {
	writer, err := os.Create("tmp.log")
	if err != nil {
		return
	}

	log.SetOutput(writer)
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetLevel(logrus.WarnLevel)
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

}
