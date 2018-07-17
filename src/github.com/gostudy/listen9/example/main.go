package main

import "fmt"

func main()  {
	var a []string = make([]string,5,10)
	fmt.Println("a:",a)  //空字符串
	for i := 0; i < 10; i++ {
		a = append(a,fmt.Sprintf("%d",i))  //格式化输入到终端
	}
	fmt.Println("a",a)
}