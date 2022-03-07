package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hetianyi/easygo/logger"
	"reflect"
	"runtime"
	"strings"
	"sync"
)

var (
	middlewares = map[string]func() gin.HandlerFunc{}
	mu          sync.Mutex
)

// RegisterMiddleWare 注册一个gin MiddleWare
func RegisterMiddleWare(middleWare func() gin.HandlerFunc) {
	mu.Lock()
	defer mu.Unlock()

	name := getFuncName(middleWare)
	if middlewares[name] != nil {
		panic("MiddleWare register failed due to: MiddleWare \"" + name + "\" already registered.")
	}
	logger.Debug("register MiddleWare :: ", name)
	middlewares[name] = middleWare
}

func getMiddleWareByName(name string) func() gin.HandlerFunc {
	return middlewares[name]
}

func getFuncName(f interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	index := strings.LastIndex(fullName, ".")
	return fullName[index+1:]
}
