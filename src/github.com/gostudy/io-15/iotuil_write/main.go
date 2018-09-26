package main

import (
	"io/ioutil"
	"fmt"
)

func main()  {
	str := "准备要写入的数据"
	err  := ioutil.WriteFile("./test.txt",[]byte(str), 0755)   //字节数组转byte
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}

}
