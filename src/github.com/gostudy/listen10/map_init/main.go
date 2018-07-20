package main

import "fmt"


func main()  {
	var a map[string]int = map[string]int {  //声明map时初始化可直接赋值
		"stu01": 1000,
		"stu02": 2000,
		"stu03": 3000,
	}
	fmt.Println(a)
	a["stu01"] = 111111111
	a["stu03"] = 333333333
	a["stu04"] = 444444444
	fmt.Println(a)

	var key string = "stu04"
	fmt.Printf("the value of key[%s] is :%d\n",key,a[key])
}
