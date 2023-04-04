package logger_adaptor

import (
	"testing"

	"turato.com/design-pattern/adaptor_pattern/logger_adaptor/adapter"
	logger2 "turato.com/design-pattern/adaptor_pattern/logger_adaptor/logger"
)

// logger目录中为日志基础组件
// adapter目录中为日志适配器的具体实现
func TestLoggerAdaptorPattern(t *testing.T) {
	var logger logger2.Logger
	// 初始化 logrus
	logger = adapter.NewLogrusAdapter()
	logger.Info("test logrus adaptor")

	// 初始化 zap
	logger = adapter.NewZapAdapter()
	logger.Info("test zap adaptor")

	// 结果：在当前目录下会生成两个日志文件，一个是利用 logrus 打印的日志，一个是利用 zap 打印的
}
