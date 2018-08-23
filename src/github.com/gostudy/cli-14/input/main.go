package main

import "fmt"

func testScanf()  {
	var a int
	var b string
	var c float32
	//fmt.Scanf("%d%s%f",&a,&b, &c) 方法1
	fmt.Scanf("%d\n", &a)  // \n至到读到\n为止
	fmt.Scanf("%s\n", &b)
	fmt.Scanf("%f\n", &c)
	fmt.Printf("a=%d b=%s c=%f\n", a,b,c)
}

func testScan()  {
	var a int
	var b string
	var c float32
	fmt.Scan(&a,&b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a,b,c)
}

func testScanln()  {
	var a int
	var b string
	var c float32
	fmt.Scanln(&a,&b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a,b,c)
}
func main()  {

	//testScanf()
	//testScan()
	testScanln()   //a b c三个值必须在一行，遇到换行符结束

}