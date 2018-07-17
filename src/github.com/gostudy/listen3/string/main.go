package main

import (
	"fmt"
	//"syscall"
)

func testString() {
	var str = "hello"  //string类型位byt类型数组,默认字符串是只读的
	fmt.Printf("str[0]=%c\n len(str)=%d\n",str[0],len(str))

	for index,val := range str {
		fmt.Printf("str[%d]=%c\n",index,val)
	}

	var byteSlice []byte
	byteSlice = []byte(str)  //将str类型转成byte切片类型
	byteSlice[0] = 'o'  //修改第一个字符位o
	str = string(byteSlice)

	fmt.Println("after modify:",str)

	fmt.Printf("len(str)=%d\n",len(str))

	str = "你好，golang"  //一个中文字符占3个字节
	fmt.Printf("长度:%d\n",len(str))

	var b rune = '中'  //rune设置utf-8,占一个字节
	fmt.Printf("b=%c\n",b)

	var runeSlice []rune
	runeSlice = []rune(str)
	fmt.Printf("str 长度：%d\n len(str)=%d\n",len(runeSlice),len(str))
}



//hello逆序v1
func testReverseStringV1()  {
	var str = "hello"
	var bytes []byte = []byte(str)

	for i := 0; i < len(str)/2;i++ {
		tmp := bytes[len(str)-i-1]
		bytes[len(str)-i-1] = bytes[i]
		bytes[i] = tmp
	}

	str = string(bytes)
	fmt.Printf(str)
}


func testReverseStringV2() {
	var str = "hello中文" 			 //当成字符交换
	var r []rune = []rune(str)

	for i := 0; i < len(r)/2;i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp	
	}

	str = string(r)
	fmt.Printf(str)
}


//回文，倒着和顺着念是一样的

func testHuiWen()  {
	var str = "上海自来水来自海上" 			 //当成字符交换
	var r []rune = []rune(str)

	for i := 0; i < len(r)/2;i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}

	str2 := string(r)
	if str2 == str {
		fmt.Println(str,"is huiwen")
	} else {
		fmt.Println(str,"is not huiwen")

	}
}

func main()  {
	//testString()
	//testReverseStringV1()
	//testReverseStringV2()
	testHuiWen()
}