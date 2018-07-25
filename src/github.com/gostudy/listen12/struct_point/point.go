package main

import "fmt"

type User struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

func main() {
	var user *User
	//user.Age = 10
	fmt.Printf("user=%v\n", user)

	var user01 *User = &User{}  //&User{}分配对象，并返回指针
	/*
		(*user01).Age = 18
		(*user01).Username = "user01"
	*/

	user01.Age = 18  //编译器帮你按上面方法(*user01)转换
	user01.Username = "user01"
	fmt.Printf("user01=%#v\n", user01)

	var user02 *User = &User{  //声明定义时候完成初始化
		Username: "jude",
		Age:      18,
	}
	fmt.Printf("user02=%#v\n", user02)

	var user03 *User = new(User)  //new初始化，分配内存
	user03.Username = "lj"
	user03.Age = 18

	fmt.Printf("user03=%#v\n", user03)
}