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
