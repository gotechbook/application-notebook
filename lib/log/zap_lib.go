package log

import (
	"fmt"
	"os"
	"time"

	"github.com/gotechbook/application-notebook/config"
	"github.com/gotechbook/application-notebook/helper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	level zapcore.Level
)

func Zap() (logger *zap.Logger) {
	var (
		director = config.C.Zap.Director
		zLevel   = config.C.Zap.Level
		showLine = config.C.Zap.ShowLine
	)
	if ok, _ := helper.PathExists(director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", director)
		_ = os.Mkdir(director, os.ModePerm)
	}
	switch zLevel { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if showLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (cfg zapcore.EncoderConfig) {
	cfg = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  config.C.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case config.C.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		cfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	case config.C.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		cfg.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case config.C.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	case config.C.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		cfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return cfg
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if config.C.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(config.V.GetString("zap.prefix") + "2006/01/02 - 15:04:05.000"))
}
