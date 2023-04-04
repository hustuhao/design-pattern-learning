package logger

// Logger 统一的日志接口
// 适配器结构体需要实现该接口
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}
