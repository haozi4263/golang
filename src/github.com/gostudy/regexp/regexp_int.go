package main
import (
	"fmt"
	"regexp"
)

func main()  {
	buf := "13.14 567 avsdd 11111.23 1. 8.55 1ddff3 6.66"

	//解析正则表达式,+ 匹配前一个字符1次或多次
	reg := regexp.MustCompile(`\d+\.\d+`)   //匹配开头结尾前一次或多次是数字，中间为.
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}
	//提取关键信息
	//result := reg.FindAllString(buf, -1)//放入一个切片
	result := reg.FindAllStringSubmatch(buf, -1)//分组打印
	fmt.Println("result = ", result)
}