package main

import "fmt"


a -->
func testMake1()  {
	var a []int
	a = make([]int,5,10)  //5切片的长度--下标，10切片的容量，占10个元素的数组
	a[0] = 10
	a[1] = 20
	fmt.Printf("a=%v\n",a)
}

func main()  {
	testMake1()
}