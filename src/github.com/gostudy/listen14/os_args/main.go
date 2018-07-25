package main

import (
	"os"
	"fmt"
)

func main()  {
	fmt.Println("args[0]=", os.Args[0])  //将切片参数第0个参数打印出来，默认为程序名称
	if len(os.Args) > 1 {
		for index, v := range os.Args {
			if index == 0 {
				continue
			}
			fmt.Printf("args[%d]=%v\n", index, v)
		}
	}
}
