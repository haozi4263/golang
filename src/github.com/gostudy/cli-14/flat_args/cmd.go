package main

import (
	"flag"
	"fmt"
	"os/exec"
)

var recusive bool
var app string
var port int

func init()  {
	//
	flag.BoolVar(&recusive, "r", false, " 使用-r选项")
	flag.StringVar(&app, "t", "hello golang", "使用-t选项")
	flag.IntVar(&port, "p", 8888, "使用-p选项")


	flag.Parse()  //使选项参数生效
}

func main()  {
	fmt.Printf("recusive:%v\n", recusive)
	fmt.Printf("app:%v\n", app)
	fmt.Printf("port:%v\n",port)

	ping1()
}


func ping1()  {

		app := `ls /
			pwd
			ls -l
            docker ps
		`
		cmd := exec.Command("/bin/bash", "-c", app)
		out, _ :=cmd.Output()
		fmt.Println(string(out))
		//fmt.Println("app:", app)

}
