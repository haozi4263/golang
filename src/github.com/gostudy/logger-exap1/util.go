package logger

import "runtime"

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()//获取函数名字
		lineNo = line
	}
	return
}
