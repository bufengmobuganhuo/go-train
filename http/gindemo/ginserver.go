package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const requestIdKey = "requestId"

func main() {
	// 返回一个Engine
	r := gin.Default()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	r.Use(func(ctx *gin.Context) {
		s := time.Now()
		// 记录path, 类似一个before的切面
		logger.Info("incoming request", zap.String("path", ctx.Request.URL.Path))
		// 继续执行，这个用来区分切面的执行时间，多个middleware可以多次调用
		ctx.Next()
		// 记录response code, after的切面
		logger.Info("incoming request", zap.Int("status", ctx.Writer.Status()),
			zap.Duration("elapsed", time.Since(s)))
	}, func(ctx *gin.Context) {
		// before的切面
		logger.Info("set requestId")
		// 在请求中添加新键值对
		ctx.Set(requestIdKey, rand.Int())
		ctx.Next()
	})

	// 定义一个接受GET的方法
	r.GET("/ping", func(ctx *gin.Context) {
		logger.Info("request reached")
		h := gin.H{
			"message": "pong",
		}
		// 获取设置的键值对
		if rid, exists := ctx.Get(requestIdKey); exists {
			// 返回添加新键值对
			h[requestIdKey] = rid
		}
		ctx.JSON(200, h)
	})
	r.GET("/hello", func(ctx *gin.Context) {
		logger.Info("request reached")
		ctx.String(200, "hello")
	})

	// 启动服务器
	r.Run()
}
