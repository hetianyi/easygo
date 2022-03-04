package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDateString(t *testing.T) {
	fmt.Println(GetDateString(time.Now()))
	fmt.Println(GetCNDateString(time.Now()))
	fmt.Println(GetLongDateString(time.Now()))
	fmt.Println(GetLongCNDateString(time.Now()))
	fmt.Println(GetShortDateString(time.Now()))
	fmt.Println(GetShortCNDateString(time.Now()))
	fmt.Println(GetLongFullDateString(time.Now()))
}

func TestGetTimestamp(t *testing.T) {
	fmt.Println(GetTimestamp(time.Now()))
	fmt.Println(GetMillionSecond(time.Now()))
}

func TestCreateTime(t *testing.T) {
	fmt.Println(CreateTime(GetTimestamp(time.Now()) / 1000))
}

func TestGetNanosecond(t *testing.T) {
	fmt.Println(GetTimestamp(time.Now()))
	fmt.Println(GetNanosecond(time.Now()))
}

func TestParseTimeFromRFC3339(t *testing.T) {
	a, _ := ParseTimeFromRFC3339("2022-02-04T15:04:05Z")
	fmt.Println(GetLongDateString(a))
}

func TestParseTime(t *testing.T) {
	a, _ := ParseTime("2022-02-04 15:04:05")
	fmt.Println(GetLongDateString(a))
}

func TestParseDate(t *testing.T) {
	a, _ := ParseDate("2022-02-04")
	fmt.Println(GetLongDateString(a))
}

func TestGetWeekDay(t *testing.T) {
	fmt.Println(GetWeekDay(time.Now()))
	fmt.Println(GetCNWeekDay(time.Now()))
}

func TestGetDay(t *testing.T) {
	t1 := time.Now()
	t2 := time.Now().Add(time.Second * 55)
	t3 := time.Now().Add(time.Minute*3 + time.Second*55)
	t4 := time.Now().Add(time.Hour*3 + time.Minute*3 + time.Second*55)
	t5 := time.Now().Add(time.Hour*26 + time.Minute*3 + time.Second*55)
	t6 := time.Now().Add(time.Hour*60 + time.Minute*3 + time.Second*55)
	fmt.Println(GetDuration(t1, t2))
	fmt.Println(GetDuration(t1, t3))
	fmt.Println(GetDuration(t1, t4))
	fmt.Println(GetDuration(t1, t5))
	fmt.Println(GetDuration(t1, t6))

	fmt.Println(GetCNDuration(t1, t2))
	fmt.Println(GetCNDuration(t1, t3))
	fmt.Println(GetCNDuration(t1, t4))
	fmt.Println(GetCNDuration(t1, t5))
	fmt.Println(GetCNDuration(t1, t6))
}
