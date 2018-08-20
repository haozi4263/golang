package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Sex int
	Age int
}

func main()  {
	var s Student
	v := reflect.ValueOf(&s) //传入结构体地址

	v.Elem().Field(0).SetString("jude")//设置第0个字段值,设置结构体的值
	v.Elem().FieldByName("Sex").SetInt(22)
	v.Elem().FieldByName("Age").SetInt(19)

	fmt.Printf("s: %#v\n", s)

}
