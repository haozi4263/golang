package main

import "fmt"

func testPoint1()  {
	var a int32
	a = 1000
	fmt.Printf("the addr of a:=%p,a:%d\n",&a,a)//输出a的指针值和变量值

	var b *int32
	fmt.Printf("the add of b:=%p,b%v\n",&b,b)  //指针没有初始化默认试空地址bil,需要先初始化，不然复制会空指针

	//*b = 100   没有初始化无法赋值
	if b == nil {
		fmt.Println("b is nil addr")
	}
	b = &a  //将a地址赋值给a完成初始化
	fmt.Printf("the last add of b:=%p,b%v\n",&b,b)
}

func testPoint2()  {
	var a int = 200
	var b *int = &a
	fmt.Printf("b指向的地址存储的值为：%d\n",*b)  //取b的内存存储的值

	*b = 1000
	fmt.Printf("b指向的地址存储的值为：%d\n",*b)  //修改
	fmt.Printf("a=%d\n",a)
}

func modify(a *int)  {  //p是a的拷贝
	*a = 100  //a修改b的值
}

func testPoint3()  {
	var b int = 10
	p := &b
	modify(p)
	fmt.Printf("b:%d\n",b) //b=100
}

func modify_arr(a *[]int)  {
	(*a)[0] = 100
}
func testPoint4()  {
	var b [3]int = [3]int{1,2,3}
}



func main()  {
	//testPoint1()
	//testPoint2()
	testPoint3()
}

//指针存储的试变量地址