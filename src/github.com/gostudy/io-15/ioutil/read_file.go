package main

import (
	"io/ioutil"
	"fmt"
)

func main()  {
	content, err  := ioutil.ReadFile("./read_file.go")//读取所有内容
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}

	fmt.Println(string(content))  //打印二进制转string
}
