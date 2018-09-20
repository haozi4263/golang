package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	r.GET("/user/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get user  info succ",
		})
	})
	r.POST("/user/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "create user info succ",
		})
	})
	r.PUT("/user/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "update user info succ",
		})
	})
	r.DELETE("/user/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete user info succ",
		})
	})
	r.Run()
}

