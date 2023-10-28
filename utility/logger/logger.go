package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/mhshahin/cool-service-go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugaredLogger *zap.SugaredLogger

func InitLogger(cfg *config.AppConfig) error {
	filePath := fmt.Sprintf("%s/log-%s.log", cfg.Log.LogFilePath, time.Now().Format("2006-01-02T15:04"))

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), getLogLevel(cfg.Log.FileLogLevel)),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), getLogLevel(cfg.Log.StdoutLogLevel)),
	)

	sugaredLogger = zap.New(core, zap.AddCaller()).Sugar()

	return nil
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func GetSugaredLogger() *zap.SugaredLogger {
	return sugaredLogger
}
