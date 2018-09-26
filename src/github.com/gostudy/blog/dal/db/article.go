package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gostudy/blog/model"
	//"database/sql"
	"fmt"
)

func InsertArticle(article *model.ArticleDetail) (articelId int64, err error)  {
	if article == nil {
		err = fmt.Errorf("invalid article parameter")
		return
	}

	sqlstr := `insert into
			   article(content, summary, title, username,
				category_id, view_count, comment_count)	
		values(?, ?, ?, ?, ?, ?, ?)	`
	result, err := DB.Exec(sqlstr, article.Content, article.Summary,
		article.Title, article.Username, article.ArticleInfo.CategoryId,
		article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		return
	}
	articelId, err = result.LastInsertId()
	return

    //fmt.Println(result)
	return
}

func GetArticleList(pageNum, pageSize int) (articleInfoList []*model.ArticleInfo, err error)  {  //获取文章列表
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid paramater, page_num:%d, page_size:%d\n", pageNum, pageSize)
		return
	}

	sqlstr := ` select 
				id, summary, title, view_count, create_time, comment_count
				username
				from
				article
				where status = 1
				order by create_time desc
				limit ?,?`
	err = DB.Select(&articleInfoList, sqlstr, pageNum, pageSize)
	return
}

func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error)  {
	if articleId < 0 {
		err = fmt.Errorf("invalid parmeter, article_id:%v\n",articleId)
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlstr := `select
				id, summary, title, view_count, content,
				create_time, comment_count, username, category_id
				from article
				where id = ?
				and status = 1
				`
	err = DB.Get(articleDetail, sqlstr, articleId)
	fmt.Printf("article_info:%#v\n", articleDetail)
	return
} //根据id获取文章信息

func GetRelativeArticle(articleId int64) (articleList []*model.RelativeArticle, err error)  {
	var categoryId int64
	sqlstr := "select category_id from article where id=?"
	err = DB.Get(&categoryId, sqlstr,articleId )
	if err != nil {
		return
	}
	sqlstr = "select id ,title from article where category_id=? and id !=? limit 10"
	DB.Select(&articleList, sqlstr, categoryId, articleId)
	return  //获取相关文章
}