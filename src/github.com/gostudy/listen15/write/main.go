package main

import (
"os"
"fmt"
)

func main()  {
	file, err := os.OpenFile("./test.txt",os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)  //trunc清空
	if err != nil {
		fmt.Println("open files is failed,err", err)
		return
	}

	defer file.Close()
	str := "hello"
	file.Write([]byte(str))
	file.WriteString(str)
}
