package main

import "fmt"

func sendData(sendch chan<- int)  {
	sendch <- 10  //入队
}

func readData(sendch <-chan  int)  {  //单向管道，无法写入管道
	//sendch <- 20
	data := <- sendch  //出队
	fmt.Println(data)
}


func main()  {
	chn1 := make(chan int)
	go sendData(chn1)
	readData(chn1)
}