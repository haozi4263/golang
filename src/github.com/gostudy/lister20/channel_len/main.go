package main

import "fmt"

func main()  {
	ch := make(chan string, 3)
	ch <- "hello"
	ch <- "golang"
	fmt.Println("容量等于:", cap(ch))
	fmt.Println("长度等于:", len(ch))  //长度等于2
	fmt.Println("读取channel值并扔掉", <-ch )
	fmt.Println("channel长度等于", len(ch))//取出一个元素后，长度等于1

}
