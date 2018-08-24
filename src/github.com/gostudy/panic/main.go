package main

import (
	"fmt"
)

func testa()  {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaa")
}

func testb() {
	//fmt.Println("bbbbbbbbbbbbbbbbbbbbbb")
	//显示调用panic函数，导致程序终端
	panic("this is a panic testb")
}

func testc()  {
	fmt.Println("ccccccccccccccccccccccccc")
}

func main()  {
	testa()
	testb()
	testc()
}

