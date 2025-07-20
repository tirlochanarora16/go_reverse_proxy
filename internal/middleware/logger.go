package middleware

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLogger() {
	enccoderConfig := zap.NewProductionEncoderConfig()
	enccoderConfig.TimeKey = "time"
	enccoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	enccoderConfig.StacktraceKey = ""

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(enccoderConfig),
		zapcore.AddSync(os.Stdout), // destination for the logs
		zap.InfoLevel,
	)

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(enccoderConfig),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   "./logs/proxy.log",
			MaxSize:    10,
			MaxBackups: 3,
			MaxAge:     30,
			Compress:   true,
		}),
		zap.InfoLevel,
	)

	core := zapcore.NewTee(consoleCore, fileCore)

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}
