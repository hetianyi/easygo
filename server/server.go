package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/hetianyi/easygo/app"
	"github.com/hetianyi/easygo/logger"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"time"
)

func Start(serverConfig *app.Server) error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	gin.ForceConsoleColor()
	if serverConfig.UseGinLogger {
		r.Use(gin.Logger())
	}
	r.Use(gin.Recovery())

	if vrs, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for _, v := range serverConfig.Validators {
			validator := getValidatorByName(v)
			if validator == nil {
				panic("server start failed due to: validator not found: \"" + v + "\", are you forget to register it?")
			}
			vrs.RegisterValidation(v, validator)
		}
	}
	for _, m := range serverConfig.MiddleWares {
		mw := getMiddleWareByName(m)
		if mw == nil {
			panic("server start failed due to: middleware not found: \"" + m + "\", are you forget to register it?")
		}
		r.Use(mw())
	}

	for _, g := range serverConfig.ApiGroup {
		addGroup(r, nil, g)
	}

	listen := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)

	srv := &http.Server{
		Handler: r,
		Addr:    listen,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("starting server error: %s\n", err)
		}
	}()

	logger.Info("http server is listening on: ", listen)

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Info("Server exiting")

	return nil
}

func addGroup(r *gin.Engine, group *gin.RouterGroup, g app.ApiGroup) {
	if group == nil {
		var mws = getGroupHandlers(g.MiddleWares)
		group = r.Group(g.Prefix, mws...)
	}

	for _, a := range g.Apis {
		method, url, err := parseRoute(a.Route)
		if err != nil {
			panic(err)
		}
		handlers := getApiHandlers(a.Handlers)
		group.Handle(method, url, handlers...)
	}
	for _, subGroup := range g.ApiGroup {
		var smws = getGroupHandlers(g.MiddleWares)
		sg := group.Group(subGroup.Prefix, smws...)
		addGroup(r, sg, subGroup)
	}
}

var allowedHttpRouteMethod = map[string]bool{
	"post":    true,
	"get":     true,
	"delete":  true,
	"patch":   true,
	"put":     true,
	"options": true,
	"head":    true,
	"trace":   true,
	"connect": true,
}

var (
	routePattern = regexp.MustCompile("(?i)^\\s*(get|post|delete|put|trace|connect|options|head|patch)\\s+/(.*)\\s*$")
)

func parseRoute(route string) (method string, url string, err error) {
	if !routePattern.MatchString(route) {
		err = errors.New("invalid route: " + route + ", route pattern must match '<Method> <Url>'")
		return
	}
	method = strings.ToUpper(routePattern.ReplaceAllString(route, "$1"))
	url = "/" + routePattern.ReplaceAllString(route, "$2")
	return
}

func getGroupHandlers(middlewares []string) []gin.HandlerFunc {
	var mws []gin.HandlerFunc
	for _, m := range middlewares {
		mw := getMiddleWareByName(m)
		if mw == nil {
			panic("server start failed due to: middleware not found: \"" + m + "\", are you forget to register it?")
		}
		mws = append(mws, mw())
	}
	return mws
}

func getApiHandlers(handlerNames []string) []gin.HandlerFunc {
	var handlers []gin.HandlerFunc
	for _, n := range handlerNames {
		handler := getApiHandlerByName(n)
		if handler == nil {
			logger.Warn("api handler not found: \"" + n + "\", are you forget to register it?")
			continue
		}
		handlers = append(handlers, handler)
	}
	return handlers
}
