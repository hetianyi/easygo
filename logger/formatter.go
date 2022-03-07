package logger

import (
	"fmt"
	"github.com/hetianyi/easygo/timex"
	. "github.com/logrusorgru/aurora"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// DefaultFormatter 是默认的日志格式器，输出格式为
//
//  [I] [PREFIX] 2022-03-02 21:35:43 | xxxxx
//  [E] [PREFIX] 2022-03-02 21:35:43 | xxxxx
func DefaultFormatter(l *logger, t time.Time, level LogLevel, v ...interface{}) {
	var log string
	if l.ColorableOutput {
		log = fmt.Sprintf("[%s] %s%s %s | %s\n",
			levelColor(level), l.Prefix,
			timeColor(timex.GetLongDateString(t)),
			GetCaller(l.callLevel, l.ColorableOutput),
			logColor(level, fmt.Sprint(v...)))
	} else {
		log = fmt.Sprintf("[%s] %s%s %s | %s\n",
			levelShortNameMapping[level], l.Prefix,
			timex.GetLongDateString(t),
			GetCaller(l.callLevel, l.ColorableOutput),
			fmt.Sprint(v...))
	}
	l.Out.Write([]byte(log))
}

// SimpleFormatter 是简化的日志格式器，输出格式为
//
//  [I] [PREFIX] 2022-03-02 21:35:43 | xxxxx
//  [E] [PREFIX] 2022-03-02 21:35:43 | xxxxx
func SimpleFormatter(l *logger, t time.Time, level LogLevel, v ...interface{}) {
	log := fmt.Sprintf("[%s] %s%s | %s\n",
		levelColor(level), l.Prefix,
		timeColor(timex.GetLongDateString(t)),
		logColor(level, fmt.Sprint(v...)))
	l.Out.Write([]byte(log))
}

func levelColor(level LogLevel) string {
	switch level {
	case LevelDebug:
		return BrightBlack(levelShortNameMapping[level]).String()
	case LevelInfo:
		return BrightGreen(levelShortNameMapping[level]).String()
	case LevelWarn:
		return BrightYellow(levelShortNameMapping[level]).String()
	case LevelError:
		return Red(levelShortNameMapping[level]).String()
	case LevelFatal:
		return SlowBlink(BgRed(levelShortNameMapping[level])).String()
	}
	return ""
}

func timeColor(t string) string {
	return Cyan(t).String()
}

func logColor(level LogLevel, content string) string {
	switch level {
	case LevelDebug:
		return Reverse(BrightBlack(content)).String()
	case LevelInfo:
		return content
	case LevelWarn:
		return BrightYellow(content).String()
	case LevelError:
		return Red(content).String()
	case LevelFatal:
		return SlowBlink(BgRed(content)).String()
	}
	return content
}

func GetCaller(callLevel int, colorable bool) string {
	_, file, line, success := runtime.Caller(callLevel)
	if success {
		if colorable {
			return Yellow(strings.Join([]string{"[", file[strings.LastIndex(file, "/")+1:], ":", strconv.Itoa(line), "]"}, "")).String()
		} else {
			return strings.Join([]string{"[", file[strings.LastIndex(file, "/")+1:], ":", strconv.Itoa(line), "]"}, "")
		}
	}
	return " [unknown] "
}
