package main

import "fmt"
import (
	"os/exec"
	"strings"
	"bytes"
)


func main()  {
	f, err := exec.LookPath("ls /")  //LookPath在环境变量中查找科执行二进制文件
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)  //bin/ls


	cmd := exec.Command("ls", "/jenkins")  //查看当前目录下文件
	out, err1 := cmd.CombinedOutput()
	if err1 != nil {
		fmt.Println(err)
		}
	fmt.Println(string(out))
}


//	cmd:= exec.Command("ping", "www.baidu.com")
//	cmd.Stdin = strings.NewReader("some input")
//	var out bytes.Buffer
//	cmd.Stdout = &out
//	err1 := cmd.Run()
//	if err != nil {
//		fmt.Println(err1)
//	}
//
//	fmt.Printf("in all caps: %q\n", out.String())
//}
