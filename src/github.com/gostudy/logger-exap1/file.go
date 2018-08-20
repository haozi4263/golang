package logger

import (
	"os"
	"fmt"
	"time"
	"path"
)

type FileLogger struct {
	level int
	logPath string
	logName string
	funcName string
	file	 *os.File
	warnfile *os.File
}

func (f *FileLogger) init()  {   //初始化结构体日志
	logname := fmt.Sprintf("%s/%s.log",f.logPath, f.logName)
	file, err := os.OpenFile(logname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err:%v", logname, err))
	}

	f.file = file

	//写错误日志和fatal日志的文件
	logname = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	file, err = os.OpenFile(logname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err:%v", logname, err))
	}

	f.warnfile = file

}

func NewLogger(level int, logPath string, logName string) LoggerInterfacel  {
	logger :=  &FileLogger{
		level:level,
		logPath:logPath,
		logName:logName,
	}
	logger.init()
	return logger
}

func (f *FileLogger) writeLogger(file *os.File, level int, loging string, args ...interface{})  {
	if (f.level > Debug)  {
		return
	}

	now := time.Now()
	nowStr := now.Format("2006/1/2 15:04:05.999")// 传入时间模板
	levelStr := GetLevelint(level)

	fileName, funcName, lineNo := GetLineInfo()//打印日志当前文件名字，函数名字，行号
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(loging, args...)
	fmt.Fprintf(f.file, "%s %s (%s:%s:%d) %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg )

}

func (f *FileLogger) SetLevel(level int)  {
	if level < Debug || level > Fatal {
		level = Debug
	}
	f.level = level
}

func (f *FileLogger) Debug(loging string, args ...interface{})  {
	f.writeLogger(f.file, Debug,loging, args...)
}

func (f *FileLogger) Traceing(loging string, args ...interface{})  {
	f.writeLogger(f.file, Traceing,loging, args...)
}

func (f *FileLogger) Info(loging string, args ...interface{})  {
	f.writeLogger(f.file, Info,loging, args...)
}

func (f *FileLogger) Warn(loging string, args ...interface{})  {
	f.writeLogger(f.file, Warn,loging, args...)
}

func (f *FileLogger) Error(loging string, args ...interface{})  {
	f.writeLogger(f.file, Error,loging, args...)
}

func (f *FileLogger) Fatal(loging string, args ...interface{})  {
	f.writeLogger(f.file, Fatal,loging, args...)
}


func (f *FileLogger) Close() {
	f.file.Close()
	f.warnfile.Close()
}
