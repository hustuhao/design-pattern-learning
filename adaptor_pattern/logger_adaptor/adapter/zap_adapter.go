package adapter

import (
	"io"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapAdapter zap 日志适配器
type ZapAdapter struct {
	logger *zap.Logger
}

// NewZapAdapter 创建 zap 适配器
func NewZapAdapter() *ZapAdapter {
	adapter := new(ZapAdapter)
	adapter.logger = NewZapLogger("./") // 指定日志文件路径为当前目录
	return adapter
}

func NewZapLogger(logPath string) *zap.Logger {
	exist, _ := pathExists(logPath)
	if !exist {
		// 创建文件夹
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			log.Fatal("init logger mkdir failed!", err)
		}
	}

	var (
		err        error
		mainWriter io.Writer = os.Stdout
		errWriter  io.Writer = os.Stdout
	)

	mainWriter, err = rotatelogs.New(
		logPath+"/zap.app.log.%Y%m%d",
		rotatelogs.WithMaxAge(5*24*time.Hour),
	)
	if err != nil {
		log.Fatal("log init fail")
	}

	errWriter, err = rotatelogs.New(
		logPath+"/zap.error.log.%Y%m%d",
		rotatelogs.WithMaxAge(5*24*time.Hour),
	)
	if err != nil {
		log.Fatal("log init fail")
	}

	// err 日志另外单独写一个文件
	errLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	mainLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	jsonEnc := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    "func",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEnc, zapcore.Lock(zapcore.AddSync(errWriter)), errLevel),
		zapcore.NewCore(jsonEnc, zapcore.Lock(zapcore.AddSync(mainWriter)), mainLevel),
	)

	return zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(errLevel),
	)
}

func (a *ZapAdapter) Debug(msg string) {
	a.logger.Debug(msg)
}

func (a *ZapAdapter) Info(msg string) {
	a.logger.Info(msg)
}

func (a *ZapAdapter) Warn(msg string) {
	a.logger.Warn(msg)
}

func (a *ZapAdapter) Error(msg string) {
	a.logger.Error(msg)
}
