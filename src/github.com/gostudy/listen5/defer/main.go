package main

import (
	"fmt"

)

func testDefer1()  {
	defer fmt.Println("hello")   //defer不会立即执行，在函数返回时执行，适合做资源释放
	defer fmt.Println("hello v2")
	defer fmt.Println("hello v3")  //多个defer执行顺序先进后出，逆序
	fmt.Println("aaaa")
	fmt.Println("bbbbb")
}

func testDerfer2()  {
	for i :=0; i < 5; i++ {
		defer fmt.Println("i=%d",i)
	}

	fmt.Printf("running\n")
	fmt.Printf("return\n")
}

func testDefer3()  {
	var i int = 0
	defer fmt.Printf("defer i=%d\n",i)  //defer=0,一开始已经定义值0
	i = 1000
	fmt.Printf("i=%d\n",i)
}



func main()  {
	testDefer1()
	testDerfer2()
	testDefer3()


}