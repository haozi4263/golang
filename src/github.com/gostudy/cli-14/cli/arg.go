package main

import (
	"fmt"
	"os"
)

func main()  {
	cmd := os.Args[0]
	fmt.Printf("program name: %s\n", cmd)  //打印当前程序的绝对路径


	for i,a := range os.Args[1:] {
		fmt.Printf("argument %d is %s\n", i+1, a)
	}
}
