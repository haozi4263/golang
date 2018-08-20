package main

import (
	"runtime"
	"fmt"
)

var i int

func calc()  {
	i++
}

func main()  {
	cpu := runtime.NumCPU()
	fmt.Println("cpu", cpu)  //获取cpu核

	runtime.GOMAXPROCS(1)  //设置1个核运行程序，16后默认使用所有cpu核

	for i :=0; i < 10; i++ {
		go calc()
	}
}