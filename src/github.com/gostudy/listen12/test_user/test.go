package main

import (
	"github.com/gostudy/listen12/user"  //引用自定义包
	"fmt"
)

func main() {
	//var u user.User //包名.结构体名---使用自定义静态包
	//u.Age = 18
	//u.Sex = "男"
	//fmt.Printf("user:=%v Sex:=%v\n", u.Age,u.Sex)

	u := user.NewUser("user01","男",11,"xxx.jpg")   //构照函数,返回u就是指向user类型的一个指针
	fmt.Printf("user=%#v\n",u)

}

