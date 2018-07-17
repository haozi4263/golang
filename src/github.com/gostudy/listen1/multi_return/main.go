package main

import (
	"fmt"
)
//多返回值
func add(a int, b int) (int,int)  {
	return a + b,a - b
}

func main()  {
	//接受返回值
	sum, sub := add(2,5)
	fmt.Println(sum,sub)
}