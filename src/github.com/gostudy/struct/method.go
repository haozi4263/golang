package main

import (
	"fmt"
)


//只要接收者类型不一样，这个方法就算是同名，也是不通方法，不会报错
type Person struct {
	name string
	sex	 byte
	age  int
}

// (tmp Person)带有接受者的函数叫方法
func (tmp Person) PirntInof()  {
	fmt.Println("tmp = ", tmp)
}

//通过一个函数，给成员赋值
func (p * Person) SetInfo(n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
}

func main()  {
	//定义一个结构体变量
	var p Person
	(&p).SetInfo("jude", 'f', 11)
	p.PirntInof()
}
