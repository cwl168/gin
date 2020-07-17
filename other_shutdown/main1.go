package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
测试运行：
1 go run main.og
2 curl "localhost:8085" 请求
3 Ctrl + C 停止进程
4 curl 不会处理请求  返回curl: (52) Empty reply from server
*/
//非优雅关停服务

func main() {
	router := gin.Default()
	//设置路由
	router.GET("/", func(c *gin.Context) {
		//模拟超时请求
		time.Sleep(5 * time.Second)
		fmt.Println("Welcome Gin Server11111")
		c.String(http.StatusOK, "Welcome Gin Server22222")
	})
	router.Run(":8085")

}
