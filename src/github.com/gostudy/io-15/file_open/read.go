package main

import (
	"fmt"
	"os"
	"io"
)

func main()  {
	//只读方式打开
	file, err := os.Open("./read.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	var content []byte  //定义切片，将数组内存追加到切片
	var buf [128]byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {  //判断文件读取完成,eof结束
			break
		}

		if err != nil { //判断出错
			fmt.Println("read failed:", err)
			return
		}
		content = append(content, buf[:n]...) //将数组buf内容追加到切片content中, ...展开


	}
	fmt.Println(string(content)) //打印切片内容并转换成字符串
}

