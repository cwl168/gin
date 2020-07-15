package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

//获取body内容  curl -H "Content-Type: application/json" -X POST -d '{"user_id": "123", "coin":100, "success":1, "msg":"OK!!!" }' "http://localhost:8080/user/12action"
func main()  {
	router := gin.Default()

	router.Any("/user/*action", func(c *gin.Context) {
		bodyBytes,err:=ioutil.ReadAll(c.Request.Body)
		if err!=nil{
			c.String(http.StatusBadRequest, err.Error())
		}
		c.String(http.StatusOK, string(bodyBytes))
	})

	router.Run(":8080")
}
