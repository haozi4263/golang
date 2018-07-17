package main

import (
	"time"
	"fmt"

)

func testTime() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	send := now.Second()

	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, send)

	timestamp := now.Unix()
	fmt.Printf("timestamp is:%d\n", timestamp) //时间戳

}
func testTimestamp(timestamp int64) {

	timeObj := time.Unix(timestamp,0)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	send := timeObj.Second()

	fmt.Printf("current timestamp:%d\n",timestamp)
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n",year,month,day,hour,minute,send)

	}

func processTask()  {
	fmt.Printf("DO task\n")
}
func testTicker()  {
	ticker := time.Tick(1*time.Second)
	for i := range  ticker {    		//每秒将当前时间放入到管道中给for循环
		fmt.Printf("%v\n",i)
		processTask()   //使用定时器执行任务
	}
}

func testConst() {
	fmt.Printf("nano second:%d\n", time.Nanosecond)
	fmt.Printf("micro second:%v\n", time.Microsecond)
	fmt.Printf("mili second:L%d\n", time.Millisecond)
	fmt.Printf("second:%d\n", time.Second)

}

func testFormat()  {
	now := time.Now()
	timeVar := now.Format("2006/1/2 15:04:05")  //go诞生时间必须写这个时间点作为模板，使用2006-1-2
	fmt.Printf("time:%s\n",timeVar)
}

func testFormat2()  {
	now := time.Now()
	timeVar := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d\n",now.Year(), now.Month(), now.Day() ,now.Hour(), now.Minute(), now.Second()) //printf会把值返回到终端
	fmt.Printf("time:%s\n",timeVar)
}




// 获取代码执行时间，精确到微秒，当前时间-代码执行完成时间
func testCost()  {
	start := time.Now().UnixNano()   //纳秒
	for i :=0; i < 10; i ++ {
		time.Sleep(time.Millisecond)    //业务代码
	}
	end := time.Now().UnixNano()
	cost := (end - start)/1000    //换成纳秒
	fmt.Printf("code run:%d us\n",cost)
}


func main()  {
	//testTime()
	timestamp := time.Now().Unix()
	testTimestamp(timestamp)
	//testTicker()
	testConst()
	testFormat()
	testFormat2()
	testCost()
}
