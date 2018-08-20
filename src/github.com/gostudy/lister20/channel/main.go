package main

import "fmt"

func main()  {
	var c chan int  //定义整数的管道，定义完后不能直接操所默认为nil,需要初始化
	fmt.Printf("c=%v\n", c)

	c = make(chan int, 1)  //初始化管道，存储1个元素，指向一块内存地址
	fmt.Printf("c=%v\n", c)

	c <- 100  //入队，将100插入c队列中
	data := <- c  //出队，将数据存储至data
	fmt.Printf("data:%v\n", data)

	//<-c   //出队，不保存数据

}
