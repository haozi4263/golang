package main

import (
	"fmt"
	"sort"
)

func main()  {
	var a[5]int = [5]int{1,2,3,4,5}  //定义数组
	sort.Ints(a[:])   //对切片进行排序
	fmt.Println("a",a)

	var b[5]string = [5]string{"ac","ec","be","fa","ii"}//对字符串进行排序，按字母abcde
	sort.Strings(b[:])
	fmt.Println("b",b)


	var c[5]float64 = [5]float64{11.11,11.333,23123.44,0.22,1.223}//对浮点型排序
	sort.Float64s(c[:])
	fmt.Println("c",c)
}
