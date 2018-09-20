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

//接收者Person为普通变量，非指针，值语义，本质是值的一份拷贝,无法修改值
func (p Person) SetInfoValue(n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetInfoValue &p = %p\n", &p)
}

//接收者Person为指针变量，引用传递
func (p * Person) SetInfo(n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("SetInfoValue p = ", p)
}
func main()  {
	s1 := Person{"jude", 'm', 22}
	//var s1 Person
	fmt.Printf("&s1 = %p\n", &s1)

	//值语义
	s1.SetInfoValue("mike", 'm', 11)
	fmt.Println("s1 = ", s1)//无法修改

	//引用语义
	(&s1).SetInfo("youyou", 'f', 11)
	fmt.Println("s1 = ", s1)//无法修改
}
