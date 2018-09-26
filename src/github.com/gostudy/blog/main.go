package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gostudy/blog/dal/db"
	"github.com/gostudy/blog/controller"
	"fmt"
)

func main()  {
	route := gin.Default()

	dns := "golang:111111@tcp(192.168.10.105)/golang?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}

	route.Static("/static/", "./static")
	route.LoadHTMLGlob("views/*")

	route.GET("/", controller.IndexHandle)
	route.GET("/article/new/", controller.NewArticle)
	route.POST("/article/submit/", controller.ArticleSubmit)
	route.GET("/article/detail", controller.ArticleDetail)






	fmt.Printf("正在启动服务器...")
	route.Run(":8080")
}