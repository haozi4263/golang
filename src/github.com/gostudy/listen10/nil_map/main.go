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
	}
}