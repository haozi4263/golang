package main

import (
	"fmt"
	"time"
)

func produce(c chan  int)  {
	c <- 10000
}

func consume(c chan  int)  {
	data := <- c
	fmt.Printf("data:%v", data)
}


func main()  {
	var c chan int  //定义整数的管道，定义完后不能直接操所默认为nil,需要初始化
	fmt.Printf("c=%v\n", c)

	c = make(chan int)  //不带缓冲区,必须需要有消费者去取队列才能写入队列和取出
	go produce(c)
	go consume(c)  //消费者
	time.Sleep(time.Second * 5)

}
