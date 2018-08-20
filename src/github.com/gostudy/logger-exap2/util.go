package logger

import (
	"runtime"
	"os"
	"time"
	"path"
	"fmt"
)

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()//获取函数名字
		lineNo = line
	}
	return
}



func writeLogger(file *os.File, level int, loging string, args ...interface{})  {
	now := time.Now()
	nowStr := now.Format("2006/1/2 15:04:05.999")// 传入时间模板
	levelStr := GetLevelint(level)

	fileName, funcName, lineNo := GetLineInfo()//打印日志当前文件名字，函数名字，行号
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(loging, args...)
	fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg )

}
