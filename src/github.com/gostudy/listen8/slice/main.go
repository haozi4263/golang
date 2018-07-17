package main

import (
	"fmt"

)

func testSlice4()  {
	a := [...]int{1,2,3,4,5,7,8,9,11}
	fmt.Sprintf("array a:%v type of a:%T\n",a,a)

	b := a[2:5]    //修改切片值数组a不变，切片是数组的引用
	fmt.Printf("slice b:%v type of b:%T\n",b,b)
	//b[0] = b[0] + 10    方法1：
	//b[1] = b[1] + 20
	//b[2] = b[2] + 30

	//for index,val := range b { 方法2
	//	fmt.Printf("修改前: b[%d]=%d\n",index,val)
	//}

	for index := range b {
		b[index] = b[index] + 10
	}
	fmt.Printf("after modify slice b,aaaay a:%v type of a:%T\n",a,a)

}

func testSlice5()  {
	a := [...]int{1,2,3}
	s1 := a[:]   //切片修改
	s2 := a[:]
	s1[0] = 100
	fmt.Printf("a=%v s2=%v\n",a,s2)
	s2[1] = 200
	fmt.Printf("a=%v s1=%v\n",a,s1)
}












func main()  {
	//testSlice4()
	testSlice5()
}