package main

import "fmt"

func modify(a map[string]int)  {
	a["modify001"] = 22222
}


func main()  {
	var a map[string]int
	fmt.Printf("a:%v\n",a)

	if a == nil {
		a = make(map[string]int,16)
		fmt.Printf("a=%v\n",a)
		a["stu01"] = 1000
		a["stu02"] = 2000
		a["stu03"] = 3000
		//delete(a,"stu02")
		fmt.Printf("a=%v\n",a)

		b := a
		b["stu03"] = 2000
		fmt.Printf("a=%v\n",a)  //a=2000.a b指向同一个地址,赋值修改
		modify(a)
		fmt.Printf("modify a=%v\n",a)  //函数修改map
	}
}