package logic

import (
	"github.com/gostudy/blog/model"
	"github.com/gostudy/blog/dal/db"
	"fmt"
)


func GetAllCategoryList() (allcategoryList []*model.Category, err error) {
	//1. 从数据库中，获取文章分类列表
	allcategoryList, err = db.GetAllCategoryList()
	if err != nil {
		fmt.Printf("1 get article list failed, err:%v\n", err)
		return
	}

	return
}

func GetCategoryList() (categoryList []*model.Category, err error)  {
	//从数据库中获取所有分类列表
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get categorylist failed, err:%v\n", err)
		return
	}
	return
}