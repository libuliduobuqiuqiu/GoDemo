package zapdemo

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 输出到分割日志logger
func initRotateLogCore(config zapcore.EncoderConfig) zapcore.Core {
	fileEncoder := zapcore.NewJSONEncoder(config)

	rotateLogger := &lumberjack.Logger{
		Filename:   "app2.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     3,
		Compress:   true,
	}

	return zapcore.NewCore(fileEncoder, zapcore.AddSync(rotateLogger), zapcore.InfoLevel)
}

// 输出到前端
func initConsoleCore(config zapcore.EncoderConfig) zapcore.Core {
	consoleEncoder := zapcore.NewJSONEncoder(config)

	return zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.ErrorLevel)
}

// 将不同级别的日志输出到对应的日志文件（输出到前端，输出到分割日志）
func InitCore() []zapcore.Core {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	rotateCore := initRotateLogCore(config)

	// consoleCore := initConsoleCore(config)
	return []zapcore.Core{rotateCore}
}
