package main

import "fmt"



func main()  {
	var a map[string]int
	fmt.Printf("a:%v\n",a)  //没初始化默认为空，不能直接赋值
	//a["abc"] = 100  无法直接赋值
	if a == nil {
		a = make(map[string]int,16)  //记录条数为16，扩容是先将之前map拷贝到新的map
		fmt.Printf("befer a:%v\n",a)
		a["stu01"] = 1000  //初始化key,value
		a["stu02"] = 2000
		a["stu03"] = 3000
		fmt.Printf("after a:%#v\n",a)  //%#v
		delete(a,"stu03")  //删除stu03
		fmt.Printf("delete a:%#v\n",a)  //%#v
		fmt.Printf("length is %v\n",len(a))  //取map长度

		for key,_ := range a {
			delete(a,key)   //删除所有
		}
		fmt.Printf("delete all  a:%#v\n",a)  //%#v
	}
}