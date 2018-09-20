package main

import (
	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"message": "login succ",
	})
}

func read(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"message": "read succ",
	})
}
func submit(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"message": "submit succ",
	})
}

func main()  {
	route := gin.Default()

	v1 := route.Group("/v1")
	{
		v1.POST("/login", login)
		v1.POST("/submit", submit)
		v1.POST("/read", read)
	}

	v2 := route.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
		v2.POST("/read", read)
	}
	route.Run()
}

