package model

type Category struct {  //分类
	CategoryId 		int64		`db:"id"`  //分类id
	CategoryName	string		`db:"category_name"`  //分类名字
	CategoryNo 		int			`db:"category_no"`   //编号
}

