package main

import "fmt"

func swap(a int,b int)  {  //a,b值交换，swap函数是main函数值拷贝，不会影响main函数值
	fmt.Printf("before a=%d b=%d\n",a,b)
	a,b = b,a
	fmt.Printf("after a=%d b=%d\n",a,b)
}


func swap_point(a *int,b *int)  {  //a,b值交换，swap函数是main函数值拷贝，不会影响main函数值
	fmt.Printf("before a=%d b=%d\n",*a,*b)
	*a,*b = *b,*a
	fmt.Printf("after a=%d b=%d\n",*a,*b)
}



func main()  {
	var a int = 10
	var b int = 20
	swap(a,b)
	fmt.Printf("in main a=%d b=%d\n",a,b)
	swap_point(&a,&b)
	fmt.Printf("in zhizhen a=%d b=%d\n",a,b)
}