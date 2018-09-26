package logic

import (
	"github.com/gostudy/blog/model"
	"github.com/gostudy/blog/dal/db"
	"fmt"
	"math"
)

func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64)  {
LABEL:
	for _, article := range articleInfoList {
		categoryId := article.CategoryId
		for _, id := range ids {
			if id == categoryId {
				continue LABEL
			}
		}
		ids = append(ids, categoryId)
	}
	return
}


func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error)   {
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		fmt.Printf("1 get article list failed, err:%v\n", err)
		return
	}

	if len(articleInfoList) == 0 {
		return
	}

	categoryIds := getCategoryIds(articleInfoList)

	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("2 get catetoryid failed, err:%v\n", err)
		return
	}


	//聚合数据
	for _, article := range articleInfoList {
		fmt.Printf("content:%s\n", article.Summary)
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}


/*
	content := c.PostForm("content")
	author := c.PostForm("author")
	categoryId := c.Params("category_id")
	title := c.PostForm("title")
 */

func InsertArticle(content, author, title string, categoryId int64 ) (err error)   {
	//构照Articledetail的类
	articleDetail := &model.ArticleDetail{}
	articleDetail.Content = content
	articleDetail.Username = author
	articleDetail.Title = title
	articleDetail.ArticleInfo.CategoryId = categoryId
	//摘要内容截取128个字，将用户提交的content内容utf转成utf8切片
	contentUtf8 := []rune(content)
	minLength := int(math.Min(float64(len(contentUtf8)), 128.0))
	articleDetail.Summary = string([]rune(content)[:minLength])

	article_id, err := db.InsertArticle(articleDetail)
	fmt.Printf("insert articl succ, id:%d, err:%v\n", article_id, err)
	return
}

func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error)  {
	//1.获取文章的信息
	articleDetail, err = db.GetArticleDetail(articleId)
	if err != nil {
		return
	}

	//2.获取文章对应的分类信息
	category, err := db.GetCategoryById(articleDetail.ArticleInfo.CategoryId) //指针类型
	if err != nil {
		fmt.Printf("GetCategoryById failed:%v\n", err)
		return
	}
	articleDetail.Category = *category  //赋值结构图category为值类型需要加*
	return
}

func GetRelativeArticleList(articleId int64) (articleList []*model.RelativeArticle, err error)  {
	articleList, err = db.GetRelativeArticle(articleId)
	return
}



