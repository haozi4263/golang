package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	c, err := redis.Dial("tcp", "192.168.10.106:6379")
	if err != nil {
		fmt.Printf("conn redis is failed, err:%v\n", err)
		return
	}
	defer c.Close()
	_, err = c.Do("Set", "abc", 1000)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("Get", "abc"))  //jieguo zhuan int
	if err != nil {
		fmt.Println("get abc failed, ", err)
		return
	}
	fmt.Println(r)
}
