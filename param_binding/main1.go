package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Person1 struct {
	Name    string `form:"name"`
	Address string `form:"address"`
	//JoiningDate    time.Time `form:"joining_date" time_format:"2006-01-02 15:04:05" time_utc:"1"`
	JoiningDate time.Time `form:"joining_date" time_format:"2006-01-02 15:04:05" time_utc:"1"`
}

//curl -X POST localhost:8085/testing --header 'Content-Type: application/json' --data-raw '{"name":"cwl","address":"sh","joining_date":"2020-03-03"}'
func main() {
	route := gin.Default()
	route.GET("/testing", startPag1e)
	route.POST("/testing", startPag1e)
	route.Run(":8085")
}

func startPag1e(c *gin.Context) {
	var person Person1
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if error := c.ShouldBind(&person); error == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.JoiningDate)
		c.String(200, "%v\n", person)
	} else {
		c.String(200, "%v\n", error)
	}

	c.String(200, fmt.Sprintf("%#v", person))
}
