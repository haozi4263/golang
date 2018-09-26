package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gostudy/blog/model"
	"github.com/jmoiron/sqlx"
)

//func InsertCategory(category *model.Category) (categoryId int64, err error) {
//	sqlstr := "insert into category(category_name, category_no) values(?, ?)"
//	result, err := DB.Exec(sqlstr, category.CategoryName, category.CategoryId)
//	if err != nil {
//		fmt.Printf("insert category failed err:%v\n", err)
//		return
//	}
//	categoryId, err = result.LastInsertId()
//	return
//}

func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error)  {
	sqlstr, args, err := sqlx.In("select id, category_name, category_no from category where id in(?)", categoryIds)
	if err != nil {
		return
	}
	err = DB.Select(&categoryList, sqlstr, args...)
	return
}

func GetAllCategoryList() (categoryList []*model.Category, err error) {

	sqlstr := "select id, category_name, category_no from category order by category_no asc"
	err = DB.Select(&categoryList, sqlstr)
	return
}

func GetCategoryById(id int64) (category *model.Category, err error)  {  //根据id获取分类信息
	category = &model.Category{}  //初始化分配内存
	sqlstr := "select id, category_name, category_no from category where id = ?"

	err = DB.Get(category, sqlstr, id)
	return
}

