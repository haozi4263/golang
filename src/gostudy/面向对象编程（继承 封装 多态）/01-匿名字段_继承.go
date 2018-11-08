package main

import (
	"fmt"
)

type Person struct {
	name string
	sex byte 
	age int
}

type Student struct {
	Person
	id int
	add string
}

func main()  {
	var s1 Student = Student{  //顺序初始化
		Person{"jude",'m',18}, 1, "wuhan"}
	fmt.Println("s1:", s1)
	//自动推导类型
	s2 := Student{Person{"jude",'m',12}, 1, "wuhan"}
	fmt.Println("s2:", s2)
	fmt.Printf("s2 = %+v\n", s2) //+v显示更详细

	s3 := Student{id:2}  //指定字段初始化，为初始化自动默认空
	fmt.Printf( "s3 = %+v\n", s3)

	s4 := Student{Person:Person{name:"jude"},id:3} //匿名字段初始化
	fmt.Printf( "s4 = %+v\n", s4)
	}


