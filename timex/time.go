// Copyright (C) 2019 tisnyo <tisnyo@gmail.com>.
//
// This file provides some functions about time operation.

package timex

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock      = *new(sync.Mutex)
	increment = 0
)

// GetDateString gets short date format like '2018-11-11'.
func GetDateString(t time.Time) string {
	return fmt.Sprintf("%d-%s-%s", GetYear(t), format2(int64(GetMonth(t))), format2(int64(GetDay(t))))
}

// GetCNDateString gets short date format like '2018年11月11日'.
func GetCNDateString(t time.Time) string {
	return fmt.Sprintf("%d年%s月%s日", GetYear(t), format2(int64(GetMonth(t))), format2(int64(GetDay(t))))
}

// GetLongDateString gets date format like '2018-11-11 12:12:12'.
func GetLongDateString(t time.Time) string {
	return fmt.Sprintf("%d-%s-%s %s:%s:%s",
		GetYear(t), format2(int64(GetMonth(t))), format2(int64(GetDay(t))),
		format2(int64(GetHour(t))), format2(int64(GetMinute(t))), format2(int64(GetSecond(t))),
	)
}

// GetLongCNDateString gets date format like '2018年11月11日 12时12分12秒'.
func GetLongCNDateString(t time.Time) string {
	return fmt.Sprintf("%d年%s月%s日 %s时%s分%s秒",
		GetYear(t), format2(int64(GetMonth(t))), format2(int64(GetDay(t))),
		format2(int64(GetHour(t))), format2(int64(GetMinute(t))), format2(int64(GetSecond(t))),
	)
}

// GetShortDateString gets time format like '12:12:12'.
func GetShortDateString(t time.Time) string {
	return fmt.Sprintf("%s:%s:%s",
		format2(int64(GetHour(t))), format2(int64(GetMinute(t))), format2(int64(GetSecond(t))))
}

// GetShortCNDateString gets time format like '12时12分12秒'.
func GetShortCNDateString(t time.Time) string {
	return fmt.Sprintf("%s时%s分%s秒",
		format2(int64(GetHour(t))), format2(int64(GetMinute(t))), format2(int64(GetSecond(t))))
}

// GetLongFullDateString gets short date format like '2018-11-11 12:12:12,233'.
func GetLongFullDateString(t time.Time) string {
	return fmt.Sprintf("%d-%s-%s %s:%s:%s.%s",
		GetYear(t), format2(int64(GetMonth(t))), format2(int64(GetDay(t))),
		format2(int64(GetHour(t))), format2(int64(GetMinute(t))), format2(int64(GetSecond(t))), format3(int64(GetMillionSecond(t))),
	)
}

// GetTimestamp gets current timestamp in milliseconds.
func GetTimestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// CreateTime creates a time from a millis second.
func CreateTime(millis int64) time.Time {
	return time.Unix(millis/1000, 0)
}

// GetNanosecond gets current timestamp in Nanosecond.
func GetNanosecond(t time.Time) int64 {
	return t.UnixNano()
}

// GetYear gets year number.
func GetYear(t time.Time) int {
	return t.Year()
}

// GetMonth gets month number.
func GetMonth(t time.Time) int {
	return int(t.Month())
}

// GetDay gets the day of the month.
func GetDay(t time.Time) int {
	return t.Day()
}

// GetHour gets hour number.
func GetHour(t time.Time) int {
	return t.Hour()
}

// GetMinute gets minute number.
func GetMinute(t time.Time) int {
	return t.Minute()
}

// GetSecond gets second number.
func GetSecond(t time.Time) int {
	return t.Second()
}

// GetWeekDay 获取当前日期是星期几
func GetWeekDay(t time.Time) string {
	return t.Weekday().String()
}

// GetCNWeekDay 获取当前日期是星期几
func GetCNWeekDay(t time.Time) string {
	switch t.Weekday() {
	case time.Sunday:
		return "周日"
	case time.Monday:
		return "周一"
	case time.Tuesday:
		return "周二"
	case time.Wednesday:
		return "周三"
	case time.Thursday:
		return "周四"
	case time.Friday:
		return "周五"
	default:
		return "周六"
	}
}

// GetMillionSecond gets millionSecond number.
func GetMillionSecond(t time.Time) int {
	return t.Nanosecond() / 1e6
}

// ParseTimeFromRFC3339 解析时间格式：
//
// 2006-01-02T15:04:05Z
func ParseTimeFromRFC3339(timeString string) (time.Time, error) {
	return time.Parse(time.RFC3339, timeString)
}

// ParseTime 解析时间格式：
//
// 2006-01-02 15:04:05
func ParseTime(timeString string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", timeString)
}

// ParseDate 解析时间格式：
//
// 2006-01-02
func ParseDate(timeString string) (time.Time, error) {
	return time.Parse("2006-01-02", timeString)
}

func GetDuration(start, end time.Time) string {
	val := (GetTimestamp(end) - GetTimestamp(start)) / 1000
	if val < 60 {
		return fmt.Sprintf("%ds", val)
	}
	if val < 3600 {
		return fmt.Sprintf("%dm %ds", val/60, val%60)
	}
	if val < 86400 {
		return fmt.Sprintf("%dh %dm %ds", val/3600, val%3600/60, val%60)
	}
	return fmt.Sprintf("%dd %dh %dm %ds", val/86400, val%86400/3600, val%3600/60, val%60)
}

func GetCNDuration(start, end time.Time) string {
	val := (GetTimestamp(end) - GetTimestamp(start)) / 1000
	if val < 60 {
		return fmt.Sprintf("%d秒", val)
	}
	if val < 3600 {
		return fmt.Sprintf("%d分%d秒", val/60, val%60)
	}
	if val < 86400 {
		return fmt.Sprintf("%d小时%d分%d秒", val/3600, val%3600/60, val%60)
	}
	return fmt.Sprintf("%d天%d小时%d分%d秒", val/86400, val%86400/3600, val%3600/60, val%60)
}

func format2(input int64) string {
	return fmt.Sprintf("%02d", input)
}

func format3(input int64) string {
	return fmt.Sprintf("%03d", input)
}
