package convert

import (
	"fmt"
	"github.com/hetianyi/easygo/logger"
	"reflect"
	"testing"
)

func TestNumberToStr(t *testing.T) {

	logger.Info("byte ---------------------------")
	var a1 byte = 1
	var a2 byte = 255
	fmt.Println(NumberToStr(a1))
	fmt.Println(NumberToStr(a2))

	logger.Info("int ---------------------------")
	var b1 int = 16547
	var b2 int = -23123123
	fmt.Println(NumberToStr(b1))
	fmt.Println(NumberToStr(b2))

	logger.Info("int8 ---------------------------")
	var d1 int8 = 100
	var d2 int8 = -100
	fmt.Println(NumberToStr(d1))
	fmt.Println(NumberToStr(d2))

	logger.Info("int16 ---------------------------")
	var e1 int16 = 500
	var e2 int16 = -5000
	fmt.Println(NumberToStr(e1))
	fmt.Println(NumberToStr(e2))

	logger.Info("int32 ---------------------------")
	var f1 int32 = 500
	var f2 int32 = -5000
	fmt.Println(NumberToStr(f1))
	fmt.Println(NumberToStr(f2))

	logger.Info("int64 ---------------------------")
	var g1 int64 = 5007893534534534
	var g2 int64 = -500785345345345345
	fmt.Println(NumberToStr(g1))
	fmt.Println(NumberToStr(g2))

	logger.Info("uint ---------------------------")
	var h1 uint = 16512313
	fmt.Println(NumberToStr(h1))

	logger.Info("uint8 ---------------------------")
	var i1 uint8 = 165
	fmt.Println(NumberToStr(i1))

	logger.Info("uint16 ---------------------------")
	var j1 uint16 = 16567
	fmt.Println(NumberToStr(j1))

	logger.Info("uint32 ---------------------------")
	var k1 uint32 = 1656733242
	fmt.Println(NumberToStr(k1))

	logger.Info("uint64 ---------------------------")
	var l1 uint64 = 5007853453453453452
	fmt.Println(NumberToStr(l1))

	logger.Info("float32 ---------------------------")
	var m1 float32 = 5052.789123123
	fmt.Println(NumberToStr(m1))

	logger.Info("float64 ---------------------------")
	var n1 float64 = 50.78911231231231
	fmt.Println(NumberToStr(n1))
}

func TestStrToNumber(t *testing.T) {

	n0, _ := StrToNumber("101", Byte)
	fmt.Println(n0, reflect.TypeOf(n0))

	n1, _ := StrToNumber("101", Int)
	fmt.Println(n1, reflect.TypeOf(n1))

	n2, _ := StrToNumber("101", Uint)
	fmt.Println(n2, reflect.TypeOf(n2))

	n3, _ := StrToNumber("101", Int8)
	fmt.Println(n3, reflect.TypeOf(n3))

	n4, _ := StrToNumber("101", Uint8)
	fmt.Println(n4, reflect.TypeOf(n4))

	n5, _ := StrToNumber("101", Int16)
	fmt.Println(n5, reflect.TypeOf(n5))

	n6, _ := StrToNumber("101", Uint16)
	fmt.Println(n6, reflect.TypeOf(n6))

	n7, _ := StrToNumber("101", Int32)
	fmt.Println(n7, reflect.TypeOf(n7))

	n8, _ := StrToNumber("101", Uint32)
	fmt.Println(n8, reflect.TypeOf(n8))

	n9, _ := StrToNumber("101", Int64)
	fmt.Println(n9, reflect.TypeOf(n9))

	n10, _ := StrToNumber("101", Uint64)
	fmt.Println(n10, reflect.TypeOf(n10))

	n11, _ := StrToNumber("101.567", Float32)
	fmt.Println(n11, reflect.TypeOf(n11))

	n12, _ := StrToNumber("101123.4123", Float64)
	fmt.Println(n12, reflect.TypeOf(n12))
}
