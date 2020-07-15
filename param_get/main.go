package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取Get参数
func main()  {
	router := gin.Default()

	router.GET("/user/*action", func(c *gin.Context) {
		firstName:=c.DefaultQuery("first_name","wang")
		lastName:=c.DefaultQuery("last_name","kai")
		c.String(http.StatusOK, "%s,%s",firstName,lastName)
	})

	router.Run(":8080")
}
