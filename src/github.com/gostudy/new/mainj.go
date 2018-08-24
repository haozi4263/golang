package main

import (
	"fmt"
)

func main()  {
	//a := 10 //整形变量a=10
	var  p *int //定义指针变量,动态分配内存空间，不需要释放，gc内部垃圾回收
	//p = &a
	// q := new(int)
	p = new(int ) //p必须是int,指向int类型,初始化内存地址
	*p = 666
	fmt.Println("*p=", *p)
}

