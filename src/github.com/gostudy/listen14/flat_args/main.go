package main

import (
	"flag"
	"fmt"
)

var recusive bool
var test string
var level int

func init()  {

	flag.BoolVar(&recusive, "r", false, " 使用-r选项")
	flag.StringVar(&test, "t", "hello golang", "使用-t选项")
	flag.IntVar(&level, "h", 888, "使用-h选项")

	flag.Parse()
}

func main()  {
	fmt.Printf("recusive:%v\n", recusive)
	fmt.Printf("test:%v\n", test)
	fmt.Printf("level:%v\n",level)
}