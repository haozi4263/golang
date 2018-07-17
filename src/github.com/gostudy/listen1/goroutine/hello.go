package main

import "fmt"
import "time"
/*
多行注释
*/



//子线程
func calc() {
	for i :=0;i < 10; i++ {
		time.Sleep(1*time.Second)
		fmt.Println("run",i,"times")
	}
	fmt.Println("calc finish")
}


// 主线程
func main() {
	//go 设置并行
    go calc()
	fmt.Println("i extied")
	time.Sleep(time.Second * 11)
}
