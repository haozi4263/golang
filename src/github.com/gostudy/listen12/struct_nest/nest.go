package main

import "fmt"


type Address struct {
	Province string
	City string
}

type User struct {
	Username string
	Sex string
	address Address  //嵌套结构体
	//Address *Address  指针类型
}

func main()  {  //初始化User
	user :=&User {
		Username:"jude",
		Sex:"man",
		address: Address{  //嵌套结构体初始化
			//address: &Address{  //指针类型结构体积初始化
			Province:"beijing",
			City:"beijing",
		},
	}
	fmt.Printf("user=%#v", user)
}