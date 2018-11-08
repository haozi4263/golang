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
}

func main() {
	s1 := Student{Person{"jude", 'm', 12}, 1, "wuhan"}
	s1.name = "lj"
	s1.sex = 'f'
	s1.Person = Person{"golang", 'f', 11}  //匿名字段操作
	fmt.Println(s1.name, s1.id, s1.age, s1.add, s1.sex)
}