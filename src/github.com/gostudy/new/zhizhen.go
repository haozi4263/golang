package main

import "fmt"

func swap(a, b int)  {
	a, b = b, a
	fmt.Printf("swap: a=%d, b=%d\n", a, b)
}

func swap1(p1, p2 *int)  {
	*p1, *p2 = *p2, *p1

}

//普通变量做函数的参数
func main()  {
	a, b := 10, 20
	//通过一个函数交换a和b的内存
	swap(a, b)  //变量本身传递，值传递（站在变量角度）
	fmt.Printf("main: a%d, b%d\n", a, b)

	swap1(&a, &b)  //地址传递
	fmt.Printf("main: a%d, b%d\n", a, b)

}
