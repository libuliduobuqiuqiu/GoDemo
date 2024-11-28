package logrusdemo

import (
	"os"

	"github.com/sirupsen/logrus"
)

type consoleHook struct {
	formatter *logrus.JSONFormatter
}

func (c *consoleHook) Levels() []logrus.Level {
	return logrus.AllLevels[:3]
}

func (c *consoleHook) Fire(entry *logrus.Entry) error {

	data, err := c.formatter.Format(entry)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(data)

	return err
}

func NewConsoleHook() logrus.Hook {
	return &consoleHook{formatter: &logrus.JSONFormatter{}}
}
