package app

import (
	"github.com/hetianyi/easygo/convert"
	json "github.com/json-iterator/go"
	"os"
	"reflect"
	"testing"
)

func TestNewApp(t *testing.T) {
	c := &Server{
		Host:         "0.0.0.0",
		Port:         8080,
		GinMode:      "release",
		UseGinLogger: false,
		MiddleWares:  nil,
		Validators:   nil,
		ApiGroup:     nil,
	}

	typ := reflect.TypeOf(c)
	println(typ.Name, " ", typ.Kind(), " ", typ.Kind().String())

	elem := reflect.ValueOf(&c).Elem()
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		println(f.Name, " ", f.Type.Kind(), " ", f.Type.String())

		envValue, exist := os.LookupEnv(f.Tag.Get("env"))
		if exist {
			v := convertType(envValue, f.Type.Kind())
			elem.Field(i).Set(v)
		}
	}
	toString, _ := json.MarshalToString(&c)
	println(toString)
}

func resolveStruct(a interface{}) {
	realType, is := isStruct(a)
	if !is {
		panic("无法处理，输入不是一个struct类型")
	}
	val := extractPointer(a)
	for i := 0; i < realType.NumField(); i++ {
		f := realType.Field(i)
		println(f.Name, " ", f.Type.Kind(), " ", f.Type.String())

		envValue, exist := os.LookupEnv(f.Tag.Get("env"))
		if exist {
			v := convertType(envValue, f.Type.Kind())
			val.Field(i).Set(v)
		}
	}
	println(a)
}

func TestNewApp2(t *testing.T) {
	c := &Server{
		Host:         "0.0.0.0",
		Port:         8080,
		GinMode:      "release",
		UseGinLogger: false,
		MiddleWares:  nil,
		Validators:   nil,
		ApiGroup:     nil,
	}

	typ := reflect.TypeOf(&c)
	println(typ.Name, " ", typ.Kind(), " ", typ.Kind().String())
	for {
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
			continue
		}
		break
	}
	println(typ.Name, " ", typ.Kind(), " ", typ.Kind().String())

	resolveStruct(&Server{
		Host:         "0.0.0.0",
		Port:         8080,
		GinMode:      "release",
		UseGinLogger: false,
		MiddleWares:  nil,
		Validators:   nil,
		ApiGroup:     nil,
	})
}

func isStruct(s interface{}) (reflect.Type, bool) {
	typ := reflect.TypeOf(s)
	println(typ.Name, " ", typ.Kind(), " ", typ.Kind().String())
	for {
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
			continue
		}
		break
	}
	return typ, typ.Kind() == reflect.Struct
}

// 剥掉指针
func extractPointer(s interface{}) reflect.Value {
	typ := reflect.TypeOf(s)
	val := reflect.ValueOf(s).Elem()
	println(typ.Name, " ", typ.Kind(), " ", typ.Kind().String())
	for {
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
			val = val.Elem()
			continue
		}

		break
	}
	return val
}

func convertType(value string, typ reflect.Kind) reflect.Value {
	var v interface{}
	switch typ {
	case reflect.Bool:
		v, _ = convert.StrToBool(value)
		break
	case reflect.Int:
		v, _ = convert.StrToInt(value)
		break
	case reflect.Int8:
		v, _ = convert.StrToInt8(value)
		break
	case reflect.Int16:
		v, _ = convert.StrToInt16(value)
		break
	case reflect.Int32:
		v, _ = convert.StrToInt32(value)
		break
	case reflect.Int64:
		v, _ = convert.StrToInt64(value)
		break
	case reflect.Uint:
		v, _ = convert.StrToUint(value)
		break
	case reflect.Uint8:
		v, _ = convert.StrToUint8(value)
		break
	case reflect.Uint16:
		v, _ = convert.StrToUint16(value)
		break
	case reflect.Uint32:
		v, _ = convert.StrToUint32(value)
		break
	case reflect.Uint64:
		v, _ = convert.StrToUint64(value)
		break
	case reflect.Float32:
		v, _ = convert.StrToFloat32(value)
		break
	case reflect.Float64:
		v, _ = convert.StrToFloat64(value)
		break
	default:
		v = value
	}
	return reflect.ValueOf(v)
}
