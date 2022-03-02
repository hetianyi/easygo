package logger

import (
	"github.com/hetianyi/easygo/base"
	"io"
	"os"
	"sync"
	"time"
)

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var (
	levelNameMapping      = []string{"debug", "info", "warn", "error", "fatal"}
	levelShortNameMapping = []string{"D", "I", "W", "E", "F"}
)

type logger struct {
	Out       io.Writer
	lock      sync.Mutex
	Prefix    string
	Level     LogLevel
	formatter Formatter
}

type Formatter func(l *logger, t time.Time, level LogLevel, v ...interface{})

// New 初始化一个日志实例
func New(out io.Writer, prefix string, level LogLevel, formatter Formatter) *logger {
	return &logger{
		Out:       out,
		Prefix:    prefix,
		Level:     level,
		formatter: base.TValue(formatter == nil, DefaultFormatter, formatter).(Formatter),
	}
}

// NewDefaultLogger 初始化一个默认日志实例
func NewDefaultLogger() *logger {
	return &logger{
		Out:       os.Stdout,
		Level:     LevelInfo,
		formatter: DefaultFormatter,
	}
}

// Info 以Info级别打印日志
func (l *logger) output(level LogLevel, v ...interface{}) {
	// 日志级别低于设置的级别，则不处理
	if level < l.Level {
		return
	}
	l.formatter(l, time.Now(), level, v...)
}

// SetOut 设置日志输入
func (l *logger) SetOut(out io.Writer) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.Out = out
}

// SetFormatter 设置日志格式器
func (l *logger) SetFormatter(formatter Formatter) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.formatter = formatter
}

// SetPrefix 设置日志前缀
func (l *logger) SetPrefix(prefix string) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.Prefix = prefix
}

// Debug 以Debug级别打印日志
func (l *logger) Debug(v ...interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.output(LevelDebug, v...)
}

// Info 以Info级别打印日志
func (l *logger) Info(v ...interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.output(LevelInfo, v...)
}

// Warn 以Warn级别打印日志
func (l *logger) Warn(v ...interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.output(LevelWarn, v...)
}

// Error 以Error级别打印日志
func (l *logger) Error(v ...interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.output(LevelError, v...)
}

// Fatal 以Fatal级别打印日志
func (l *logger) Fatal(v ...interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.output(LevelFatal, v...)
}

func ParseLevel(levelName string) LogLevel {
	for i, v := range levelNameMapping {
		if v == levelName {
			return LogLevel(i)
		}
	}
	return LevelInfo
}
