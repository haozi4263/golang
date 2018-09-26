package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gostudy/blog/logic"
	"fmt"
	"net/http"
	"strconv"
)

func IndexHandle(c *gin.Context)  {
	articleRecordList, err := logic.GetArticleRecordList(0 , 15)
	if err != nil {
		fmt.Printf("get article failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	allCategoryList, err := logic.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
	}

	var data map[string]interface{} = make(map[string]interface{}, 10)
	data["article_list"] = articleRecordList
	data["category_list"] = allCategoryList
	c.HTML(http.StatusOK, "views/index.html", data)

}

func NewArticle(c *gin.Context)  {
	categoryList, err := logic.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/post_article.html", categoryList)  //传入切片进模版
}

func ArticleSubmit(c *gin.Context)  {
	content := c.PostForm("content")
	author := c.PostForm("author")
	categoryIdStr := c.PostForm("category_id")  //字符串
	title := c.PostForm("title")

	categoryId, err := strconv.ParseInt(categoryIdStr, 10 , 64)
	if err != nil {
		fmt.Printf("参数错误", err)
		return
	}
	err = logic.InsertArticle(content, author, title, categoryId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")  //301重定向
}

func ArticleDetail(c *gin.Context)  {  //处理文章详情页
	articleIdStr := c.Query("article_id")  //字符串

	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	articleDetail, err := logic.GetArticleDetail(articleId)
	if err != nil {
		fmt.Printf("get article detail failed,article_id:%d err:%v\n", articleId, err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	relativeArticle, err := logic.GetRelativeArticleList(articleId)
	if err != nil {
		fmt.Printf("get relativeArticle failed, err:%v\n", err)
	}
	//定义map并初始化用于存储relativeArticle和articleDetail，因模版只能传1个变量
	var m map[string]interface{} = make(map[string]interface{}, 10)
	m["detail"] = articleDetail
	m["relative_article"] = relativeArticle


	//c.HTML(http.StatusOK, "views/detail.html", articleDetail)
	c.HTML(http.StatusOK, "views/detail.html", m)
	return
}


























