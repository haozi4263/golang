package main

import "fmt"


type Animal interface {   //定义接口,定义要给动物的接口，有3个方法
	Talk()
	Eat()
	Name()  string
}

type Dog struct {

}

func (d Dog) Talk() {
	fmt.Println("汪汪")
}

func (d Dog) Eat()  {
	fmt.Println("我在吃骨头")
}

func (d Dog) Name() string  {
	fmt.Println("im xiaoming")
	return "xiaoming"
}


type Pig struct {

}
func (d Pig) Talk() {
	fmt.Println("坑坑坑")
}

func (d Pig) Eat()  {
	fmt.Println("我在吃草")
}

func (d Pig) Name() string  {
	fmt.Println("我是猪八戒")
	return "猪八戒"
}

func testInterface1()  {
	var d Dog
	var a Animal
	a = d

	a.Talk()
	a.Eat()
	a.Name()

	var pig Pig   //定义实例
	a = pig
	a.Eat()
	a.Talk()
	a.Name()
}

func just(a Animal)  {
	//d, ok := a.(Dog)  方法1
	//p, ok := a.(Dog)  方法1
	switch v := a.(type) {  //方法2
	case Dog:
		fmt.Printf("v is dog:%v\n",v)
	case Pig:
		fmt.Printf("v is pig:%v\n",v)
	default:
		fmt.Printf("not support")
	}
}

func testjust()  {
	var d Dog
	just(d)
}



func main()  {
	testInterface1()
	//testjust()
}
