package main

import "fmt"

type Interger int   //这是Interger为int的别名

func (i Interger) Print()  {  //为int类型增加方法
	fmt.Println(i)
}


func main()  {
	var a Interger
	a = 1000
	a.Print()

	var b int = 200
	a = Interger(b)  //必须将b强制转换成a的类型
	a.Print()
}
