package main

import "fmt"


type User struct {
	Username string
	Sex	string
	int   //匿名字段
	string  //匿名字段,默认使用类型字段名字使用名字
}

func main()  {
	var user User
	user.Username = "jude"
	user.Sex = "man"
	user.int = 19
	user.string = "golang"

	fmt.Printf("user=%#v",user)
}