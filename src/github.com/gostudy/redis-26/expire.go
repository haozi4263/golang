package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "192.168.10.106:6379")
	if err != nil {
		fmt.Printf("'conn is failed, err:%v\n", err)
		return
	}
	defer c.Close()

	_, err = c.Do("expire", "java", 5)
	if err != nil {
		fmt.Println(err)
		return
	}
}