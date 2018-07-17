package main

import "fmt"

func sayhelllo()  {   		//无参数和返回值函数
	fmt.Printf("hello word\n")
}

func add(a , b ,c int) int  {    //参数a,参数b,多个参数类型一样，前面类型可省略
	sum := a + b +c
	return sum
}

func calc(a,b int) (int,int)  {   //返回多个返回值需要用（）
	sum := a + b
	sub := a - b
	return sum,sub
}



func calc1(a,b int) (sum int, sub int)  {   //返回值命名，返回值使用变量命名
	sum = a + b  //不需要使用:=
	sub = a - b
	return
}

func cacl_v1(b ...int) int {  //接受0个或多个参数，是一个数组,可变参数*********
	sum := 0
	for i := 0;i < len(b);i++ {
		sum = sum + b[i]
	}
	return sum
}




func cacl_v2(a int,b ...int) int {  //a参数必须要传
	sum := a
	for i := 0;i < len(b);i++ {
		sum = sum + b[i]
	}
	return sum
}











func main()  {
	sayhelllo()

	s := add(100,200,400)  //变量s接受add函数返回值,接受参数
	fmt.Println(s)

	//sum,sub := calc(200,100)  //2个变量接受返回值
	//fmt.Println(sum,sub)

	//sum,sub := calc1(500,200)  //2个变量接受返回值
	//fmt.Println(sum,sub)

	//sum,sub := calc1(500,200)  //2个变量接受返回值
	//fmt.Println(sum,sub)

	//sum,_ := calc1(500,200)  //_忽略某个或多个返回值
	//fmt.Println(sum)
	//
	//sum,sub := calc1(500,200)  //2个变量接受返回值
	//fmt.Println(sum,sub)



	sum1 := cacl_v1(22,22,33,44,11)
	fmt.Printf("sum1=%d\n",sum1)

	sum2 := cacl_v2(10,33)
	fmt.Printf("sum2=%d\n",sum2)
}