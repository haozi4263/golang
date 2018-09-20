package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context) {
		//fangfa1:
		c.JSON(http.StatusOK, gin.H{"messsage": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		//fangfa2
		var msg struct{
			Name string `json:"user"`
			Message string
			Number int
		}
		msg.Name = "jude"
		msg.Message = "hey"
		msg.Number = 18
		c.JSON(http.StatusOK, msg)
	})
	r.Run()
}

