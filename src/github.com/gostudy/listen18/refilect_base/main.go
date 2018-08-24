package main

import (
	"fmt"
	"reflect"
)

func reflect_type(a interface{})  {
	t := reflect.TypeOf(a)  //获取x变量类型信息
	fmt.Printf("type of a is:%v\n", t)

	k := t.Kind()
	switch (k) {
	case reflect.Int64:
		fmt.Printf("a is int64\n")
	case reflect.String:
		fmt.Printf("a is string_pkg\n")
	case reflect.Float32:
		fmt.Printf("a is float32\n")
	}

}

func reflect_value(a interface{})  {
	v := reflect.ValueOf(a)
	k := v.Kind()//等于t :=reflect.TypeOf(a)
	switch (k) {
	case reflect.Int64:
		fmt.Printf("a is int64, store value is:%d",v.Int())
	case reflect.String:
		fmt.Printf("a is string_pkg, store value is:%s",v.String())
	case reflect.Float32:
		fmt.Printf("a is float32, store value is:%f",v.Float())
	}
}



func reflect_set_value(a interface{})  {
	v := reflect.ValueOf(a)
	k := v.Kind()//等于t :=reflect.TypeOf(a)
	switch (k) {
	case reflect.Int64:
		v.SetInt(100)
		fmt.Printf("a is int64, store value is:%d",v.Int())
	case reflect.Ptr:  //接受指针类型
		v.Elem().SetInt(100000)    //获取指针指向的变量，完成赋值操作
		//fmt.Printf("a is string_pkg, store value is:%s",v.String())
	case reflect.Float64:
		v.SetFloat(6.6)
		fmt.Printf("a is float32, store value is:%f",v.Float())
	}
}

func main()  {
	//var x float64 = 3.4
	var x int64 = 111
	//reflect_type(x)
	//reflect_value(x)
	reflect_set_value(&x)  //需要传地址进去
	fmt.Printf("x value is %v\n", x)
}


/*
var *p int = new(int)
*p = 100
 */