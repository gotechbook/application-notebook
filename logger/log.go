package logger

import (
	"go.uber.org/zap"
)

var (
	L *zap.Logger
)

func Info(msg string, data interface{}) {
	L.Info(msg, zap.Any("data", data))
}

func Debug(msg string, data interface{}) {
	L.Debug(msg, zap.Any("data", data))
}

func Error(msg string, err error, data interface{}) {
	L.Error(msg, zap.Error(err), zap.Any("data", data))
}
