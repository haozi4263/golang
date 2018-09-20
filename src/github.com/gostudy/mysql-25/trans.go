package main

import (
	"database/sql"  //sql接口
	_"github.com/go-sql-driver/mysql"  //导入mysql驱动,执行init函数，进行注册驱动
	"fmt"

)

var DB *sql.DB  //从database/sql中导入

func initDb() error  {
	var err error
	dsn := "golang:111111@tcp(192.168.10.105:3306)/golang"
	DB, err = sql.Open("mysql", dsn)  //open函数调用mysql-driver调用map
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)  //最大连接
	DB.SetMaxIdleConns(16)  //空闲连接
	return  nil
	//defer db.Close()
}


func testTrans()  {
	conn, err := DB.Begin()
	if err != nil {
		if conn != nil {
			conn.Rollback()
		}
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}
	sqlstr := "update user set age = 111112 where id = ?"
	_, err = conn.Exec(sqlstr,1)  //将id=1年龄修改22
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec sql:%s faile , err:%v\n", sqlstr, err)
		return
	}

	sqlstr = "update user set age = 222222; where id = ?"
	_, err = conn.Exec(sqlstr, 2)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec sql:%s faile , err:%v\n", sqlstr, err)
		return
	}
	err = conn.Commit()
	if err != nil {
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
}

func main()  {
	err := initDb()
	if err != nil {
		fmt.Printf("init db faild, err:%v\n", err)
		return
	}
	testTrans()
}