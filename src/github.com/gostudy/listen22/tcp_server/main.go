package main

import "fmt"
import "net"

func main()  {
	listen, err := net.Listen("tcp", "0.0.0.0:1111")
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed,err:%v\n", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn)  {

	defer conn.Close()
	for {
		var buf[128]byte  //定义数组
		n, err := conn.Read(buf[:])  //接受一个切片
		if err != nil {
			fmt.Printf("read from conn faild ,err:%v\n", err)
			break;
		}
		str := string(buf[:n])  //转字符串，0-n
		fmt.Printf("recv from client, data:%v\n", str)
	}
}
