package main

import "sync"
import (
	"fmt"
	"time"
)

func process(i int, wg *sync.WaitGroup)  {
	fmt.Println("started Goroution", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() //计数-1
}

func main()  {
	no := 3
	var wg sync.WaitGroup  //初始化结构体，默认计数等于0
	for i := 0; i < no ; i++ {
		wg.Add(1)   //计数+1
		go process(i, &wg)
	}
	wg.Wait()//计数=0返回主线程退出
	fmt.Println("All go routines finished executing")
}
