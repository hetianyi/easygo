package app

import (
	"github.com/hetianyi/easygo/logger"
	json "github.com/json-iterator/go"
	"os"
	"reflect"
	"strings"
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

	typ := reflect.TypeOf(*c)
	println(typ.Name, " ", typ.Kind(), " ", typ.Kind().String())

	elem := reflect.ValueOf(c).Elem()
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
	if !isPtr(a) {
		panic("无法处理，输入必须是一个struct指针")
	}
	realType, is := isStruct(a)
	if !is {
		panic("无法处理，输入不是一个struct类型")
	}
	val := extractPointer(a)
	for i := 0; i < realType.NumField(); i++ {
		f := realType.Field(i)
		println(f.Name, " ", f.Type.Kind(), " ", f.Type.Kind().String())
		if !isBasicType(f.Type.Kind()) {
			if f.Type.Kind() == reflect.Struct {
				fieldVal := val.Field(i).Interface()

				cc := reflect.ValueOf(&fieldVal).Elem()

				println(val.Field(i).CanAddr())
				println(cc.CanAddr())
				println(reflect.ValueOf(fieldVal).Kind().String())
				println(reflect.TypeOf(fieldVal).Kind().String())

				tmp := reflect.New(reflect.TypeOf(fieldVal)).Elem()
				println(tmp.CanAddr())
				tmp.Set(val.Field(i))
				println(tmp.Field(1).CanAddr())
				tmpIt := tmp.Interface()

				println(reflect.ValueOf(tmpIt).CanAddr())
				println(reflect.ValueOf(&tmpIt).CanAddr())
				println(reflect.ValueOf(&tmpIt).Elem().CanAddr())

				resolveStruct(&tmpIt)
				println(val.Field(i).CanAddr())
				val.Field(i).Set(reflect.ValueOf(tmpIt))
			}
			continue
		}

		tag := strings.TrimSpace(f.Tag.Get("env"))
		if tag == "" {
			continue
		}
		envValue, exist := os.LookupEnv(tag)
		if exist {
			v := convertType(envValue, f.Type.Kind())
			println(val.Field(i).CanAddr())
			val.Field(i).Set(v)
		}
	}
	toString, _ := json.MarshalToString(a)
	println(toString)
}

func TestNewApp2(t *testing.T) {
	ConfigureEnvVariables(&Config{
		Server: &Server{
			Host:         "0.0.0.0",
			Port:         8080,
			GinMode:      "release",
			UseGinLogger: false,
			MiddleWares:  nil,
			Validators:   nil,
			ApiGroup:     nil,
		},
	})
}

func TestNewApp3(t *testing.T) {
	a := &Server{}
	var b interface{}
	b = Server{}
	atyp := reflect.TypeOf(a)
	btyp := reflect.TypeOf(&b)

	atyp = atyp.Elem()
	btyp = btyp.Elem()

	println(atyp.Name, " ", atyp.Kind(), " ", atyp.Kind().String())
	println(btyp.Name, " ", btyp.Kind(), " ", btyp.Kind().String())
}

func extractPointer1(s interface{}) {
	val := reflect.ValueOf(s)
	for {
		if val.Kind() == reflect.Ptr || val.Kind() == reflect.Interface {
			if val.Elem().Kind() != reflect.Ptr && val.Elem().Kind() != reflect.Interface {
				val = val.Elem()
				break
			} else {
				val = val.Elem()
			}
			continue
		}
		break
	}

	realType := reflect.TypeOf(val.Interface())

	for i := 0; i < realType.NumField(); i++ {
		f := realType.Field(i)
		println(f.Name, " ", f.Type.Kind(), " ", f.Type.Kind().String())
		if !isBasicType(f.Type.Kind()) {
			continue
		}
		tag := strings.TrimSpace(f.Tag.Get("env"))
		if tag == "" {
			continue
		}
		envValue, exist := os.LookupEnv(f.Tag.Get("env"))
		if exist {
			v := convertType(envValue, f.Type.Kind())
			val.Field(i).Set(v)
		}
	}
}

func TestNewApp4(t *testing.T) {
	a := &Config{
		Server: &Server{
			Host:         "0.0.0.0",
			Port:         8080,
			GinMode:      "release",
			UseGinLogger: false,
			MiddleWares:  nil,
			Validators:   nil,
			ApiGroup:     nil,
		},
	}

	b := a.Server
	extractPointer1(&b)
}

func TestNewApp5(t *testing.T) {
	var b interface{}
	b = Server{}
	println(reflect.TypeOf(&b).Elem().String())
	println(reflect.TypeOf(&b).Elem().Kind().String())
	println(reflect.ValueOf(&b).Kind().String())
	println(reflect.ValueOf(&b).Elem().Elem().Kind().String())
}

func TestConfigureEnvVariables(t *testing.T) {

	c := &Config{
		Server: &Server{
			Host:         "0.0.0.0",
			Port:         8080,
			GinMode:      "release",
			UseGinLogger: false,
			MiddleWares:  nil,
			Validators:   nil,
			ApiGroup:     nil,
		}}
	ConfigureEnvVariables(c)
	toString, _ := json.MarshalToString(&c)
	println(toString)
}

func TestApp_LoadFromYamlFile(t *testing.T) {
	ap := NewApp()
	err := ap.LoadFromYamlFile("app.yaml")
	if err != nil {
		logger.Fatal(err)
	}
	toString, _ := json.MarshalToString(ap)
	println(toString)
}
