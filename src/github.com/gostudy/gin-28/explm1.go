package main

import "github.com/gin-gonic/gin"

func testPing(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main()  {
	//default fanhui yige moren de luyou yinqin
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", testPing)

	//r.Run()
	r.Run(":9090")
}