package main

import (
	"fmt"
)

func testa()  {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaa")
}

func testb(x int) {
	var a [10]int
	a[x] = 111 //当x为20时候，导致数据越界，产生一个panic,导致程序崩溃,内部实现
}


func main()  {
	testa()
	testb(20)

}


