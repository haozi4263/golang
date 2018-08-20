package log

import (
	"fmt"
)

type FileLog struct {

}

func NewFileLog(file string) LogInterface  {  //返回log的接口
	return &FileLog{}
}

func (f *FileLog) LogDebug(msg string) {
	fmt.Println("file",msg)
}

func (f *FileLog) LogWarn(msg string) {
	fmt.Println("file",msg)
}
