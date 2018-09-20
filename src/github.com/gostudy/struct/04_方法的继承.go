package main

import (
	"fmt"
)


type Person struct {
	name string
	sex	 byte
	age  int
}

//Person类型，实现了一个方法
func (tmp *Person) PrintInfo()  {
	fmt.Printf("name = %s sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

//有个学生，继承了Person字段，成员和方法都继承了
type Student struct {
	Person //匿名字段
	id	int
	addr string
}

func main()  {
	s := Student{Person{"jude", 'f', 11}, 11, "wuhan"}
	s.PrintInfo()  //继承方法
}