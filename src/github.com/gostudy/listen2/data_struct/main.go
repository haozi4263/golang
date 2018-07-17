package main  //main可以执行程序，只能有一个main函数

import (
	"fmt"
	"strings"  //切割包
	"github.com/pingguoxueyuan/gostudy/listen2/access" //引用access包
)

func testBool() {
	var a bool   //默认false
	fmt.Println(a)
	a = true
	fmt.Println(a)

	a = !a
	fmt.Println(a)


	var b bool
	if a == true && b == true {   //于操作
		fmt.Println("right")
	} else {
		fmt.Println("&&not right")
	}


	if a == true || b == true {   //于操作
		fmt.Println("right")
	} else {
		fmt.Println("||not right")
	}



	fmt.Printf("%t\n %t\n",a,b)  //格式化输出布尔类型
}

func testInt() {
	var a int8  //-128,+127
	a = 18
	fmt.Println("a=",a)
	a = -12
	fmt.Println("a=",a)


	var b int
	b = 123232323232
	fmt.Println("b=",b)

	var c float32 = 3.342
	fmt.Println(c)
	fmt.Printf("a=%d a=%x c=%f\n",a,a,c)  //%a十进制，%x十六进制，%c,浮点数
}

func testStr()  {
	var a string = "hello word"
	fmt.Println(a)  //默认位空

	b := a
	fmt.Println(b)

	c := "c:\nhello"
	fmt.Println(c)

	c = `你好，golang。`   //c := "\nc:hello" //""可以处理逻辑 `` 原样输出
	fmt.Printf("a=%v b=%v c=%v\n",a,b,c)   //%v万能占位符，%s字符串占位符号

	clen := len(c)
	fmt.Printf("len of d=%d\n",clen)  //判断字符串长度

	c = c + c
	fmt.Printf("c=%s\n",c)  //字符串拼接方式1

	c = fmt.Sprintf("%s%s",c,c)
	fmt.Printf("c=%s\n",c)  //字符串内置函数拼接方式2

	ips := "192.168.1.1;192.168.1.2"
	ipArrsy := strings.Split(ips,";") //切割成数组
	fmt.Printf("first ip :%s\n",ipArrsy[0])
	fmt.Printf("first ip :%s\n",ipArrsy[1])

	result := strings.Contains(ips,"192.168.1.1") //判断字符串是否存在返回布尔值
	fmt.Println(result)

	str := "http://baidu.com"
	if strings.HasPrefix(str,"http") {  //判断前缀是否http
		fmt.Printf("str is ture http url\n")
	} else {
		fmt.Printf("str is not http url")
	}

	if strings.HasSuffix(str,"baidu.com") {  //判断结尾是否http
		fmt.Printf("str is baidu web\n")
	} else {
		fmt.Printf("str is not baidu web")
	}


	index := strings.Index(str,"baidu")  //判断baidu在字符串str第几个位置，从前往后
	fmt.Printf("baidu is index:%d\n",index)

	index = strings.LastIndex(str,"baidu")  //判断baidu在字符串str第几个位置，从后往前
	fmt.Printf("baidu last index:%d\n",index)


	var strArr []string = []string{"192.165.1.1","192.168.1.2","192.168.1.3"} //将数组已;分隔符拼接成字符串
	resultStr := strings.Join(strArr,";")
	fmt.Printf("result=%s\n",resultStr)
}

func testAccess() {
	//fmt.Printf("access.a=%d\n",access.a)  /小写变量或函数只能在内部包调用
	fmt.Printf("access.a=%d\n",access.A)
}

func testOperator() {  //算数操作符针对变量操作，逻辑操作符针对条件判断操作
	var a int = 2
	if a ==2 {
		fmt.Printf("is right\n")
	} else {
		fmt.Printf("is not right")
	}
}

func main() {
	//testBool
	//testInt()
	testStr()
	testOperator()
	testAccess()
}
