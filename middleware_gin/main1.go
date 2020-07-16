package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	//日志信息写入到文件中
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New()
	//使用Logger中间件 实际作用打印请求的url  如果使用Recovery中间件 panic错误不会导致程序挂
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "defalue_name")
		panic("panic test")
		c.String(200, "%s", name)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
