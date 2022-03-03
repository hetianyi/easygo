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
func initDefaultLogger() {
	mu.Lock()
	defer mu.Unlock()

	if defaultLogger == nil {
		defaultLogger = newDefaultLogger()
	}
}

func initLogger() {
	if defaultLogger == nil {
		initDefaultLogger()
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
func Debug(v ...interface{}) {
	initLogger()
	defaultLogger.Debug(v...)
}

// Info 以Info级别打印日志
func Info(v ...interface{}) {
	initLogger()
	defaultLogger.Info(v...)
}

// Warn 以Warn级别打印日志
func Warn(v ...interface{}) {
	initLogger()
	defaultLogger.Warn(v...)
}

// Error 以Error级别打印日志
func Error(v ...interface{}) {
	initLogger()
	defaultLogger.Error(v...)
}

// Fatal 以Fatal级别打印日志
func Fatal(v ...interface{}) {
	initLogger()
	defaultLogger.Fatal(v...)
}
