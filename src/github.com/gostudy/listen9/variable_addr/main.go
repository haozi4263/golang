package main

import "fmt"

func main()  {
	var a int32
	a = 100  //通过变量来操作内存地址，占用4字节
	fmt.Printf("the var addr of a %p",&a)  //取a变量地址
}
