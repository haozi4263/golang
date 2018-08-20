package logger

import (
	"os"
	"fmt"
	//"time"
	//"path"

)

type FileLogger struct {
	level int
	logPath string
	logName string
	flie 	*os.File
	warnfile *os.File
}



func (f *FileLogger) init()  {   //初始化结构体日志
	filename := fmt.Sprintf("%s/%s.log",f.logPath, f.logName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err:%v", filename, err))
	}

	f.flie = file

	//写错误日志和fatal日志的文件
	filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err:%v", filename, err))
	}

	f.warnfile = file

}


func NewFileLogger(config map[string]string ) (log LogInterface, err error) {  //实例化
	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("not found log_path")
		return
	}
	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("not found log_name")
		return
	}

	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found log_level")
		return
	}
	level := getLogLevel(logLevel)
	log =  &FileLogger{
		level:level,
		logPath:logPath,
		logName:logName,
	}
	log.Init()
	return
}




func (f *FileLogger) SetLevel(level int)  {   //接口方法
	if level < LogLevelDebug || level > LogLevelFatal {  //用户输入错误将日志定义为debug
		level = LogLevelDebug
	}
	f.level = level
}



func (f *FileLogger) Debug(format string, args ...interface{})  {   //接口方法
	if (f.level > LogLevelDebug) {
		return
	}
	writeLog(f.flie, LogLevelDebug, format, args...)

}

func (f *FileLogger) Trace(format string, args ...interface{})  {   //接口方法
	if (f.level > LogLevelTrace) {
		return
	}
	writeLog(f.flie, LogLevelTrace, format, args...)
}

func (f *FileLogger) Info(format string, args ...interface{})  {   //接口方法
	if (f.level > LogLevelInfo) {
		return
	}
	writeLog(f.flie, LogLevelInfo, format, args...)
}

func (f *FileLogger) Warn(format string, args ...interface{})  {   //接口方法
	if (f.level > LogLevelWarn) {
		return
	}
	writeLog(f.warnfile, LogLevelWarn, format, args...)
}

func (f *FileLogger) Error(format string, args ...interface{})  {   //接口方法
	if (f.level > LogLevelError) {
		return
	}
	writeLog(f.warnfile, LogLevelError, format, args...)
}

func (f *FileLogger) Fatal(format string, args ...interface{})  {   //接口方法
	if (f.level > LogLevelFatal) {
		return
	}
	writeLog(f.warnfile, LogLevelFatal, format, args...)
}

func (f *FileLogger) Close() {
	f.flie.Close()
	f.warnfile.Close()
}

func (f *FileLogger) Init() {

}