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

	for {
		v, ok := <-ch
		if ok == false {
			fmt.Println("chan is closed")
			break  //退出循环
		}
		fmt.Println("received", v,ok)
	}
}