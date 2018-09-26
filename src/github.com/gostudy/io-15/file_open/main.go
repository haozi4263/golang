package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//只读的方式打开
	file, err := os.Open("./file.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}

	defer file.Close()
	var content []byte
	var buf [128]byte  //1次读取128字节
	for {
		n, err := file.Read(buf[:])  //文件读取，接受1个切片
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read file:", err)
			return
		}
		content = append(content, buf[:n]...)
	}

	fmt.Println(string(content))
}