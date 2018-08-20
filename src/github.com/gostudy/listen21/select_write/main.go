package main

import (
	"fmt"
	//"sync"
	"time"

)

func server1(ch chan string)  {
	time.Sleep(time.Second*6)
	ch <- "response from server1"
}

func write(ch chan string)  {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write success")
		default:
			fmt.Println("channel is full")  //判断chnnel是否满
		}
		time.Sleep(time.Millisecond*500)
	}
}


func main()  {
	select {} //阻塞分支
	output1 := make(chan string, 10)
	go write(output1)
	for s := range output1 {
		fmt.Println("recv:", s)
		time.Sleep(time.Second)
	}

}