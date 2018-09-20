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

	_, err = c.Do("lpush", "book_list","java", "php", "python")
	if err  != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.String(c.Do("lpop", "book_list"))  //hou jing xian chu xiaofei wan hou duilei zhong bucunzai
	r, err := redis.String(c.Do("lpop", "book_list"))  //xian jing xian chu
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}