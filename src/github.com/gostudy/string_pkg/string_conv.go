package main

import (
	"fmt"
	"strconv"

)

func main() {
	//转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	//第二个数为要追加的数，第三个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcgohello")

	fmt.Println("slice = ", string(slice))

	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println("str = ", str)

	//f:指打印格式，一小数方式， -1：指小数点位数， 64：已float64处理
	str = strconv.FormatFloat(3.12, 'f', -1, 64)
	fmt.Println("str = ", str)

	//整型转换字符串，常用
	str = strconv.Itoa(6666)
	fmt.Println("str = ", str)

	//字符串转其他类型

	var flag bool
	var err error
	flag, err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err =", err)
	}

	//把字符串转整型
	a, _ := strconv.Atoi("564")
	fmt.Println("a= ", a)

}
