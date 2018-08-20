package main

import (
	"fmt"
	"reflect"
)


type Student struct {
	Name string `json:"jjjjjjjjjj" db:"redis"`
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
	v := reflect.ValueOf(&s)//传入指针
	t := v.Type()

	field0 := t.Elem().Field(0)
	fmt.Printf("tag json=%s\n", field0.Tag.Get("json"))  //获取结构体tag信息
	fmt.Printf("tag db=%s\n", field0.Tag.Get("db"))


}
