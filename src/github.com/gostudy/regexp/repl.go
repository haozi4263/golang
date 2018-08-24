package main

import (
	"fmt"
	"regexp"
	"strings"
)


//查找a.e替换结果中替换a为zzzz
func main()  {
	str := `abcd a111d aed`

	reg := regexp.MustCompile(`a.d`)  //查找a开头d结尾，中间任意单个字符
	if reg == nil {
		fmt.Println("reg err= ", reg)
	}

	reg_sult := reg.FindAllString(str, -1)  //提取关键信息
	fmt.Println("reg_sult = ",reg_sult)  // aed

	var data string
	for _, data = range  reg_sult {

		fmt.Println("data= ",data)
		repl := strings.Replace(data, "a","zzzz", -1) //替换data中a为zzzz
		fmt.Println("替换 = ", repl)


	}


}

