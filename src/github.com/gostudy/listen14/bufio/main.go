package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	var str string
	//fmt.Scanf("%s",&str)  //默认已空格作为分隔符，无法读取一行数据
	//fmt.Println("read from fmt:",str)


	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')  //指定分隔符\n, byt字节,第一个参数str，第二个参数错误
	fmt.Println("read from bufio:", str)
}