package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.GET("/moreXML", func(c *gin.Context) {
		type MessageRecord struct {
			Name string
			Message string
			Number  int
		}
		var msg MessageRecord
		msg.Name = "jude"
		msg.Message = "hello"
		msg.Number = 11
		c.XML(http.StatusOK, msg)
	})
	r.Run()
}

