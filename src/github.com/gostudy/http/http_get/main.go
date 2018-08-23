package main

import (
	"fmt"
	"net/http"
)

func main()  {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http.get err =", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("status = ", resp.Status)
	fmt.Println("statusCode = ", resp.StatusCode)
	fmt.Println("Header = ", resp.Header)
	//fmt.Println("body = ", resp.Body)

	buf := make([]byte, 4*1024)
	var tmp string

	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("read err = ", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("tmp = ", tmp)


}
