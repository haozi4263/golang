package main

import (
	"fmt"
	"encoding/json"
)

type IT struct {
	Id		int			`json:"id"`   // json--json
	Name 	string		`json:"name"`
	KeCheng	[]string	`json:"kecheng"`
	IsOk	bool		`json:",string"`
	Age 	float64		`json:",string"`
}

func main()  {
	jude := IT{1, "jude", []string{"golang", "python", "java"},true,18}
	//json, err := json.Marshal(jude)//标准输出
	json, err := json.MarshalIndent(jude,"", "	")  //
	if err != nil {
		fmt.Printf("json marsh is faild, err:%v\n", err)
		return
	}
	fmt.Println("json ", string(json))
}

