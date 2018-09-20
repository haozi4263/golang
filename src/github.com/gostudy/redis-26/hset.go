package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	c, err := redis.Dial("tcp", "192.168.10.106:6379")
	if err != nil {
		fmt.Printf("'conn is failed, err:%v\n", err)
		return
	}
	defer c.Close()

	_, err = c.Do("HSet", "Books", "golang", 1000)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("HGet", "Books", "golang"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}