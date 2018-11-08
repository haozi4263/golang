package main

import "fmt"
type Person struct {
	name string
	sex byte
	age int
}
type mystr string //自定义类型，给一个类型改名
type Student struct {
	Person   //结构题匿名字段
	int   //基础类型的匿名字段
	mystr
}

func main()  {
	s := Student{Person{"jude",'m',18}, 666, "go"}
	s.int = 777
	s.mystr = "golang"
	s.Person = Person{"jude1", 'f', 11}
	fmt.Printf("s: %+v\n", s)
	fmt.Println(s.Person, s.int, s.sex)
}