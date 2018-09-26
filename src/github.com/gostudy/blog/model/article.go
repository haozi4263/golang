package model

import "time"

type ArticleInfo struct {  //文章详情
	Id 				int64		`db:"id"`
	CategoryId		int64		`db:"category_id"`  //分类id
	Summary			string		`db:"summary"`		//摘要
	Title			string		`db:"title"`		//标题
	ViewCount		uint32		`db:"view_count"`	//访问次数
	Createtime		time.Time	`db:"create_time"`	//创建时间
	CommentCount	uint32		`db:"comment_count"`	//评论次数
	Username		string		`db:"username"`
}

type ArticleDetail struct {
	ArticleInfo		//匿名字段
	Content			string		`db"content"`  //内容
	Category		//分类
}

type ArticleRecord struct {  //文章记录
	ArticleInfo
	Category
}