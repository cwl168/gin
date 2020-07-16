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
//parsing time ""2006-01-02 15:04:05"" as ""2006-01-02T15:04:05Z07:00"": cannot parse " 15:04:05"" as "T"

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
	//time_format is designed for custom tags such as form, query, header, etc., does not contain json, yaml library
	//Birthday time.Time `form:"birthday" time_format:"2006-01-02 15:04:05" time_utc:"1"`
	JoiningDate time.Time `form:"joining_date"  json:"joining_date" time_format:"2006-01-02 15:04:05" time_utc:"1"`
	//JoiningDate time.Time `form:"joining_date"  time_format:"2006-01-02 15:04:05" time_utc:"1"`  //接受不了joining_date参数，必须joining_date
}

//获取绑定Get参数或者Post参数
//curl -X POST localhost:8085/testing --header 'Content-Type: application/json' --data-raw '{"name":"cwl","address":"sh","joining_date":"2019-03-29T09:32:52Z"}'  RFC3339 格式才不会报错
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
		//log.Println(person.Birthday)
		//2020-03-29 09:32:52 +0000 UTC转化为标准时间格式
		log.Println(person.JoiningDate.Format("2006-01-02 15:04:05"))
		c.String(200, "%v\n", person)
	} else {
		c.String(200, "%v\n", error)
	}

	c.String(200, fmt.Sprintf("%#v", person))
}
