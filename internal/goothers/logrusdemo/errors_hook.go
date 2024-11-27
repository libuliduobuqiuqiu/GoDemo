package logrusdemo

import (
	"runtime"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type stackHook struct{}

func (s *stackHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (s *stackHook) Fire(entry *logrus.Entry) error {

	if e, ok := entry.Data[logrus.ErrorKey]; ok {
		if tmp, ok := e.(error); ok {
			entry.Message = tmp.Error()
		}

		SetCallers(entry)
		delete(entry.Data, logrus.ErrorKey)
	}
	return nil
}

// 捕获堆栈信息
// 固定格式，不灵活
// 适合直接获取和输出堆栈信息
func SetStack(entry *logrus.Entry) {
	buf := make([]byte, 1<<16)
	n := runtime.Stack(buf, false)
	entry.Data["stack"] = string(buf[:n])
}

type stack struct {
	File     string
	Line     int
	FuncName string
}

// 捕获堆栈信息,需要解析,可以灵活定制
// skip 指定跳过调用栈桢数量
func SetCallers(entry *logrus.Entry) {
	var stacks []stack
	pcs := make([]uintptr, 32)
	num := runtime.Callers(6, pcs)

	frames := runtime.CallersFrames(pcs[:num])
	for {
		frame, more := frames.Next()
		stacks = append(stacks, stack{
			File:     frame.File,
			Line:     frame.Line,
			FuncName: frame.Function,
		})

		if !more {
			break
		}
	}
	entry.Data["stack"] = stacks
}
