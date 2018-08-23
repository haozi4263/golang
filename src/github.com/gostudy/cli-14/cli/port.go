package main

import (
	"fmt"
	"flag"
)

func main()  {
	var port int
	flag.IntVar(&port, "p", 8000, "-h帮助信息 默认：8000")
	flag.Parse()  //flag包对设置的变量赋值
	fmt.Printf("port=%d\n", port)
	//fmt.Printf("other args: %+v\n", flag.Args())  //很多CLI程序同时包含有标识和没有标识的参数。flag.Args() 将会直接返回哪些没有标识的参数。
	fmt.Println(flag.Args())
}
