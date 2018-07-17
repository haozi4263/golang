package main

import "fmt"

func testFor1()  {    //变量值存在for循环内部
	var i int
	for i = 1; i <10;i++ {
		fmt.Printf("i=%d\n",i)
	}
	fmt.Printf("last i=%d\n",i)
}


func testFor2()  {    //变量值存在for循环内部
	var i int
	for i = 1; i <10;i++ {
		if i > 5 {
			break   //跳出循环
		}
		fmt.Printf("i=%d\n",i)
	}
	fmt.Printf("last i=%d\n",i)
}


func testFor3()  {    //变量值存在for循环内部
	var i int
	for i = 1; i <10;i++ {
		if i % 2 ==0 {
			continue   //跳出本次循环
		}
		fmt.Printf("i=%d\n",i)
	}
	fmt.Printf("last i=%d\n",i)
}

func testFor4()  {   //省略第一部分for表达式和第三部分for表达式
	i := 0;
	for ; i <= 10; {
		fmt.Printf("testF4 i=%d\n",i)
		i = i+2
	}
}


func testFor5()  {   //省略第一部分和第三部分
	i := 0;
	for  i <= 10 {  //省略2个;号
		fmt.Printf("testF5 i=%d\n",i)
		i = i+2
	}
}

func testMultSign()  {   //同时赋值多个变量
	var a int
	var b int
	var c string
	a,b,c = 10,20,"hello"
	fmt.Printf("a=%d b=%d c=%v",a,b,c)
}

func testFor6()  {
	for num,i := 10,1;i<=10 && num <= 19;i,num = i +1,num+1 {
		fmt.Printf("testFor6  %d * %d =%d\n",num,i,num*i)
	}
}

func testFor7()  {    			//无限循环
	for {
		fmt.Printf("死循环")
	}
}

func main()  {
	//testFor1()
	testFor2()
	testFor3()
	testFor4()
	testFor5()
	testFor6()
	testFor7()
	testMultSign()
}
