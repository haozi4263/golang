package main

import "fmt"

func test(a interface{})  {
	// s := a.(int)
	s, ok := a.(int)			//类型断言
	if !ok {
		fmt.Println(s)
		return
	}

}

func testInterface1()  {
	var a int = 100
	test(a)
	var b string = "golang"
	test(b)

}

func testSwitch(a interface{})  {
	switch a.(type) {
	case string:
		fmt.Printf("a is string,value:%v\n", a.(string))
	case int:
		fmt.Printf("a is int value:%v\n", a.(int))
	case int32:
		fmt.Printf("a is int value:%v\n", a.(int32))
	default:
		fmt.Println("no support type")
	}
}

func testSwitch1(a interface{})  {  //方法2，推荐使用
	switch v := a.(type) {
	case string:
		fmt.Printf("a is string,value:%v\n", v)
	case int:
		fmt.Printf("a is int value:%v\n", v)
	case int32:
		fmt.Printf("a is int value:%v\n", v)
	default:
		fmt.Println("no support type")
	}
}


func testInterface2()  {
	var a int = 100
	testSwitch(a)
	var b string = "golang"
	testSwitch(b)
}

func testInterface3()  {
	var a int = 100
	testSwitch(a)
	var b string = "golang"
	testSwitch(b)
}

func main()  {
	//testInterface1()
	testInterface2()
	testInterface3()
}