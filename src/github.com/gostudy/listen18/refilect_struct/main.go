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
	v := reflect.ValueOf(s)
	t := v.Type()  //等于t := reflect.TypeOf(s)
	//t := reflect.TypeOf(s)
	kind := t.Kind()
	switch (kind) {
	case reflect.Int64:
		fmt.Printf("s is int64\n")
	case reflect.Float64:
		fmt.Printf("s is float64\n")
	case reflect.Struct:
		fmt.Printf("s is struct\n")
		fmt.Printf("fied num of s is %d\n",v.NumField() )  //获取字段数量
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fmt.Printf("name:%s type:%v value:%v\n", t.Field(i).Name,
				field.Type(),field.Interface())   //获取结构体字段名字，字段类型，字段值
		}
	default:
		fmt.Printf("default\n")
	}


}
