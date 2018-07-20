package main

import (
	"fmt"
	"github.com/gostudy/listen11/calc"  //引用自定义包
	_"github.com/gostudy/listen11/calc"  //引用自定义包,不引用里面函数
)

func init()  {
	fmt.Printf("init func in main\n")   //先执行init函数，做程序初始化工作，可以写多个init函数
}
func init()  {
	fmt.Printf("init2 func in main\n")   //先执行init函数，做程序初始化工作，可以写多个init函数
}

func main()  {
	var sum = calc.Add(2,3)
	fmt.Printf("sum=%d\n",sum)
}