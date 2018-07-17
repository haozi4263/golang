package main

import (
	"fmt"

)

func testIf() {
	var a int = 3
	if a == 1 {
		fmt.Printf("a=1\n")
	} else if a == 2 {
		fmt.Printf("a=2\n")
	} else if a == 3 {
		fmt.Printf("a=3\n")
	} else if a == 4 {
		fmt.Printf("a=4\n")
	} else {
		fmt.Printf("a=5\n")
	}

}

func testSwitch()  {
	var a int = 4
	switch a {
	case 1:
		fmt.Printf("a=1\n")
	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
	case 4:
		fmt.Printf("a=4\n")
	case 5:
		fmt.Printf("a=5\n")
	}
}


func testSwitchV2()  {
	//var a int = 4
	switch a := 3; a {    //变量放在saitch内
	case 1:
		fmt.Printf("a=1\n")
	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
	case 4:
		fmt.Printf("a=4\n")
	case 5:
		fmt.Printf("a=5\n")
	default:
		fmt.Printf("invalid a =%d\n",a)  //未命中会走defalut
	}
}

func testSwitchV3()  {
	switch a := 6; a {    //变量放在saitch内
	case 1,2,3,4,5:   //case后接多个值
		fmt.Printf("a>=1 and a <=5\n")
	case 6,7,8,9,10:
		fmt.Printf("a>=6 and a <=10\n")
	default:
		fmt.Printf("a >10")  //未命中会走defalut
	}
}

func testSwitchV4()  {
	var num = 60
	switch  {
	case num >= 0 && num <= 25:    //一个case分支命中，其他case分支不执行
		fmt.Printf("a>=0 and a<=25 a=%d\n",num)
	case num >= 25 && num <= 50:
		fmt.Printf("a>=0 and a<=25 a=%d\n",num)
	case num >= 50 && num <= 75:
		fmt.Printf("a>=50 and a<=75 a=%d\n",num)
	case num >= 75 && num <= 100:
		fmt.Printf("a>=75 and a<=100 a=%d\n",num)
	default:
		fmt.Printf("a >100\n")
	}
}



func testSwitchV5()  {
	var num = 60
	switch  {
	case num >= 0 && num <= 25:    //一个case分支命中，后面其他case分支不执行
		fmt.Printf("a>=0 and a<=25 a=%d\n",num)
	case num >= 25 && num <= 50:
		fmt.Printf("a>=0 and a<=25 a=%d\n",num)

	case num >= 50 && num <= 75:
		fmt.Printf("a>=50 and a<=75 a=%d\n",num)
		fallthrough			    //会穿透到下一个case语句中执行
	case num >= 75 && num <= 100:
		fmt.Printf("a>=75 and a<=100 a=%d\n",num)
	default:
		fmt.Printf("a >100\n")
	}
}


//久久乘法表
func testMulti()  {
	//1*1=1
	//1*2=2  2*2=4
	//1*3=3	 2*3=6	3*3=9

	for i :=1;i < 10; i++{
		for j := 1;j <= i; j++ {
			fmt.Printf("%d * %d = %d\t",j,i,j*i)  \\ \t等于4个空白字符
		}
		fmt.Println()
		}
}









func main()  {
	//testIf()
	//testSwitch()
	//testSwitchV2()
	//testSwitchV3()
	//testSwitchV4()
	//testSwitchV5()

	testMulti()

}
