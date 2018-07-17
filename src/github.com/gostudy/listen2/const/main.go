package main

import (
	"fmt"
)

func main()  {
	const a int = 100  //必须先赋值
	const b string = "hello"
	fmt.Printf("a=%d\n b=%s\n",a,b)  //%sd为占位符，会被a的值替换

	const (
		c int = 100
		d string = "hellow"
		e   //会继承d值
	)
	fmt.Printf("c=%d\n d=%s\n e=%s\n ",c,d,e)

	const (
		f = iota  //f=0,g=1,h=2 iota递增
		g
		h
	)
	fmt.Printf("f=%d\n g=%d\n h=%d\n ",f,g,h)

	const (
		a1 = 1 << iota  //0移1位，a1=1,iota默认值位0
		a2 = 1 << iota  //1移1位，a2=2
		a3 = 1 << iota  //2移2位，a3=4
	)
	fmt.Printf("a1=%d\n a2=%d\n a3=%d\n ",a1,a2,a3)




	const (
		A = iota  //A=0
		B		  //B++=1
		C		  //C++=2
		D = 8	  //D=8
		E		  //继承D=8
		F = iota //A-F递增到5
		G		 //G++=6
	)
	fmt.Printf("A=%d\n B=%d\n C=%d\n D=%d\n E=%d\n F=%d\n G=%d\n" ,A,B,C,D,E,F,G)
}
