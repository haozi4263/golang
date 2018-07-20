package main

import (
	"fmt"
	"time"
	"math/rand"
)

func sliceMap()  {
	rand.Seed(time.Now().UnixNano())
	var s []map[string]int    //定义空切片，元素为map
	s = make([]map[string]int,5,16)  //初始化切片5个元素，长度16 s为map
	for index,val := range s {     //遍历切片
		fmt.Printf("slice[%d]=%v\n",index,val)    //打印空切片等于map，需要先map初始化
	}
	fmt.Println()
	s[0] = make(map[string]int,16)  //初始化map
	s[0]["stu01"] = 100
	s[0]["stu02"] = 200
	s[0]["stu03"] = 300

	for index,val := range s {     //遍历切片
		fmt.Printf("slice[%d]=%v\n",index,val)    //打印空切片等于map，需要先map初始化
	}
}


func mapSlice()  {
	rand.Seed(time.Now().UnixNano())
	var s map[string][]int   //定义map的value值是一个空切片[]int
	s = make(map[string][]int,16)  //初始化map，定义容量为16,需要先对map内value值的切片进行初始化
	//s["100"][0]  = 100  会报错未初始化
	key := "stu01"
	value, ok := s[key]
		if !ok {
			s[key] = make([]int,0,16)   //初始化切片
			value = s[key]
		}
		value = append(value,100)
		value = append(value,200)
		value = append(value,300)
		s [key] = value
		fmt.Printf("map:%v\n",s)
}


func main()  {   //切片类型的map
	//sliceMap()
	mapSlice()
}