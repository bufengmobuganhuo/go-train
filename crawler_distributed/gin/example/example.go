package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":"pong",
		})
	})
	// 监听并在localhost:8080上启动服务
	r.Run()
}