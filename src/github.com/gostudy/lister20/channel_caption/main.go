package main

import "fmt"

func main()  {
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "golang"
	//ch <- "没有空间了,无法插入了,报死锁deadlock"
	c1 := <- ch // c1=hello
	c2 := <- ch //c2=golang

	fmt.Println(c1, c2)
}

