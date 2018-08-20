package iniconfig

import (
	"strings"
	"fmt"
	"reflect"
	"errors"
)

func Marshal(data interface{}) (result []byte, err error)  {  //把结构体序列化ini配置文件
	return
}

func UnMarshal(data []byte, result interface{}) (err error) {
	//fmt.Println(string(data))  //打印原始配置文件

	lineArr := strings.Split(string(data), "\n")  //每行组成的字符数组
	for _, v := range lineArr {
		fmt.Printf("%s\n", v)
	}   //打印解析的配置

	typeInfo := reflect.TypeOf(result)  //获取用户传入的类型信息
	if typeInfo.Kind() != reflect.Ptr {  //判断用户传入是否为指针
		err = errors.New("please pass address")
		return
	}
	typeInfo2 := typeInfo.Elem()
	if typeInfo2.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
	}

	var lastSectionFieldName string
	for index, line := range lineArr {
		line = strings.TrimSpace(line)   //去掉字符串前后空格
		if len(line) == 0 {  //判断空行
			continue
		}

		if line[0] == ';' || line[0] == '#' {  //判断是否为注释符号
			continue
		}

		if line[0] == '['   && len(line) <= 2  {  //判断第一个字符为[括号并且字符长度不小于2，非法
			err = fmt.Errorf("syntax error, invaild section:%s, lineNo:%d", line, index+1)
			return
		}

		if line[0] == '[' && line[len(line)-1] != ']' {
			err = fmt.Errorf("syntax error, invaild section:%s, lineNo:%d", line, index+1)
			return
		}

		sectionName := strings.TrimSpace(line[1:len(line)-1 ] ) //获取中括号内字符,并去掉前后空格
		if len(sectionName) == 0 {
			err = fmt.Errorf("syntax error, invaild section:%s, lineNo:%d", line, index+1)
			return
		}

		for i := 0; i < typeInfo2.NumField(); i++ {  //获取结构体内字段个数
			field := typeInfo2.Field(i)
			tagValue := field.Tag.Get("ini")   //获取结构体的value
			if tagValue == sectionName {
				lastSectionFieldName = field.Name
				fmt.Println("field name:", lastSectionFieldName)
				break
			}
		}
	}
	return
}