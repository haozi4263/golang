package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(200, gin.H{
			"message": "pong",
			"username": username,
			"address": address,
		})
	})
	r.Run(	)
}


//run http://localhost:8080/user/search/abc/wuhan
// http://localhost:8080/user/search?username=jude&address=wuhan