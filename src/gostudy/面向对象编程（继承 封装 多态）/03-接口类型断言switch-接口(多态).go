package main

import "fmt"

type Student struct {
	name string
	id int
}
func main()  {
	i := make([]interface{}, 3)//空接口切片
	i[0] = 1
	i[1] = "golang"
	i[2] = Student{"jude", 11}

	//类型查询，类型断言
	//第一个返回下标，第二个返回下标值， date分别是i[0], i[1], i[2]
	for index, date := range i {
		//第一个返回的是值，第二个返回判断结果的真假
		if value, ok := date.(int); ok == true {
			fmt.Printf("x[%d],类型为int,内容为%d\n", index, value)
		}else if value, ok := date.(string); ok == true  {
			fmt.Printf("x[%d],类型为int,内容为%s\n", index, value)
		} else if  value, ok := date.(Student); ok == true {
			fmt.Printf("x[%d],类型为Student,内容为: name: %s, id: %d\n", index, value.name, value.id)
		}
	}
}
