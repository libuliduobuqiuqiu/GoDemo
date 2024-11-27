package zapdemo

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 设置日志输出等级、日志输出文件、堆栈信息打印、日志输出格式、
// 日志通过钩子保存记录到日志文件
//

func UseZapLogging() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)

	rotateLogger := &lumberjack.Logger{
		Filename:   "app2.log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     3,
		Compress:   true,
	}

	core := zapcore.NewCore(fileEncoder, zapcore.AddSync(rotateLogger), zapcore.InfoLevel)
	logger := zap.New(core)
	defer logger.Sync()
	logger.Info("failed to featch the Url",
		zap.String("url", "http://www.google.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
