package reflectx

import (
	"github.com/hetianyi/easygo/convert"
	"os"
	"reflect"
	"strings"
)

// Server 为http服务器的配置
type Server struct {
	Host         string `yaml:"host" env:"HTTP_HOST"`                   // HTTP服务监听地址
	Port         int    `yaml:"port" env:"HTTP_PORT"`                   // HTTP服务监听端口
	GinMode      string `yaml:"mode" env:"HTTP_GIN_MODE"`               // 服务启动模式：debug|release
	UseGinLogger bool   `yaml:"useGinLogger" env:"HTTP_USE_GIN_LOGGER"` // 是否开启gin的默认Logger
}

type Config struct {
	Age    *int `yaml:"useGinLogger" env:"AGE"`
	Server *Server
}

func Start(a interface{}) {
	Configure(reflect.ValueOf(a))
}

func Configure(a reflect.Value, field ...reflect.StructField) {
	if a.Kind() == reflect.Ptr {
		println(a.Kind().String())
		Configure(a.Elem(), field...)
		return
	}
	//println(a.Type().String(), a.Kind().String(), a.CanAddr())
	if a.Kind() == reflect.Struct {
		IteratorStruct(a)
		return
	}
	println(a.String())
	println(a.Kind().String())
	if isBasicType(field[0].Type.Kind()) {
		resolveBasicField(a, field[0])
	}
}

func resolveBasicField(a reflect.Value, field reflect.StructField) {
	tag := strings.TrimSpace(field.Tag.Get("env"))
	if tag == "" {
		return
	}
	envValue, exist := os.LookupEnv(field.Tag.Get("env"))
	if exist {
		v := convertType(envValue, field.Type.Kind())
		a.Set(v)
	}
}

func IteratorStruct(a reflect.Value) {
	typ := a.Type()
	for i := 0; i < a.NumField(); i++ {
		f := a.Field(i)
		println(f.String())
		Configure(f, typ.Field(i))
	}
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

func isBasicType(k reflect.Kind) bool {
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
