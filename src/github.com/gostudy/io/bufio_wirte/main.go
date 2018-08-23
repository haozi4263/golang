package main

import (
	"os"
	"fmt"
	"bufio"
)

func main()  {
	file, err := os.OpenFile("./test.txt",os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)  //trunc清空
	if err != nil {
		fmt.Println("open files is failed,err", err)
		return
	}

	defer file.Close()
	write := bufio.NewWriter(file)
	for i := 0; i < 10; i ++ {
		write.WriteString("hello golang\n" )  //写入buf缓冲区
	}
	write.Flush()  //写入磁盘

}
