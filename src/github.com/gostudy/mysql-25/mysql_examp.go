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

type User struct {
	Id int64 `db:"id"`   //根据数据库字段类型动态选择字段对应类型，数据库id为int
	Name sql.NullString `db:"name"`  //当name字段为空时需哟啊加上
	Age int `db:"age"`
}

func testQueryMultiRow()  {
	sqlstr := "SELECT id, name,age from user where id > ?"
	rows, err := DB.Query(sqlstr,0 )  //多行查询id>2
	defer func() {
		if rows != nil {
			rows.Close()  //关闭mysql连接，不然会mysql一致增长
		}
	}()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	for rows.Next() {  //如果next返回true,rows存在数据,当next返回false，退出程序
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("user:%#v\n", user)
	}
}


func testQueryDate()  {
	sqlstr := "SELECT id, name,age from user where id=?"
	row := DB.QueryRow(sqlstr,2 )  //单行查询id=2
	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		fmt.Printf("scan failed,err:%v\n", err)
		return
	}

	fmt.Printf("id:%d name:%s age:%d\n", user.Id, user.Name, user.Age)


}

func testInsectDate()  {
	sqlstr := "insert into user(name, age) VALUES (?, ?)"  //？号为mysql占位符
	resule, err := DB.Exec(sqlstr, "tom", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	id, err := resule.LastInsertId()  //获取插入数据
	if err != nil {
		fmt.Printf("get last insert failed, err:%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)

}


func testUpdateDate()  {
	sqlstr := "UPDATE user set name=? WHERE id =?"  //？号为mysql占位符
	resule, err := DB.Exec(sqlstr, "jude", 4)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	//id, err := resule.LastInsertId()  //获取插入数据
	//if err != nil {
	//	fmt.Printf("get last insert failed, err:%v\n", err)
	//	return
	//}
	affected, err := resule.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows is failed, err:%v\n", err)
	}
	fmt.Printf("update db success, affected rows:%d\n", affected)

}



func testDeleteDate()  {
	sqlstr := "DELETE from user WHERE  id=?"  //？号为mysql占位符
	resule, err := DB.Exec(sqlstr,4)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	//id, err := resule.LastInsertId()  //获取插入数据
	//if err != nil {
	//	fmt.Printf("get last insert failed, err:%v\n", err)
	//	return
	//}
	affected, err := resule.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows is failed, err:%v\n", err)
	}
	fmt.Printf("delete db success, affected rows:%d\n", affected)

}

func testPrepareDate()  {  //使用预处理方式
	sqlstr := "SELECT id, name,age from user where id > ?"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	defer  func() {
		if stmt != nil {
			stmt.Close()
		}
	}()  //关闭

	rows, err := stmt.Query(0)
	defer func() {
		if rows != nil {
			rows.Close()  //关闭mysql连接，不然会mysql一致增长
		}
	}()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	for rows.Next() {  //如果next返回true,rows存在数据,当next返回false，退出程序
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("user:%#v\n", user)
	}
}


func testPrepareInsectDate()  {
	sqlstr := "insert into user(name, age) VALUES (?, ?)"  //？号为mysql占位符
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	result ,err := stmt.Exec("jude", 4)
	id, err := result.LastInsertId()  //获取插入数据
	if err != nil {
		fmt.Printf("get last insert failed, err:%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)

}




func main()  {
	err := initDb()
	if err != nil {
		fmt.Printf("init db faild, err:%v\n", err)
		return
	}

	//testQueryDate()
	//testQueryMultiRow()
	//testInsectDate()
	//testUpdateDate()
	//testDeleteDate()
	//testPrepareDate()
	testPrepareInsectDate()
}


