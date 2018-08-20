package main

import (
	"fmt"
	"net"
	"io"
)

func main()  {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}

	data := "GET / HTTP/1.1\r\n"
	data += "HOST: www.baidu.com\r\n"
	data += "connnect: close\r\n"
	data += "\r\n\r\n"  //结束

	_, err = io.WriteString(conn, data)
	if err != nil {
		fmt.Printf("write is failed, err:%v\n", err)
		return
	}
	var buf[1024]byte  //定义数组
	for {
		n, err := conn.Read(buf[:]) //传入切片，读取数据
		if err != nil || n == 0 {
			break
			}
		fmt.Println(string(buf[:n])) //将数据打印转string
	}
}
