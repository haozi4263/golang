package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main()  {
	route := gin.Default()
	route.LoadHTMLGlob("tmp/**/*")
	route.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "jude",
		})
	})
	route.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "zhanghoa",
		})
	})
	route.Run()
}