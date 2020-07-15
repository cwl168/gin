package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//静态文件夹
// 进入router_static目录  build -o router_static && ./router_static  访问 http://localhost:8080/assets/a.html
func main() {

	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}