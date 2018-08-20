package log

import (
	"fmt"
)

type ConsoleLog struct {

}

func NewConsoleLog(file string) LogInterface  {   //返回日志库接口
	return &ConsoleLog{}
}

func (f *ConsoleLog) LogDebug(msg string) {  //实现接口方法
	fmt.Println("console",msg)
}

func (f *ConsoleLog) LogWarn(msg string) {
	fmt.Println("console",msg)
}
