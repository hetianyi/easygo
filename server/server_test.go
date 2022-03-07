package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hetianyi/easygo/app"
	"github.com/hetianyi/easygo/logger"
	"net/http"
	"reflect"
	"runtime"
	"testing"
)

const (
	trackingCookieName = "_trkid"
)

func init() {
	RegisterApiHandler(V1Demo1)
	RegisterApiHandler(V1Demo2)
	RegisterApiHandler(V2Demo1)
	RegisterApiHandler(V2Demo2)
	RegisterApiHandler(V2V1Demo1)
	RegisterApiHandler(V2V1Demo2)

	RegisterMiddleWare(MiddleWareCommon)
	RegisterMiddleWare(MiddleWareV1)
	RegisterMiddleWare(MiddleWareV2)
}

func MiddleWareCommon() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("MiddleWareCommon")
		c.Next()
	}
}
func MiddleWareV1() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("MiddleWareV1")
		c.Next()
	}
}
func MiddleWareV2() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("MiddleWareV2")
		c.Next()
	}
}

func V1Demo1(c *gin.Context) {
	c.String(http.StatusOK, "v1 %d", 1)
}
func V1Demo2(c *gin.Context) {
	c.String(http.StatusOK, "v1 %d", 2)
}
func V2Demo1(c *gin.Context) {
	c.String(http.StatusOK, "v2 %d", 1)
}
func V2Demo2(c *gin.Context) {
	c.String(http.StatusOK, "v2 %d", 2)
}
func V2V1Demo1(c *gin.Context) {
	c.String(http.StatusOK, "v2>v1 %d", 3)
}
func V2V1Demo2(c *gin.Context) {
	c.String(http.StatusOK, "v2>v2 %d", 4)
}

func TestStart(t *testing.T) {
	newApp := app.NewApp()
	err := newApp.LoadFromYamlFile("app.yaml")
	if err != nil {
		logger.Fatal(err)
	}
	Start(&newApp.Config.Server)
}

func TestStart2(t *testing.T) {
	logger.Info(runtime.FuncForPC(reflect.ValueOf(MiddleWareCommon).Pointer()).Name())
	logger.Info(reflect.TypeOf(MiddleWareCommon))
	logger.Info(reflect.TypeOf(MiddleWareCommon).Name())
}
