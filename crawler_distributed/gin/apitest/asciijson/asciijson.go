package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/someJSON", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			// 返回时需要对'<'转义
			"tag": "<br>",
		}
		// 可以生成非ASCII编码的字符
		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		ctx.AsciiJSON(http.StatusOK, data)
	})
	r.Run(":8081")
}
