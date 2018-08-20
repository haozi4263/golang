package main

import "fmt"
import (
	"net"
	"bufio"
	"os"
	"strings"
)

func main()  {
	conn, err := net.Dial("tcp", "0.0.0.0:1111")  //拨号连接服务端
	if err != nil {
		fmt.Printf("dial failed ,err:%v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)  //传入终端文件句柄
	for {   //获取终端用户输入
		data ,err := reader.ReadString('\n')   //读一行祖字符串，已\n为分隔符
		if err != nil {
			fmt.Printf("read from console failed\n", err)
			break
		}
		data = strings.TrimSpace(data)
		_, err = conn.Write([]byte(data))//将数据发送到服务端，接受字节数组
		if err != nil {
			fmt.Printf("wirte failed, err:%v\n", err)
			break
		}
	}
}





