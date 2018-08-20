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

func (s *Student) SetName(name string)  {
	s.Name = name
}

func (s *Student) Print()  {
	fmt.Printf("通过反射进行调用:%#v\n",s)
}

func main()  {
	var s Student
	s.SetName("xxx")
	v := reflect.ValueOf(&s)
	t := v.Type()

	fmt.Printf("struct student have %d methods\n", t.NumMethod())
	for i := 0; i < t.NumMethod(); i ++ {
		method := t.Method(i)
		fmt.Printf("struct %d method, name:%s type:%v\n", i, method.Name, method.Type)  //获取结构方法信息,按照字母排序
	}


	//通过refiect.value获取对应的方法并调用
	m1 := v.MethodByName("Print")
	var args []reflect.Value  //定义空切片,无参数方法
	m1.Call(args)

	m2 := v.MethodByName("SetName")
	var args2 []reflect.Value  //有参数方法
	name := "jude"
	nameVal := reflect.ValueOf(name)
	args2 = append(args2, nameVal)
	m2.Call(args2)

	m1.Call(args)
}