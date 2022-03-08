package reflectx

import (
	"github.com/hetianyi/easygo/convert"
	"reflect"
)

// Server 为http服务器的配置
type Server struct {
	Host         string `yaml:"host" env:"HTTP_HOST"`                   // HTTP服务监听地址
	Port         int    `yaml:"port" env:"HTTP_PORT"`                   // HTTP服务监听端口
	GinMode      string `yaml:"mode" env:"HTTP_GIN_MODE"`               // 服务启动模式：debug|release
	UseGinLogger bool   `yaml:"useGinLogger" env:"HTTP_USE_GIN_LOGGER"` // 是否开启gin的默认Logger
}

type Config struct {
	Server *Server
	Age    int `yaml:"useGinLogger" env:"AGE"`
}

func Start(a interface{}) {
	Configure(reflect.ValueOf(a))
}

func Configure(a reflect.Value) {
	if a.Kind() == reflect.Ptr {
		println(a.Kind().String())
		Configure(a.Elem())
		return
	}
	println(a.Type().String(), a.Kind().String(), a.CanAddr())
	if a.Kind() == reflect.Struct {
		IteratorStruct(a)
	}
}

func IteratorStruct(a reflect.Value) {
	for i := 0; i < a.NumField(); i++ {
		f := a.Field(i)
		Configure(f)
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
