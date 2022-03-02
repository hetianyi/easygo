package logger

import (
	"io"
	"sync"
)

var (
	defaultLogger *logger
	mu            sync.Mutex
)

// InitDefaultLogger 初始化一个默认日志实例
func InitDefaultLogger() {
	mu.Lock()
	defer mu.Unlock()

	if defaultLogger == nil {
		defaultLogger = NewDefaultLogger()
	}
}

func initLogger() {
	if defaultLogger == nil {
		InitDefaultLogger()
	}
}

// SetOut 设置日志输入
func SetOut(out io.Writer) {
	initLogger()
	defaultLogger.Out = out
}

// SetFormatter 设置日志格式器
func SetFormatter(formatter Formatter) {
	initLogger()
	defaultLogger.formatter = formatter
}

// SetPrefix 设置日志前缀
func SetPrefix(prefix string) {
	initLogger()
	defaultLogger.Prefix = prefix
}

// Debug 以Debug级别打印日志
func Debug(log string) {
	initLogger()
	defaultLogger.Debug(log)
}

// Info 以Info级别打印日志
func Info(log string) {
	initLogger()
	defaultLogger.Info(log)
}

// Warn 以Warn级别打印日志
func Warn(log string) {
	initLogger()
	defaultLogger.Warn(log)
}

// Error 以Error级别打印日志
func Error(log string) {
	initLogger()
	defaultLogger.Error(log)
}

// Fatal 以Fatal级别打印日志
func Fatal(log string) {
	initLogger()
	defaultLogger.Fatal(log)
}
