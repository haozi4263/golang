package main

import (
	"fmt"
	"math/rand"

)

func testInterFace()  {  //空的interface可以存储任何类型的值
	var a interface{}
	var b int =100
	var c float32 = 1.2
	var d string = "hello"

	a = b
	fmt.Printf("a=%v\n",a)
	a = c
	fmt.Printf("a=%v\n",a)
	a = d
	fmt.Printf("a=%v\n",a)
}

func studentStore()  {
	var stuMap map[int]map[string]interface{}    //定义map，key为id使用int，value为map的多维度 map对应value中的value为空的interface类型
	stuMap = make(map[int]map[string]interface{},16)   //初始化外层map
	//插入学生id=1，姓名=stu01,分数=99.22，年龄=18
	var id = 1
	var name = "stu01"
	var score = 99.22
	var age = 22

	value,ok := stuMap[id]
	if !ok {
		value = make(map[string]interface{},8)//对value内层的map初始化
 	}
 	value["name"] = name
 	value["id"] = id
 	value["score"] = score
 	value["age"] = age
	stuMap[id] = value

	fmt.Printf("stuMap:%#v",stuMap)
	for i := 0; i < 10;i++ {
		value,ok := stuMap[i]
		if !ok {
			value = make(map[string]interface{},8)
		}
		value["name"] = fmt.Sprintf("stu%d",i)
		value["id"] = i
		value["score"] = rand.Float32() * 100.0
		value["age"] = rand.Intn(100)
		stuMap[i] = value
	}
	fmt.Println()
	for k,v := range stuMap {
		fmt.Printf("id=%d stu info=%#v\n",k,v)
	}
}





func main()  {
	//testInterFace()

	studentStore()
}