package main

import "fmt"
import "errors"

func main()  {
	err1 := fmt.Errorf("%s","this is normal err")
	fmt.Println("err1: ",err1)

	err2 := errors.New("this is erro2")
	fmt.Println("err2: ",err2)
}