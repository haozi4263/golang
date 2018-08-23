package main

import "fmt"
import "os"

func main()  {
	file, err := os.Open("read_file.go")  //打开一个文件
	if err != nil {
		fmt.Printf("error:%v\n", err)
	}
	data := make([]byte, 1024)  //文件的信息可以读取进一个[]byte切片。Read和Write方法从切片参数获取其内的字节数。
	count, err := file.Read(data)  //读进切片
	if err != nil {
		fmt.Printf("read error:%v\n",err)
	}
	fmt.Printf("read %d bytes: %v\n", count, string(data[:count]))  //读取二进制后转string
}