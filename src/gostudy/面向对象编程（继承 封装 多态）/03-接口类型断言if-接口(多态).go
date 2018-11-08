package main

import "fmt"

func Interface(arg ...interface{})  {
}
func main()  {
	//空接口万能类型，可以保存任意类型的值
	var i interface{} = 1
	fmt.Println("i = ", i)

	i = "abc"
	fmt.Println("i = ", i)

	Interface("")

}
