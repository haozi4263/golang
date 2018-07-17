package main

import (
	"fmt"
	"flag"
	"math/rand"   //随机算法包
	"time"
)

var (
	length int
	charset string
)

const (   //定义常量
	NumStr = "0123456789"
	CharStr = "abcdefjhijklmnopqrstupwxyzABCDEFGHIJKLMNOPQRSTUPWXYZ"
	SpecStr = "!@#$%^&*"
)

//func test1()  {
//	for i := 0 ; i < len(CharStr);i ++ {
//		if CharStr[i] != ' ' {
//			fmt.Printf("%c", CharStr[i])
//		}
//	}
//}   //生成26个字母



func parseArgs()  {
	flag.IntVar(&length,"l",16,"-l 生成密码的长度")       //&取变量地址，选项名字，默认值,帮助文档
	flag.StringVar(&charset,"t","num",
		`-t 指定密码生成的字符集，
				num:只能使用数字[0-9],
				char:只使用英文字母[a-zA-z],
				mix: 使用数字和字母，
				advance： 使用数字/字母以及特殊字符`)
	flag.Parse()
}

func generatePasswd() string {
	var passwd []byte = make([]byte,length,length)  //定义切片，长度为用户指定长度，容量为lengt
	var sourceStr string   //定义原始字符串
	if charset == "num" {
		sourceStr = NumStr
	} else if charset == "char" {
		sourceStr = CharStr
	} else if charset == "mix" {
		sourceStr = fmt.Sprintf("%s%s",NumStr,CharStr)
	} else if charset == "advance"{
		sourceStr = fmt.Sprintf("%s%s%s",NumStr,CharStr,SpecStr)
	} else {
		sourceStr = NumStr
	}
	fmt.Println("source",sourceStr)    //输出原始字符是那种类型
	for i :=0; i < length; i++ {
		index := rand.Intn(len(sourceStr))   //随机生成下标字符
		passwd[i] = sourceStr[index]
	}

	return string(passwd)   //切片强制转成字符
}


func main()  {
	rand.Seed(time.Now().UnixNano())	//初始化随机数种子,使用当前时间
	parseArgs()
	//fmt.Printf("length:%d charset:%s\n",length,charset)
	//test1()
	passwd := generatePasswd()
	fmt.Println(passwd)
}


//随机生成密码器，用法example3.exe -l 8 -t chat/mix/advance