package main

import (
	"fmt"
)


type Person struct {
	name string
	sex	 byte
	age  int
}

func (p Person) SetInfoValue()  {
	fmt.Println("SetInfoValue")
}

func (p * Person) SetInfo()  {
	fmt.Println("SetInfo")
}
func main()  {
	//假如结构体变量是一个指针变量，它能够调用那些方法。这些方法就是一个集合，简称方法集
	p := &Person{"jude", 'm', 11}
	p.SetInfo()
	//内部做的转换，先把指针p,转换*p后在调用
	p.SetInfoValue()
}
