package main

import "fmt"

type Test struct {
	A int32  //int32占4字节
	B int32
	C int32
	D int
}

func main()  {  //测试struct内存地址是否为连续地址
	var t Test
	fmt.Printf("a addr:%p\n",&t.A)
	fmt.Printf("b addr:%p\n",&t.B)
	fmt.Printf("c addr:%p\n",&t.C)
	fmt.Printf("d addr:%p\n",&t.D)
}