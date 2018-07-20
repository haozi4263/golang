package main

import (
	"fmt"
	"math/rand"  //随机包
	"time"
	"sort"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	var a map[string]int = make(map[string]int,1024)

	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("stu%d",i)
		value := rand.Intn(1000)
		a[key] = value
	}

	var keys []string = make([]string,2,128)  //定义切片
	for key,value := range a {
		fmt.Printf("map[%s]=%d\n",key,value)
		keys = append(keys,key)
	}

	sort.Strings(keys)			//对切片进行排序
	for _, value := range keys {   //便利map
		fmt.Printf("key:%s val:%d\n",value,a[value])
	}
}
