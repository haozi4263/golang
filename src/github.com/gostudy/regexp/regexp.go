package main

import (
	"fmt"
	"regexp"
)

func main()  {
	buf := "abc azc a11c aac 9999 a9c tac adddc a3c"
	//1.解释规则,解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`a.c`)  //匹配已a开头中间任意单个字符已c结尾
	reg2 := regexp.MustCompile(`a\dc`)  //0-9任意单个数字
	if reg1 == nil { //解释失败，返回Nil
		fmt.Println("regexp err")
		return
	}
	if reg2 == nil { //解释失败，返回Nil
		fmt.Println("regexp err")
		return
	}
	//2.根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1) //-1提取所有
	fmt.Println("resulte1 = ", result1)

	result2 := reg2.FindAllStringSubmatch(buf, -1) //-1提取所有
	fmt.Println("resulte2 = ", result2)

}

