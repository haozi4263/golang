package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func login(w http.ResponseWriter, r *http.Request)  {
	method := r.Method
	if method == "GET" {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			fmt.Fprintf(w, "load login.html failed")
			return
		}

		t.Execute(w, nil)  //t发送给浏览器，w是io对象
	}else if method == "POST" {
		r.ParseForm()  //解析表单
		username := r.FormValue("username")
		passwd := r.FormValue("password")
		fmt.Printf("username:%s\n", username)  //提取html表单key
		fmt.Printf("password:%s\n", passwd)

		if (username == "jude" && passwd == "123") {
			fmt.Fprintf(w, "username:%s login success\n", r.FormValue("username"))
		} else {
			fmt.Fprintf(w, "username:%s login faild\n", r.FormValue("username"))
		}

	}
}


func main()  {
	http.HandleFunc("/login", login)
	fmt.Printf("http servre is running 8888....")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf("http server lister is failed, err:%v", err)
		return
	}
}