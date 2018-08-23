package main

import "fmt"
import (
	"net/http"
	"log"
)

func sayhello(w http.ResponseWriter, r *http.Request)  {  //初始化实例
	//r.ParseForm()  //解析参数，默认是不会解析的
	//fmt.Fprintf(w, "%v\n", r.Form) //这些信息是输入到服务器端的打印信息
	//fmt.Fprintf(w,"path:%s\n", r.URL.Path)
	//fmt.Fprintf(w,"scheme:%s\n", r.URL.Scheme)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header =", r.Header)
	fmt.Println("r.Body", r.Body)
	w.Write([]byte("给客户端回复数据\n"))
	fmt.Fprintf(w,"ping is ok \n")
}


func main()  {
	http.HandleFunc("/echo/ping",sayhello)  //设置访问的路由，接受函数
	fmt.Printf("http server is running ...")
	err := http.ListenAndServe(":9999", nil) //设置监听端口
	if err != nil {
		log.Fatal("server run  error")
		return
	}
}