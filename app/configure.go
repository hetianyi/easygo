package app

import (
	"errors"
	"github.com/hetianyi/easygo/convert"
	"os"
	"reflect"
	"strings"
)

// ConfigureEnvVariables 将环境变量值填充/覆盖到给定的struct配置。
// 如果需要从环境变量中读取配置，则需要将struct的字段加上env标签，例如：
//  type Server struct {
//  	Port int `env:"HTTP_PORT"`
//  }
// 注意：参数必须是struct的指针
func ConfigureEnvVariables(a interface{}) error {
	if !isPtr(a) {
		return errors.New("input argument must be a POINTER of a struct")
	}
	realType, is := isStruct(a)
	if !is {
		return errors.New("input argument must be a pointer of a STRUCT")
	}
	val := extractPointer(a)
	configure(val, realType)
	return nil
}

func configure(val reflect.Value, realType reflect.Type) reflect.Value {
	for i := 0; i < val.NumField(); i++ {
		f := realType.Field(i)
		//println(f.Name, " ", f.Type.Kind(), " ", f.Type.Kind().String())
		if !isBasicType(f.Type) {
			if isFieldStructType(val.Field(i).Type()) {
				tmp := reflect.New(reflect.TypeOf(val.Field(i).Interface())).Elem()
				//println(tmp.CanAddr())
				tmp.Set(val.Field(i))
				//println(tmp.Field(1).CanAddr())
				configure(tmp, reflect.TypeOf(val.Field(i).Interface()))
				val.Field(i).Set(tmp)
			}
			continue
		}
		resolveBasicField(f, val.Field(i))
	}
	return val
}

func resolveBasicField(ft reflect.StructField, fv reflect.Value) {
	tag := strings.TrimSpace(ft.Tag.Get("env"))
	if tag == "" {
		return
	}
	envValue, exist := os.LookupEnv(ft.Tag.Get("env"))
	if exist {
		v := convertType(envValue, ft.Type.Kind())
		fv.Set(v)
	}
}

func isPtr(a interface{}) bool {
	typ := reflect.TypeOf(a)
	return typ.Kind() == reflect.Pointer
}

func isStruct(s interface{}) (reflect.Type, bool) {
	val := reflect.ValueOf(s)
	for {
		if val.Kind() == reflect.Ptr || val.Kind() == reflect.Interface {
			val = val.Elem()
			continue
		}
		break
	}
	return val.Type(), val.Kind() == reflect.Struct
}

func extractPointer(s interface{}) reflect.Value {
	val := reflect.ValueOf(s)
	for {
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
			continue
		}
		break
	}
	if val.Kind() == reflect.Interface {
		return val.Elem()
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

func isBasicType(t reflect.Type) bool {
	if t.Kind() == reflect.Ptr {
		return isBasicType(t.Elem())
	}
	k := t.Kind()
	return k == reflect.Bool ||
		k == reflect.Int ||
		k == reflect.Int8 ||
		k == reflect.Int16 ||
		k == reflect.Int32 ||
		k == reflect.Int64 ||
		k == reflect.Uint ||
		k == reflect.Uint8 ||
		k == reflect.Uint16 ||
		k == reflect.Uint32 ||
		k == reflect.Uint64 ||
		k == reflect.Float32 ||
		k == reflect.Float64 ||
		k == reflect.String
}

func isFieldStructType(t reflect.Type) bool {
	if t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
		return isFieldStructType(t.Elem())
	}
	return t.Kind() == reflect.Struct
}
