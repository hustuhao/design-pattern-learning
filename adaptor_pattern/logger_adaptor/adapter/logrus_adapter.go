package adapter

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"turato.com/design-pattern/adaptor_pattern/logger_adaptor/logger"
)

// LogrusAdapter logrus 日志适配器
type LogrusAdapter struct {
	logger *logrus.Logger
}

//var loggerAdapter LogrusAdapter

func (a *LogrusAdapter) Debug(msg string) {
	a.logger.Debug(msg)
}

func (a *LogrusAdapter) Info(msg string) {
	a.logger.Info(msg)
}

func (a *LogrusAdapter) Warn(msg string) {
	a.logger.Warn(msg)
}

func (a *LogrusAdapter) Error(msg string) {
	a.logger.Error(msg)
}

func NewLogrusAdapter() *LogrusAdapter {
	adapter := new(LogrusAdapter)
	adapter.logger = NewLogrusLogger("./")
	return adapter
}

// path路径, 末尾要有/
// tagName 一般为进程名
func NewLogrusLogger(logPath string) *logrus.Logger {
	exist, _ := pathExists(logPath)
	if !exist {
		// 创建文件夹
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			log.Fatal("init logger mkdir failed!", err)
		}
	}
	logrus.SetOutput(ioutil.Discard) //使用hook输出日志，丢弃原有的write操作
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	appLog, err := rotatelogs.New(
		logPath+"/logrus.app.log.%Y%m%d",
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		log.Fatal("log init fail")
	}
	errorLog, err := rotatelogs.New(
		logPath+"/logrus.error.log.%Y%m%d",
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		log.Fatal("log init fail")
	}
	//为不同级别设置不同的输出目的
	lfHook := logger.NewHook(
		logger.WriterMap{
			logrus.DebugLevel: appLog,
			logrus.InfoLevel:  appLog,
			logrus.WarnLevel:  appLog,
			logrus.ErrorLevel: errorLog,
			logrus.FatalLevel: errorLog,
			logrus.PanicLevel: errorLog,
		},
		&logrus.JSONFormatter{},
	)

	logger := logrus.New()
	logger.AddHook(lfHook)
	return logger
}

// 判断文件夹是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
