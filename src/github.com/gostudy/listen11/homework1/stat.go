package main

import (
	"fmt"
	"strings"
)

func statWordCount( str string) map[string]int {   //输入字符串，返回map的key为string,value为int
	var result map[string]int = make(map[string]int,128)   //定义map并初始化
	words := strings.Split(str," ") //将字符串已空白字符分割成返回一个切片
	for _,v := range words {
		count, ok := result[v]
		if !ok {
			result[v] = 1
		} else {
			result[v] = count + 1
		}
	}
	return result
}

func main()  {
	var str = "how do you do ? du you like me ?"
	result := statWordCount(str)
	fmt.Printf("result:%#v",result)

}