package main

import (
	"fmt"
	"strings"
)

func main()  {
	//"hellogo"中是否包含"hello",包含返回true,不包含返回false
	fmt.Println(strings.Contains("hellogo", "hello"))  //true

	//Joins组合
	s := []string{"abc", "hello", "jude", "golang"}  //定义切片
	buf := strings.Join(s,"/") //已@连接
	fmt.Println("buf = ", buf)

	//Index,查找子串的位置
	fmt.Println(strings.Index("abcdhello", "hello"))//4：查看hello在字符串所在位置
	fmt.Println(strings.Index("abcdhello", "go"))//不存在返回-1

	//复制
	data := strings.Repeat("golang", 3)
	fmt.Println("data=",data)

	//Split已指定的分隔符拆分
	split := "hello@  abc@go@jude   "
	s1 := strings.Split(split, "@")
	fmt.Println("s1=", s1)

	//Trim去掉两头字符
	s2 := strings.Trim(" ///  are u ok ?    ///  "," ")  //去掉2头空格
	s3 := strings.Trim(" ///  are u ok ?    ///  "," /")  //去掉2头/
	fmt.Println("s2=", s2)
	fmt.Println("s3=", s3)

	//去掉空格，把元素放入切片中
	s4 := strings.Fields("     are u ok?    ")
	//fmt.Println("s3=", s4)
	for i, data1 := range s4{
		fmt.Println(i , ",", data1)
	}
}

