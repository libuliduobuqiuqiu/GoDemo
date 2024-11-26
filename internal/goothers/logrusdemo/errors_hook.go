package logrusdemo

import "github.com/sirupsen/logrus"

type stackHook struct {
}

func (s *stackHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (s *stackHook) Fire(entry *logrus.Entry) error {

	for _, e := entry.Data[logrus.ErrorKey] {
	}
}
