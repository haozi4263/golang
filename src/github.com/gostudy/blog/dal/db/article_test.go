package db

import (
	"github.com/gostudy/blog/model"
	"testing"
	"time"
	"fmt"
)

func init()  {
	dns := "golang:111111@tcp(192.168.242.129:3306)/golang?parseTime=true"
	err :=Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertArticle(t *testing.T) {  //插入文章
	a1 := &model.ArticleDetail{}
	a1.ArticleInfo.CategoryId = 1
	a1.ArticleInfo.CommentCount = 0
	a1.Content = "测试插入函数,测试插入函数,测试插入函数"
	a1.ArticleInfo.Createtime = time.Now()
	a1.ArticleInfo.Summary = `测试插入函数`
	a1.ArticleInfo.Title = "insertarticel"
	a1.ArticleInfo.Username = "jude"
	a1.ArticleInfo.ViewCount = 1
	a1.Category.CategoryId = 1

	articelId, err := InsertArticle(a1)
	if err != nil {
		t.Errorf("insert articel failed, err:%v\n", err)
		return
	}
	t.Logf("insert article succ, articleId:%d\n", articelId)
}

func TestGetArticelList(t *testing.T) {
	articleList, err := GetArticleList(0,15)
	if err != nil {
		t.Errorf("get articleList failed, err:%v\n", err)
		return
	}
	t.Logf("get articlelist succ, len:%d\n",len(articleList))
	for _, v := range articleList {
		fmt.Println(v)
	}
}

func TestGetArticleDetail(t *testing.T) {
	a2 ,err := GetArticleDetail(3)
	if err != nil {
		t.Errorf("get articledetail failed, err:%v\n", err)
		return
	}
	t.Logf("get articledetail succ, articledetail:%#v\n", a2)
}

func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		t.Errorf("GetCategoryById is failed:%d\n", err)
		return
	}
	t.Logf("categor:%#v\n", category)
}

func TestGetRelativeArticle(t *testing.T) {
	articleList, err := GetRelativeArticle(1)
	if err != nil {
		t.Errorf("get relative arvicle failed, err:%v\n", err)
		return
		}
	for _, v := range articleList {
		t.Logf("id:%d title:%s\n", v.ArticleId,v.Title )
	}
	}