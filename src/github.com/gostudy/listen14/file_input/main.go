package main

import (
	"fmt"
	"os"
)

func main()  {
	var a int
	var b string
	var c float32
	fmt.Fscanf(os.Stdin,"%d%s%f" ,&a, &b, &c)  //使用stdin当作文件输入
	//fmt.Fscan(os.Stdin, ,&a, &b, &c)
	//fmt.Scanf("%d%s%f", ,&a, &b, &c)

	//fmt.Println(a,b,c)   //输入 11 ddd 22.22
	fmt.Fprintln(os.Stdout, "stdout", a ,b, c )  //将stdoun当作文件输出到终端

}
