package main

import "fmt"

func producer(chn1 chan  int)  {
	for i := 0; i < 10; i++ {
		chn1 <- i  //入队
	}
	close(chn1)
}

func main()  {
	ch := make(chan int)  //不带缓冲区的管道
	go producer(ch)
	for v:= range ch {  //推荐写法，自动判断管道关闭
		fmt.Println("receive",v)
	}
}