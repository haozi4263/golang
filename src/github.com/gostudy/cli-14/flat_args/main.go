package main

import (
	"flag"
	"fmt"
)

var recusive bool
var test string
var port int

func init()  {
	//
	flag.BoolVar(&recusive, "r", false, " 使用-r选项")
	flag.StringVar(&test, "t", "hello golang", "使用-t选项")
	flag.IntVar(&port, "p", 8888, "使用-p选项")

	flag.Parse()
}

func main()  {
	fmt.Printf("recusive:%v\n", recusive)
	fmt.Printf("test:%v\n", test)
	fmt.Printf("port:%v\n",port)

	//输入方法：.\flat_args.exe -p 2222  -t "nihao" -r ture
}