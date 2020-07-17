package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//模板渲染
//1 进入other_template目录，运行go run main.go
//2 curl -X GET "localhost:8085/index"
func main() {
	router := gin.Default()
	//router.LoadHTMLFiles("templates/index.html",)
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8085")
}
