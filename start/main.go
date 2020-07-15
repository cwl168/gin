package main

import "github.com/gin-gonic/gin"


//go run main.go启动  浏览器访问 http://localhost:8080/ping
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}