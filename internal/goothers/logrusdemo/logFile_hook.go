package logrusdemo

import (
	"os"

	"github.com/sirupsen/logrus"
)

// refer: https://github.com/rifflock/lfshook
// About logrus hook.
// Attempt to write log on the text file.
type logFileHook struct {
	file      *os.File
	formatter logrus.Formatter
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
