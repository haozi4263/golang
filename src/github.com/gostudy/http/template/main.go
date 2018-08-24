package main

import (
	"fmt"
	"html/template"
	"net/http"
	//"os"
	"os"
)

type UserInfo struct {
	Name string
	Sex string
	Age int
}

func login(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./login.html")
	if err != nil {
		fmt.Fprintf(w, "load login.html is failed")
		return
	}

	user := UserInfo{
		Name: "jude",
		Sex: "男",
		Age: 11,
	}
	t.Execute(w, "jude")
	t.Execute(w, 1000)

	//m := make(map[string_pkg]interface{})
	//m["username"] = "jude"
	//m["sex"] = "男"
	//m["age"] = 18
	////t.Execute(w, "jude")
	//t.Execute(w, m)
	//t.Execute(os.Stdout, m)

}


func main()  {
	http.HandleFunc("/index", login)
	fmt.Printf("http server is running 8080....")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("listen server failed, err:%v", err)
		return
	}
}
