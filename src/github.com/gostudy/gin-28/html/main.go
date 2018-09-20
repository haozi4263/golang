package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	route := gin.Default()
	//route.Static("/img", "./static")
	route.Run()
}

