package main

import "fmt"

//判断map中key是否存在





func main()  {
	var a map[string]int
	a = make(map[string]int,16)
	a["stu01"] = 111111111
	a["stu03"] = 333333333
	a["stu04"] = 444444444
	fmt.Println(a)

	var result int
	result = a ["stu05"]
	fmt.Printf("restult:%d\n",result)  //访问map中key不存在时，int类型返回0，strint返回空

	var ok bool
	var key string = "stu01"
	result,ok = a[key]  //2个key第一个key result对应值，第二个key ok是否存在map中，返回true/flase
	if ok == false {    //判断key是否存在
		fmt.Printf("key %s is not exist\n",key)
	} else {
		fmt.Printf("key %s is %d\n",key,result)
	}

}
