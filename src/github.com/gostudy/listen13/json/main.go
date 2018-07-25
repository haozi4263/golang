package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id string
	Name string
	Sex string
}


type Class struct {
	Name string
	Count int
	Student []*Student  //切片存储学生，指针类型
}

var rawJson = `{"Name":"it","Count":100,"Student":[{"Id":"0","Name":"stu%d0","Sex":"man"},{"Id":"1","Name":"stu%d1","Sex":"man"},{"Id":"2","Name":"stu%d2","Sex":"man"},{"Id":"3","Name":"stu%d3","Sex":"man"},{"Id":"4","Name":"stu%d4","Sex":"man"},{"Id":"5","Name":"stu%d5","Sex":"man"},{"Id":"6","Name":"stu%d6","Sex":"man"},{"Id":"7","Name":"stu%d7","Sex":"man"},{"Id":"8","Name":"stu%d8","Sex":"man"},{"Id":"9","Name":"stu%d9","Sex":"man"}]}
`


func main()  {
	c := &Class{   //初始化班级
		Name: "it",
		Count: 100,
	}
	for i := 0; i < 10; i++ {   //生成10个学生
		stu := &Student{
			Name: fmt.Sprint("stu%d",i),
			Sex: "man",
			Id: fmt.Sprintf("%d",i),  //转字符串
		}
		c.Student = append(c.Student,stu)    //加入班级切片中
	}

	data, err := json.Marshal(c)     //将班级对象转换json, (字节数组，错误码)
	if err != nil {  //如果报错
		fmt.Println("json marshal falid")
		return
	}

	fmt.Printf("json:%s\n",string(data))  //将json字节码转换成str输出

	//json反序列化
	fmt.Println("反序列化")
	var c1 *Class = &Class{}
	err = json.Unmarshal([]byte(rawJson),c1)   //第一个参数字符传，第二个参数
	if err != nil {
		fmt.Println("unmarshal failed")
		return
	}
	fmt.Printf("c1:%#v\n",c1)
	for _, v := range c1.Student {
		fmt.Printf("c1:%#v\n",v)
	}

}

