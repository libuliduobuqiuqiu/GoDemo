package logrusdemo

import (
	"io"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var tmpRotateConfig = &RotateConfig{
	FileName:   "app.log",
	MaxSize:    100,
	MaxBackups: 3,
	MaxAges:    3,
}

type RotateConfig struct {
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAges    int
}

type rotateLogHook struct {
	config    *RotateConfig
	logWriter io.Writer
	formatter logrus.Formatter
}

func (r *rotateLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (r *rotateLogHook) Fire(entry *logrus.Entry) error {
	data, err := r.formatter.Format(entry)
	if err != nil {
		return err
	}

	_, err = r.logWriter.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// 切割归档日志的hook
func NewRotateLogFileHook(c *RotateConfig) logrus.Hook {
	if c == nil {
		c = tmpRotateConfig
	}

	r := &rotateLogHook{config: c, formatter: &logrus.JSONFormatter{}}
	r.logWriter = &lumberjack.Logger{
		Filename:   c.FileName,
		MaxSize:    c.MaxSize,
		MaxAge:     c.MaxAges,
		MaxBackups: c.MaxBackups,
		Compress:   true,
	}
	return r
}
