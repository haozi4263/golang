package main

import "fmt"
type Person struct {
	name string
	sex byte
	age int
}

type Student struct {
	Person
	id int
	add string
	name string  //同名字段 
}

func main()  {
	var s Student
	s.name = "jude"
	//操作的是student的name还是操作person的name
	//默认规则就近原则，如果在本作用域找到此成员，就操作此成员
	//如果没有找到，找继承字段成员
	s.sex = 'm'
	s.age = 11
	s.add = "wuhan"
	s.Person.name = "lj"  //显示调用,指定结构图调用
	fmt.Printf("s = %+v\n", s)
}