package main

import (
	"fmt"

	"time"
)


func hello(c chan bool)  {
	fmt.Println("input data to chan")
	time.Sleep(5 * time.Second)
	c <- true  //插入管道
}
func main()  {
	var exitChan chan bool
	exitChan = make(chan bool)
	go hello(exitChan)
	fmt.Println("向管道取出数据")
	<- exitChan


}
