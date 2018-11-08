package main

import "fmt"
type Person struct {
	name string
	sex byte
	age int
}
type Student struct {
	*Person   //指针类型匿名字段
	id int
	add string
}

func main()  {
	s1 := Student{&Person{"jude", 'm', 22},18,"wuhan"}
	fmt.Println("s1 =", s1)  //打印出来为指针地址
	fmt.Println(s1.name, s1.sex, s1.add, s1.age, s1.id, s1.add) //方法1

	var s2 Student
	s2.Person = new(Person)  //分配空间,在赋值
	s2.name = "jude"
	s2.add = "wuhan"
	s2.age = 11
	s2.sex = 'f'
	s2.id = 1
	fmt.Println(s2.name, s2.sex, s2.add, s2.age, s2.id, s2.add)
}

