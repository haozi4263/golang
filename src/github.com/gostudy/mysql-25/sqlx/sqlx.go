package main

import (
	//"database/sql"  //sql接口
	_"github.com/go-sql-driver/mysql"  //导入mysql驱动,执行init函数，进行注册驱动
	"github.com/jmoiron/sqlx"
	"fmt"
	"database/sql"
)

var DB *sqlx.DB  //从database/sql中导入
func initDb() error  {
	var err error
	dsn := "golang:111111@tcp(192.168.10.105:3306)/golang"
	DB, err = sqlx.Open("mysql", dsn)  //open函数调用mysql-driver调用map
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)  //最大连接
	DB.SetMaxIdleConns(16)  //空闲连接
	return  nil

}


type User struct {
	Id int64 `db:"id"`   //根据数据库字段类型动态选择字段对应类型，数据库id为int
	Name sql.NullString `db:"name"`  //当name字段为空时需哟啊加上
	Age int `db:"age"`
}



func testQuery()  {  //单行查询
	sqlstr := "select id, name, age from user where id=?"
	var user User

	err := DB.Get(&user, sqlstr, 2)
	if err != nil {
		fmt.Printf("get failed ,err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", user)
}

func testQueryMulti()  {
	sqlstr := "select id, name, age from user where id>?"
	var user []User //传入切片

	err := DB.Select(&user, sqlstr, 1)
	if err != nil {
		fmt.Printf("get failed ,err:%v\n", err)
		return
	}
	fmt.Printf("user:#%v\n", user)
}

func testUpdateMulti()  {
	sqlstr := "UPDATE user set  name=? WHERE id=?"
	result, err := DB.Exec(sqlstr, "zhanghao", 1)
	if err != nil {
		fmt.Printf("update is failed, err:%v\n", result)
		return
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("RowsAffected is failed, err:%v\n", err)
		return
	}
	fmt.Printf("affect rows:%d\n", count)
}

func main()  {
	err := initDb()
	if err != nil {
		fmt.Printf("init db faild, err:%v\n", err)
		return
	}

	//testQuery()
	//testQueryMulti()
	testUpdateMulti()
}
