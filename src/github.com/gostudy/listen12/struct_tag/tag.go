package main

import (
	"fmt"
	"encoding/json"  //json序列化
)

type User struct {
	UserName string `json:"user"`
	Sex		 string `json:"six"`
	Gg		 float32  //无tag信息会已字段名字做为key传给json序列化
}

func main()  {
	user := &User {
		UserName:"jude",
		Sex:"man",
		Gg:18.8,
	}
	data,_ := json.Marshal(user)   //序列化json字符串，kv格式(第一个表示数据，第二个错误忽略)
	fmt.Printf("json str:%s\n",string(data))  //强制将json转str,json将结构体中值替换成tag中值

}
