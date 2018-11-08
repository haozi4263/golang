package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

func (p Person) SetInfoValue()  { //方法：SetInfoValue
	fmt.Println("SetInfoValue =", p)
}

func (p *Person) SetPointer()  { //方法：SetPointer
	fmt.Println("SetPointer =", p)
}

func main() {
	//假如，结构体变量是一个普通变量，它能够调用那些方法 ，这些方法就是一个集合，简称方法集
	p := Person{"jude", 'f', 11} //普通变量传入
	p.SetPointer() //内部，先把p，转换成&p在调用，等价于(&p).SetPointer()

	p.SetInfoValue()
}