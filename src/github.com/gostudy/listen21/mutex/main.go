package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup
var mutex sync.Mutex   //定义锁，互是锁

func add()  {
	for i := 0; i < 5000; i++ {
		mutex.Lock()  //加锁
		x = x + 1
		mutex.Unlock()//释放锁
	}
	wg.Done()
}

func main()  {
	//add()
	//add()
	wg.Add(2)
	go add()
	go add()

	wg.Wait()
	fmt.Println("x:", x)
}
