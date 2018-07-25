package main

import (
	"fmt"

)

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

func modify_arr(a *[3]int)  {
	(*a)[0] = 100
}
func testPoint4()  {
	var b [3]int = [3]int{1,2,3}

	modify_arr(&b)
	fmt.Printf("b:%v\n",b)
}

func testPoint5()  {
	var a *int = new(int)
	*a = 100   //将a指向内存改成100
	fmt.Printf("*a=%d\n",*a)

	var b *[]int = new([]int)  //b指向空切片,空切片不能直接操作
	fmt.Printf("*b=%v",*b)
	//(*b)[0] = 100   空切片不能直接操作
	(*b) = make([]int,5,100)  //make初始化
	(*b)[0] = 100
	(*b)[1] = 200
	fmt.Printf("modify *b=%v",*b)
	}

func modifyInt(a int)  {
	a = 1000
}

func testPoint6()  {
	var b int = 10
	modifyInt(b)
	fmt.Printf("b=%d\n",b)   //b=10,a是b的拷贝
}

func testPoint7()  {
	var a int = 10
	var b *int = &a
	var c *int = b
	*c = 200
	fmt.Printf("*c=%d *b=%d *a=%d\n",*c,*b,a)  //a,b,c=200
}

func main()  {
	//testPoint1()
	//testPoint2()
	//testPoint3()
	//testPoint4()
	//testPoint5()
	//testPoint6()
	testPoint7()
}

//指针存储的试变量地址