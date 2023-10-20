package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("for", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "bar",
		})
	})
	r.Run() // 默认监听并在 127.0.0.1:8080 上启动服务
}
