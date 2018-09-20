package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gostudy/blog-29/dal/db"
	"github.com/gostudy/blog-29/controller"
	"fmt"
)

func main()  {
	route := gin.Default()
	dns := "golang:111111@tcp(192.168.10.105:3306)/golang?parseTime=true"  //将mysql时间字段解析成go语言时间类型，默认不开启
	err := db.Init(dns)
	fmt.Printf("mysql conn succ")
	if err != nil {
		panic(err)
	}

	route.Static("/static/", "./static")
	route.LoadHTMLGlob("views/*")  //加载模板

	route.GET("/", controller.IndexHandle)
	//发布页面
	//route.GET("/article/new/", controller.NewArticle)
	////文章提交接口
	//route.POST("/article/submit/", controller.ArticleSubmit)
	////文章上传接口
	//route.POST("/upload/file/", controller.UploadFile)
	//
	////留言页面
	//route.GET("/leave/new/", controller.LeaveNew)
	////关于我页面
	//route.GET("/about/me/", controller.AboutMe)
	fmt.Printf("正在启动服务器....\n")
	route.Run(":8080")

}

