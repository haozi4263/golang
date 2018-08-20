package logger

import (
	"fmt"
	"path"
	"runtime"
	"time"
	"os"
)

type LogData struct {
	Message      string
	TimeStr      string
	LevelStr     string
	Filename     string
	FuncName     string
	LineNo       int
	WarnAndFatal bool
}

//util.go 10

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()//获取函数名字
		lineNo = line
	}
	return
}

/*
1. 当业务调用打日志的方法时，我们把日志相关的数据写入到chan（队列）
2. 然后我们有一个后台的线程不断的从chan里面获取这些日志，最终写入到文件。
*/

func writeLog(file *os.File, level int, format string, args ...interface{})  {
	now := time.Now()
	nowStr := now.Format("2006/1/2 15:04:05.999")// 传入时间模板
	levelStr := getLevelText(level)  //获取日志级别字符串，返回日志级别数字0 1 2 3 4 5 6

	fileName, funcName, lineNo := GetLineInfo()//打印日志当前文件名字，函数名字，行号
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg )
}