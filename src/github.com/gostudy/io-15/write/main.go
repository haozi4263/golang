package main

import (
"os"
"fmt"
)

func main()  {
	file, err := os.OpenFile("./test.txt",os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)  //trunc清空,打开一个文件
	if err != nil {
		fmt.Println("open files is failed,err", err)
		return
	}

	defer file.Close()
	str := "hello golang2"
	//file.Write([]byte(str))//1.接受字节数组
	file.WriteString(str)  //2.直接写入字符串
}
