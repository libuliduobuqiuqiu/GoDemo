package zapdemo

import (
	"github.com/go-faker/faker/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 设置日志输出等级、日志输出文件、堆栈信息打印、日志输出格式、
// 日志通过钩子保存记录到日志文件

type zapLogData struct {
	Msg     string `json:"msg"`
	Url     string `json:"url"`
	Attempt string `json:"attempt"`
	Backoff int    `json:"backoff"`
}

// Zap 正常logger输出日志性能更快，suagar logger输出日志支持结构化；
// Zap 通过自定输出encoder，设置输出等级，输出格式，堆栈信息等；
// Zap通过core可以自定义多个输出；
func UseZapLogging() error {
	cores := InitCore()
	core := zapcore.NewTee(cores...)

	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	data := &zapLogData{}
	for i := 0; i < 30; i++ {

		err := faker.FakeData(data)
		if err != nil {
			return err
		}
		logger.Info(
			faker.Sentence(),
			zap.String("Message", data.Msg),
			zap.String("Url", data.Url),
			zap.String("Attempt", data.Attempt),
			zap.Int("backoff", data.Backoff),
		)
	}

	for i := 0; i < 30; i++ {
		err := faker.FakeData(data)
		if err != nil {
			return err
		}
		logger.Error(
			faker.Sentence(),
			zap.String("Message", data.Msg),
			zap.String("Url", data.Url),
			zap.String("Attempt", data.Attempt),
			zap.Int("backoff", data.Backoff),
		)

	}

	return nil
}
