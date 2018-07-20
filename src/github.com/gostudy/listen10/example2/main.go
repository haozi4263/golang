package main

import "fmt"

func modfiy(a *int)  {
	*a = 10000
}

func main()  {
	var a int = 100
	fmt.Printf("before modify: %d addr:%p\n",a,&a)
	modfiy(&a)
	fmt.Printf("after modify: %d addr:%p\n",a,&a)
}