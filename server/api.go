package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hetianyi/easygo/logger"
)

var (
	apis = map[string]gin.HandlerFunc{}
)

// RegisterApiHandler 注册一个gin API处理器
func RegisterApiHandler(apiHandler gin.HandlerFunc) {
	mu.Lock()
	defer mu.Unlock()

	name := getFuncName(apiHandler)
	if apis[name] != nil {
		panic("Api Handler register failed due to: Api Handler \"" + name + "\" already registered.")
	}
	logger.Debug("register Api Handler :: ", name)
	apis[name] = apiHandler
}

func getApiHandlerByName(route string) gin.HandlerFunc {
	return apis[route]
}
