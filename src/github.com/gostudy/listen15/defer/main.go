package main

import "fmt"

func funA() int  {   //返回值匿名
	x := 5
	defer func() {
		x += 1  //x+1=6，返回值还是5
	}()

	return x   //x=5,返回值设置5
}


func funB() (x int) {   //返回值匿名
	defer func() {
		x += 1  //x+1=6，返回值还是5
	}()

	return 5   //x=5,返回值设置5
}


func funC() (y int)  {   //返回值匿名
	x := 5
	defer func() {
		x += 1  //x+1=6，返回值还是5
	}()

	return x   //x=5,返回值设置5,y=x=5
}



func funD() (x int) {   //返回值匿名
	defer func(x int) {
		x += 1  //x+1=6，返回值还是5
	}(x)

	return 5   //x=5,返回值设置5
}



func main()  {
	fmt.Println(funA())  //x=5
	fmt.Println(funB())	 //x=6
	fmt.Println(funC())	 //x=5
	fmt.Println(funD())	 //x=5
}