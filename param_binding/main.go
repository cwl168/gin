package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

//curl -X GET 'http://localhost:8085/testing?name=cwl&address=sh&birthday=2020-01-04'
//curl -X POST -d 'name=cwl&address=sh&birthday=2020-01-04' 'http://localhost:8085/testing'
//curl -X POST -H "Content-Type: application/json" -d '{"name":"cwl","address":"sh","birthday":"2006-01-02"}' 'http://localhost:8085/testing'
type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

//获取绑定Get参数或者Post参数
func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.POST("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	//根据请求content-type来做不同的binding
	if error := c.ShouldBind(&person); error == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)

		c.String(200, "%v\n", person)
	} else {
		c.String(200, "%v\n", error)
	}
	c.String(200, fmt.Sprintf("%#v", person))
}
